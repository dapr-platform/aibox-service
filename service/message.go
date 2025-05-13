package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"aibox-service/config"
	"aibox-service/entity"
	"aibox-service/model"

	"github.com/dapr-platform/common"
	"github.com/spf13/cast"
)

// ProcessHeartbeatMessage 处理设备心跳消息
func ProcessHeartbeatMessage(message *entity.HeartbeatMessage) (upgradeTask *entity.ResponseMessage, err error) {
	if err := validateHeartbeatMessage(message); err != nil {
		return nil, err
	}

	common.Logger.Infof("收到心跳消息: 设备ID=%s, 时间=%s", message.BoxID, message.Time)

	device, err := getOrCreateDevice(message)
	if err != nil {
		return nil, fmt.Errorf("处理设备信息失败: %v", err)
	}

	

	if config.AUTO_UPGRADE {
		err = handleDeviceUpgrade(device, message)
		if err != nil {
			common.Logger.Errorf("处理设备升级失败: %v", err)
			// 继续执行，不中断心跳流程
		} 
	}

	// 获取现有升级任务
	upgradeTask, err = getExistingUpgradeTask(device)
	if err != nil {
		common.Logger.Errorf("获取升级任务失败: %v", err)
		// 继续执行，不影响心跳基本功能
		err = nil // 重置错误，避免整个心跳处理失败
	}
	
	if updateErr := updateDeviceStatus(device); updateErr != nil {
		common.Logger.Errorf("更新设备状态失败: %v", updateErr)
		// 如果没有其他错误，返回此错误
		if err == nil {
			err = fmt.Errorf("更新设备状态失败: %v", updateErr)
		}
	}

	return upgradeTask, err
}

// validateHeartbeatMessage 验证心跳消息
func validateHeartbeatMessage(message *entity.HeartbeatMessage) error {
	if message == nil {
		return fmt.Errorf("心跳消息为空")
	}
	if message.BoxID == "" {
		return fmt.Errorf("设备ID为空")
	}
	return nil
}

// getOrCreateDevice 获取或创建设备
func getOrCreateDevice(message *entity.HeartbeatMessage) (*model.Aibox_device, error) {
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"id="+message.BoxID,
	)
	if err != nil {
		return nil, fmt.Errorf("查询设备失败: %v", err)
	}

	heartbeatTime := parseHeartbeatTime(message.Time)

	if device == nil {
		return createNewDevice(message, heartbeatTime)
	}

	updateDeviceInfo(device, message, heartbeatTime)
	return device, nil
}

// parseHeartbeatTime 解析心跳时间
func parseHeartbeatTime(timeStr string) time.Time {
	heartbeatTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		common.Logger.Warnf("解析心跳时间失败: %v, 原始时间=%s, 使用当前时间", err, timeStr)
		return time.Now()
	}
	return heartbeatTime
}

// createNewDevice 创建新设备
func createNewDevice(message *entity.HeartbeatMessage, heartbeatTime time.Time) (*model.Aibox_device, error) {
	common.Logger.Infof("创建新设备: ID=%s", message.BoxID)
	newDevice := model.Aibox_device{
		ID:                  message.BoxID,
		CreatedBy:           "admin",
		CreatedTime:         common.LocalTime(time.Now()),
		UpdatedBy:           "admin",
		UpdatedTime:         common.LocalTime(time.Now()),
		Name:                message.BoxName,
		IP:                  message.IP,
		BuildTimeStr:        message.BuildTime,
		DeviceTime:          common.LocalTime(heartbeatTime),
		LatestHeartBeatTime: common.LocalTime(time.Now()),
		Status:              1, // 在线
	}

	err := common.DbUpsert[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		newDevice,
		model.Aibox_deviceTableInfo.Name,
		"id",
	)
	if err != nil {
		return nil, fmt.Errorf("创建设备失败: %v", err)
	}
	return &newDevice, nil
}

// updateDeviceInfo 更新设备信息
func updateDeviceInfo(device *model.Aibox_device, message *entity.HeartbeatMessage, heartbeatTime time.Time) {
	device.LatestHeartBeatTime = common.LocalTime(time.Now())
	device.BuildTimeStr = message.BuildTime
	device.Status = 1 // 设置为在线
	device.IP = message.IP
	device.Name = message.BoxName
	device.DeviceTime = common.LocalTime(heartbeatTime)
	device.UpdatedTime = common.LocalTime(time.Now())
	device.UpdatedBy = "admin"
}

// handleDeviceUpgrade 处理设备升级,如果有需要升级的，就加入到设备的任务列表
func handleDeviceUpgrade(device *model.Aibox_device, message *entity.HeartbeatMessage) (err error) {
	upgradeInfo, hasAppUpdate := checkDeviceHasAppUpdate(device.BuildTimeStr)
	if hasAppUpdate {

		downloadURL := buildDownloadURL(message.IP, upgradeInfo)
		return addUpgradeTask(device, upgradeInfo, downloadURL)
	}
	upgradeInfo,hasModelUpdate := checkDeviceHasModelUpdate(message.ModelInfo)
	if hasModelUpdate{
		downloadURL := buildDownloadURL(message.IP, upgradeInfo)
		return addUpgradeTask(device, upgradeInfo, downloadURL)
	}
	return nil
}

func checkDeviceHasModelUpdate(modelInfoStr string)(upgradeInfo *model.Aibox_update_info,hasModelUpdate bool){
	if modelInfoStr == ""{
		return nil,false
	}
	for _,modelInfo := range strings.Split(modelInfoStr,","){
		modelInfo = strings.TrimSpace(modelInfo)
		if modelInfo == ""{
			continue
		}
		modelInfoArray := strings.Split(modelInfo,":")
		if len(modelInfoArray) != 2{
			common.Logger.Warnf("模型信息格式错误: %s", modelInfo)
			continue
		}
		modelName := modelInfoArray[0]
		modelVersion := modelInfoArray[1]
		upgradeInfo,err := getLatestModelUpdateInfo(modelName)
		if err != nil{
			common.Logger.Warnf("获取模型更新信息失败: %v",err)
			continue
		}
		if upgradeInfo.Version != modelVersion{
			return upgradeInfo,true
		}
	}

	return nil,false
}
func getLatestModelUpdateInfo(modelName string)(upgradeInfo *model.Aibox_update_info,err error){
	updates,err :=common.DbQuery[model.Aibox_update_info](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_update_infoTableInfo.Name,
		"type=2&status=1&_order=-updated_time&filename="+modelName,
	)
	if err != nil{
		return nil,fmt.Errorf("获取模型更新信息失败: %v",err)
	}
	if len(updates) == 0{
		return nil,fmt.Errorf("没有找到可用的模型更新信息")
	}
	latestUpdate := updates[0]
	return &latestUpdate,nil
}

// buildDownloadURL 构建下载URL
func buildDownloadURL(deviceIP string, upgradeInfo *model.Aibox_update_info) string {
	params := url.Values{}
	params.Set("version", upgradeInfo.Version)
	params.Set("type", "1")
	params.Set("filename", upgradeInfo.FileName)
	return fmt.Sprintf("http://%s/api/aibox-service/file/download?%s", deviceIP, params.Encode())
}

// addUpgradeTask 添加升级任务
func addUpgradeTask(device *model.Aibox_device, upgradeInfo *model.Aibox_update_info, downloadURL string) error {
	if device.UpgradeTasks == "" {
		device.UpgradeTasks = "[]"
	}

	upgradeTasks := []entity.ResponseMessage{}
	if err := json.Unmarshal([]byte(device.UpgradeTasks), &upgradeTasks); err != nil {
		return fmt.Errorf("解析升级任务失败: %v", err)
	}

	if !hasExistingUpgradeTask(upgradeTasks, upgradeInfo.Version) {
		newTask := createUpgradeTask(upgradeInfo, downloadURL)
		upgradeTasks = append(upgradeTasks, newTask)
		if err := updateDeviceUpgradeTasks(device, upgradeTasks); err != nil {
			return fmt.Errorf("更新设备升级任务失败: %v", err)
		}
	}

	return nil
}

// hasExistingUpgradeTask 检查是否存在相同版本的升级任务
func hasExistingUpgradeTask(tasks []entity.ResponseMessage, version string) bool {
	for _, task := range tasks {
		if task.Data["version"] == version {
			return true
		}
	}
	return false
}

// createUpgradeTask 创建升级任务
func createUpgradeTask(upgradeInfo *model.Aibox_update_info, downloadURL string) entity.ResponseMessage {
	return entity.ResponseMessage{
		Action: "upgrade",
		Data: map[string]interface{}{
			"filename":    upgradeInfo.FileName,
			"version":     upgradeInfo.Version,
			"url":         downloadURL,
			"md5":         upgradeInfo.FileKey,
			"description": upgradeInfo.Description,
		},
	}
}

// updateDeviceUpgradeTasks 更新设备升级任务
func updateDeviceUpgradeTasks(device *model.Aibox_device, tasks []entity.ResponseMessage) error {
	tasksStr, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("序列化升级任务失败: %v", err)
	}
	device.UpgradeTasks = string(tasksStr)
	return nil
}

// getExistingUpgradeTask 获取现有的升级任务
func getExistingUpgradeTask(device *model.Aibox_device) (*entity.ResponseMessage, error) {
	if device.UpgradeTasks == "" {
		return nil, nil
	}

	var upgradeTasks []entity.ResponseMessage
	if err := json.Unmarshal([]byte(device.UpgradeTasks), &upgradeTasks); err != nil {
		return nil, fmt.Errorf("解析升级任务失败: %v", err)
	}

	if len(upgradeTasks) == 0 {
		return nil, nil
	}

	upgradeTask := upgradeTasks[0]
	upgradeTasks = upgradeTasks[1:]

	if err := updateDeviceUpgradeTasks(device, upgradeTasks); err != nil {
		return nil, err
	}

	return &upgradeTask, nil
}

// updateDeviceStatus 更新设备状态
func updateDeviceStatus(device *model.Aibox_device) error {
	return common.DbUpsert[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		*device,
		model.Aibox_deviceTableInfo.Name,
		"id",
	)
}

// ProcessEventMessage 处理设备事件消息
func ProcessEventMessage(message *entity.EventMessage) {
	if err := validateEventMessage(message); err != nil {
		common.Logger.Errorf("事件消息验证失败: %v", err)
		return
	}

	common.Logger.Infof("收到事件消息: 设备ID=%s, 事件类型=%s, 级别=%s",
		message.BoxID, message.EventType, message.EventLevel)

	_, err := ensureDeviceExists(message)
	if err != nil {
		common.Logger.Errorf("确保设备存在失败: %v", err)
		return
	}

	eventTime := parseEventTime(message.Time)
	event := createOrUpdateEvent(message, eventTime)
	if err := saveEvent(event); err != nil {
		common.Logger.Errorf("保存事件失败: %v", err)
	}
}

// validateEventMessage 验证事件消息
func validateEventMessage(message *entity.EventMessage) error {
	if message == nil {
		return fmt.Errorf("事件消息为空")
	}
	if message.BoxID == "" {
		return fmt.Errorf("设备ID为空")
	}
	if message.EventType == "" {
		return fmt.Errorf("事件类型为空")
	}
	return nil
}

// ensureDeviceExists 确保设备存在
func ensureDeviceExists(message *entity.EventMessage) (*model.Aibox_device, error) {
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"id="+message.BoxID,
	)
	if err != nil {
		return nil, fmt.Errorf("查询设备失败: %v", err)
	}

	if device == nil {
		return createDeviceForEvent(message)
	}

	return device, nil
}

// createDeviceForEvent 为事件创建设备
func createDeviceForEvent(message *entity.EventMessage) (*model.Aibox_device, error) {
	common.Logger.Warnf("事件关联设备不存在: %s, 尝试创建设备", message.BoxID)
	newDevice := model.Aibox_device{
		ID:                  message.BoxID,
		CreatedBy:           "admin",
		CreatedTime:         common.LocalTime(time.Now()),
		UpdatedBy:           "admin",
		UpdatedTime:         common.LocalTime(time.Now()),
		Name:                "",
		IP:                  "",
		BuildTimeStr:        "",
		LatestHeartBeatTime: common.LocalTime(time.Now()),
		Status:              1, // 在线
	}

	err := common.DbUpsert[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		newDevice,
		model.Aibox_deviceTableInfo.Name,
		"id",
	)
	if err != nil {
		return nil, fmt.Errorf("创建设备失败: %v", err)
	}

	return &newDevice, nil
}

// parseEventTime 解析事件时间
func parseEventTime(timeStr string) time.Time {
	eventTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		common.Logger.Warnf("解析事件时间失败: %v, 使用当前时间", err)
		return time.Now()
	}
	return eventTime
}

// createOrUpdateEvent 创建或更新事件
func createOrUpdateEvent(message *entity.EventMessage, eventTime time.Time) model.Aibox_event {
	dn := message.BoxID + "-" + message.EventType
	levelInt := parseEventLevel(message.EventLevel)

	return model.Aibox_event{
		ID:          message.ID,
		CreatedBy:   "admin",
		CreatedTime: common.LocalTime(eventTime),
		UpdatedBy:   "admin",
		UpdatedTime: common.LocalTime(eventTime),
		Dn:          dn,
		Title:       formatEventTitle(message.EventType, message.EventLevel),
		DeviceID:    message.BoxID,
		Content:     message.EventMessage,
		Picstr:      message.EventPicture,
		Level:       int32(levelInt),
		Status:      cast.ToInt32(message.Status),
	}
}

// saveEvent 保存事件
func saveEvent(event model.Aibox_event) error {
	return common.DbUpsert[model.Aibox_event](
		context.Background(),
		common.GetDaprClient(),
		event,
		model.Aibox_eventTableInfo.Name,
		"id",
	)
}

// parseEventLevel 解析事件级别字符串为整数
func parseEventLevel(levelStr string) int {
	level, err := strconv.Atoi(levelStr)
	if err == nil && level >= 1 && level <= 4 {
		return level
	}

	switch levelStr {
	case "critical", "紧急":
		return 1
	case "major", "严重":
		return 2
	case "minor", "轻微":
		return 3
	case "warning", "警告":
		return 4
	default:
		common.Logger.Warnf("未知事件级别: %s, 默认设置为警告级别", levelStr)
		return 4
	}
}

// formatEventTitle 根据事件类型和级别格式化事件标题
func formatEventTitle(eventType, eventLevel string) string {
	levelName := ""
	switch parseEventLevel(eventLevel) {
	case 1:
		levelName = "紧急"
	case 2:
		levelName = "严重"
	case 3:
		levelName = "轻微"
	case 4:
		levelName = "警告"
	}
	return levelName + "-" + eventType
}

// checkDeviceHasAppUpdate 检查设备是否需要软件更新
func checkDeviceHasAppUpdate(deviceBuildTime string) (*model.Aibox_update_info, bool) {
	updates, err := getLatestUpdateInfo()
	if err != nil {
		common.Logger.Errorf("获取软件更新信息失败: %v", err)
		return nil, false
	}

	if len(updates) == 0 {
		common.Logger.Debugf("没有找到可用的软件更新")
		return nil, false
	}

	latestUpdate := updates[0]
	if deviceBuildTime == "" {
		common.Logger.Warnf("设备版本信息为空，无法比较")
		return nil, false
	}

	needsUpdate, err := compareVersions(deviceBuildTime, latestUpdate.Version)
	if err != nil {
		common.Logger.Warnf("版本比较失败: %v", err)
		return nil, false
	}

	if needsUpdate {
		common.Logger.Infof("发现需要更新: 设备版本=%s, 最新版本=%s", deviceBuildTime, latestUpdate.Version)
		return &latestUpdate, true
	}

	common.Logger.Debugf("设备版本已是最新: %s", deviceBuildTime)
	return nil, false
}

// getLatestUpdateInfo 获取最新的更新信息
func getLatestUpdateInfo() ([]model.Aibox_update_info, error) {
	return common.DbQuery[model.Aibox_update_info](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_update_infoTableInfo.Name,
		"type=1&status=1&_order=-updated_time",
	)
}

// compareVersions 比较版本号
func compareVersions(deviceVersion, updateVersion string) (bool, error) {
	deviceTime, err := time.Parse("2006-01-02_15:04:05", deviceVersion)
	if err != nil {
		return false, fmt.Errorf("解析设备版本时间失败: %v", err)
	}

	updateTime, err := time.Parse("2006-01-02_15:04:05", updateVersion)
	if err != nil {
		return false, fmt.Errorf("解析更新版本时间失败: %v", err)
	}

	return updateTime.After(deviceTime), nil
}

// GetDeviceUpdateResponse 获取设备更新响应消息
func GetDeviceUpdateResponse(update *model.Aibox_update_info, r *http.Request) entity.ResponseMessage {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	downloadURL := fmt.Sprintf("%s://%s/api/aibox-service/file/download?version=%s&type=1&filename=%s",
		scheme, r.Host, update.Version, url.QueryEscape(update.FileName))

	return entity.ResponseMessage{
		Action: "upgrade",
		Data: map[string]interface{}{
			"filename":    update.FileName,
			"version":     update.Version,
			"url":         downloadURL,
			"md5":         update.FileKey,
			"description": update.Description,
		},
	}
}
