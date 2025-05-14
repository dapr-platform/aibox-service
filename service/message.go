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
	common.Logger.Debugf("[心跳处理] 开始处理设备心跳: ID=%s, IP=%s, 时间=%s, 版本=%s",
		message.BoxID, message.IP, message.Time, message.BuildTime)

	if err := validateHeartbeatMessage(message); err != nil {
		common.Logger.Errorf("[心跳处理] 心跳消息验证失败: %v", err)
		return nil, err
	}

	common.Logger.Infof("收到心跳消息: 设备ID=%s, 时间=%s", message.BoxID, message.Time)

	device, err := getOrCreateDevice(message)
	if err != nil {
		common.Logger.Errorf("[心跳处理] 获取或创建设备失败: %v", err)
		return nil, fmt.Errorf("处理设备信息失败: %v", err)
	}
	common.Logger.Debugf("[心跳处理] 获取设备信息成功: ID=%s, 构建时间=%s", device.ID, device.BuildTimeStr)

	if config.AUTO_UPGRADE {
		common.Logger.Debugf("[心跳处理] 自动升级已开启, 开始检查设备升级: ID=%s", message.BoxID)
		err = handleDeviceUpgrade(device, message)
		if err != nil {
			common.Logger.Errorf("[心跳处理] 处理设备升级失败: %v", err)
			// 继续执行，不中断心跳流程
		}
	} else {
		common.Logger.Debugf("[心跳处理] 自动升级未开启, 跳过升级检查: ID=%s", message.BoxID)
	}

	// 获取现有升级任务
	common.Logger.Debugf("[心跳处理] 开始获取设备现有升级任务: ID=%s", message.BoxID)
	upgradeTask, err = getExistingUpgradeTask(device)
	if err != nil {
		common.Logger.Errorf("[心跳处理] 获取升级任务失败: %v", err)
		// 继续执行，不影响心跳基本功能
		err = nil // 重置错误，避免整个心跳处理失败
	} else if upgradeTask != nil {
		common.Logger.Infof("[心跳处理] 找到设备升级任务: ID=%s, 版本=%v, 文件=%v",
			message.BoxID, upgradeTask.Data["version"], upgradeTask.Data["filename"])
	} else {
		common.Logger.Debugf("[心跳处理] 设备没有升级任务: ID=%s", message.BoxID)
	}

	common.Logger.Debugf("[心跳处理] 开始更新设备状态: ID=%s", message.BoxID)
	device.ModelInfo = message.ModelInfo
	if updateErr := updateDeviceStatus(device); updateErr != nil {
		common.Logger.Errorf("[心跳处理] 更新设备状态失败: %v", updateErr)
		// 如果没有其他错误，返回此错误
		if err == nil {
			err = fmt.Errorf("更新设备状态失败: %v", updateErr)
		}
	}

	common.Logger.Debugf("[心跳处理] 心跳处理完成: ID=%s, 是否返回升级任务: %v",
		message.BoxID, upgradeTask != nil)
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
	common.Logger.Debugf("[心跳验证] 心跳消息验证通过: ID=%s", message.BoxID)
	return nil
}

// getOrCreateDevice 获取或创建设备
func getOrCreateDevice(message *entity.HeartbeatMessage) (*model.Aibox_device, error) {
	common.Logger.Debugf("[设备处理] 开始查询设备: ID=%s", message.BoxID)
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"id="+message.BoxID,
	)
	if err != nil {
		common.Logger.Errorf("[设备处理] 数据库查询设备失败: %v", err)
		return nil, fmt.Errorf("查询设备失败: %v", err)
	}

	heartbeatTime := parseHeartbeatTime(message.Time)
	common.Logger.Debugf("[设备处理] 解析心跳时间: 原始=%s, 解析=%s",
		message.Time, heartbeatTime.Format("2006-01-02 15:04:05"))

	if device == nil {
		common.Logger.Debugf("[设备处理] 设备不存在, 开始创建新设备: ID=%s", message.BoxID)
		return createNewDevice(message, heartbeatTime)
	}

	common.Logger.Debugf("[设备处理] 设备已存在, 开始更新设备信息: ID=%s, 当前版本=%s",
		message.BoxID, device.BuildTimeStr)
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

	common.Logger.Debugf("[设备处理] 准备保存新设备到数据库: ID=%s, 名称=%s, 版本=%s",
		newDevice.ID, newDevice.Name, newDevice.BuildTimeStr)
	err := common.DbUpsert[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		newDevice,
		model.Aibox_deviceTableInfo.Name,
		"id",
	)
	if err != nil {
		common.Logger.Errorf("[设备处理] 创建设备数据库操作失败: %v", err)
		return nil, fmt.Errorf("创建设备失败: %v", err)
	}
	common.Logger.Infof("[设备处理] 成功创建新设备: ID=%s", message.BoxID)
	return &newDevice, nil
}

// updateDeviceInfo 更新设备信息
func updateDeviceInfo(device *model.Aibox_device, message *entity.HeartbeatMessage, heartbeatTime time.Time) {
	common.Logger.Debugf("[设备处理] 更新设备信息: ID=%s, 原版本=%s, 新版本=%s",
		device.ID, device.BuildTimeStr, message.BuildTime)

	device.LatestHeartBeatTime = common.LocalTime(time.Now())
	device.BuildTimeStr = message.BuildTime
	device.Status = 1 // 设置为在线
	device.IP = message.IP
	device.Name = message.BoxName
	device.DeviceTime = common.LocalTime(heartbeatTime)
	device.UpdatedTime = common.LocalTime(time.Now())
	device.UpdatedBy = "admin"

	common.Logger.Debugf("[设备处理] 设备信息已更新: ID=%s, 名称=%s, IP=%s",
		device.ID, device.Name, device.IP)
}

// handleDeviceUpgrade 处理设备升级,如果有需要升级的，就加入到设备的任务列表
func handleDeviceUpgrade(device *model.Aibox_device, message *entity.HeartbeatMessage) (err error) {
	common.Logger.Debugf("[升级处理] 开始检查设备是否需要应用升级: ID=%s, 当前版本=%s",
		device.ID, device.BuildTimeStr)

	upgradeInfo, hasAppUpdate := checkDeviceHasAppUpdate(device.BuildTimeStr)
	if hasAppUpdate {
		common.Logger.Infof("[升级处理] 设备需要应用升级: ID=%s, 当前版本=%s, 目标版本=%s",
			device.ID, device.BuildTimeStr, upgradeInfo.Version)

		downloadURL := buildDownloadURL(message.IP, upgradeInfo)
		common.Logger.Debugf("[升级处理] 构建下载URL: %s", downloadURL)

		return addUpgradeTask(device, upgradeInfo, downloadURL)
	}

	common.Logger.Debugf("[升级处理] 设备不需要应用升级, 检查模型升级: ID=%s", device.ID)
	common.Logger.Debugf("[升级处理] 设备模型信息: %s", message.ModelInfo)

	upgradeInfo, hasModelUpdate := checkDeviceHasModelUpdate(message.ModelInfo)
	if hasModelUpdate {
		common.Logger.Infof("[升级处理] 设备需要模型升级: ID=%s, 模型=%s, 版本=%s",
			device.ID, upgradeInfo.FileName, upgradeInfo.Version)

		downloadURL := buildDownloadURL(message.IP, upgradeInfo)
		common.Logger.Debugf("[升级处理] 构建下载URL: %s", downloadURL)

		return addUpgradeTask(device, upgradeInfo, downloadURL)
	}

	common.Logger.Debugf("[升级处理] 设备不需要任何升级: ID=%s", device.ID)
	return nil
}

func checkDeviceHasModelUpdate(modelInfoStr string) (upgradeInfo *model.Aibox_update_info, hasModelUpdate bool) {
	if modelInfoStr == "" {
		common.Logger.Debugf("[模型升级] 设备模型信息为空, 跳过检查")
		return nil, false
	}

	common.Logger.Debugf("[模型升级] 开始分析模型信息: %s", modelInfoStr)
	for _, modelInfo := range strings.Split(modelInfoStr, ",") {
		modelInfo = strings.TrimSpace(modelInfo)
		if modelInfo == "" {
			continue
		}

		modelInfoArray := strings.Split(modelInfo, ":")
		if len(modelInfoArray) != 2 {
			common.Logger.Warnf("[模型升级] 模型信息格式错误: %s", modelInfo)
			continue
		}

		modelName := modelInfoArray[0]
		modelVersion := modelInfoArray[1]
		common.Logger.Debugf("[模型升级] 检查模型: 名称=%s, 版本=%s", modelName, modelVersion)

		upgradeInfo, err := getLatestModelUpdateInfo(modelName)
		if err != nil {
			common.Logger.Warnf("[模型升级] 获取模型更新信息失败: %v", err)
			continue
		}

		// 检查版本是否为md5值
		if len(upgradeInfo.Version) != 32 {
			common.Logger.Warnf("[模型升级] 版本不是md5值, 跳过: %s", upgradeInfo.Version)
			continue
		}

		common.Logger.Debugf("[模型升级] 获取到最新模型信息: 名称=%s, 版本=%s, 当前版本=%s",
			upgradeInfo.FileName, upgradeInfo.Version, modelVersion)

		if upgradeInfo.Version != modelVersion {
			common.Logger.Infof("[模型升级] 发现模型需要升级: 名称=%s, 当前版本=%s, 最新版本=%s",
				modelName, modelVersion, upgradeInfo.Version)
			return upgradeInfo, true
		}
	}

	common.Logger.Debugf("[模型升级] 未发现需要升级的模型")
	return nil, false
}

func getLatestModelUpdateInfo(modelName string) (upgradeInfo *model.Aibox_update_info, err error) {
	common.Logger.Debugf("[模型升级] 查询模型最新信息: 名称=%s", modelName)
	updates, err := common.DbQuery[model.Aibox_update_info](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_update_infoTableInfo.Name,
		"type=2&status=1&_order=-updated_time&file_name="+modelName,
	)
	if err != nil {
		common.Logger.Errorf("[模型升级] 数据库查询模型更新信息失败: %v", err)
		return nil, fmt.Errorf("获取模型更新信息失败: %v", err)
	}
	if len(updates) == 0 {
		common.Logger.Warnf("[模型升级] 数据库中没有找到模型信息: %s", modelName)
		return nil, fmt.Errorf("没有找到可用的模型更新信息")
	}
	latestUpdate := updates[0]
	common.Logger.Debugf("[模型升级] 找到模型最新信息: 名称=%s, 版本=%s",
		latestUpdate.FileName, latestUpdate.Version)
	return &latestUpdate, nil
}

// buildDownloadURL 构建下载URL
func buildDownloadURL(deviceIP string, upgradeInfo *model.Aibox_update_info) string {
	params := url.Values{}
	params.Set("version", upgradeInfo.Version)
	params.Set("type", fmt.Sprintf("%d", upgradeInfo.Type))
	params.Set("filename", upgradeInfo.FileName)
	downloadURL := fmt.Sprintf("http://%s/api/aibox-service/file/download?%s", deviceIP, params.Encode())
	common.Logger.Debugf("[升级处理] 构建下载URL: %s", downloadURL)
	return downloadURL
}

// addUpgradeTask 添加升级任务
func addUpgradeTask(device *model.Aibox_device, upgradeInfo *model.Aibox_update_info, downloadURL string) error {
	common.Logger.Debugf("[升级任务] 开始添加升级任务: 设备=%s, 文件=%s, 版本=%s",
		device.ID, upgradeInfo.FileName, upgradeInfo.Version)

	if device.UpgradeTasks == "" {
		common.Logger.Debugf("[升级任务] 设备当前没有升级任务, 初始化为空数组")
		device.UpgradeTasks = "[]"
	}

	upgradeTasks := []entity.ResponseMessage{}
	if err := json.Unmarshal([]byte(device.UpgradeTasks), &upgradeTasks); err != nil {
		common.Logger.Errorf("[升级任务] 解析当前升级任务列表失败: %v, 原数据=%s",
			err, device.UpgradeTasks)
		return fmt.Errorf("解析升级任务失败: %v", err)
	}

	common.Logger.Debugf("[升级任务] 当前任务数量: %d", len(upgradeTasks))
	if !hasExistingUpgradeTask(upgradeTasks, upgradeInfo.Version) {
		common.Logger.Debugf("[升级任务] 未发现相同版本任务, 添加新任务: 版本=%s",
			upgradeInfo.Version)

		newTask := createUpgradeTask(upgradeInfo, downloadURL)
		upgradeTasks = append(upgradeTasks, newTask)
		common.Logger.Debugf("[升级任务] 任务添加后数量: %d", len(upgradeTasks))

		if err := updateDeviceUpgradeTasks(device, upgradeTasks); err != nil {
			common.Logger.Errorf("[升级任务] 更新设备升级任务失败: %v", err)
			return fmt.Errorf("更新设备升级任务失败: %v", err)
		}
		common.Logger.Infof("[升级任务] 成功添加升级任务: 设备=%s, 文件=%s, 版本=%s",
			device.ID, upgradeInfo.FileName, upgradeInfo.Version)
	} else {
		common.Logger.Infof("[升级任务] 已存在相同版本任务, 跳过添加: 设备=%s, 版本=%s",
			device.ID, upgradeInfo.Version)
	}

	return nil
}

// hasExistingUpgradeTask 检查是否存在相同版本的升级任务
func hasExistingUpgradeTask(tasks []entity.ResponseMessage, version string) bool {
	for _, task := range tasks {
		if task.Data["version"] == version {
			common.Logger.Debugf("[升级任务] 找到相同版本任务: 版本=%s", version)
			return true
		}
	}
	common.Logger.Debugf("[升级任务] 未找到相同版本任务: 版本=%s", version)
	return false
}

// createUpgradeTask 创建升级任务
func createUpgradeTask(upgradeInfo *model.Aibox_update_info, downloadURL string) entity.ResponseMessage {
	task := entity.ResponseMessage{
		Action: "upgrade",
		Data: map[string]interface{}{
			"filename":    upgradeInfo.FileName,
			"version":     upgradeInfo.Version,
			"url":         downloadURL,
			"md5":         upgradeInfo.FileKey,
			"description": upgradeInfo.Description,
		},
	}
	common.Logger.Debugf("[升级任务] 创建升级任务: 文件=%s, 版本=%s, URL=%s",
		upgradeInfo.FileName, upgradeInfo.Version, downloadURL)
	return task
}

// updateDeviceUpgradeTasks 更新设备升级任务
func updateDeviceUpgradeTasks(device *model.Aibox_device, tasks []entity.ResponseMessage) error {
	common.Logger.Debugf("[升级任务] 开始序列化升级任务列表: 数量=%d", len(tasks))
	tasksStr, err := json.Marshal(tasks)
	if err != nil {
		common.Logger.Errorf("[升级任务] 序列化升级任务失败: %v", err)
		return fmt.Errorf("序列化升级任务失败: %v", err)
	}

	device.UpgradeTasks = string(tasksStr)
	common.Logger.Debugf("[升级任务] 设备升级任务已更新: ID=%s, 任务JSON长度=%d",
		device.ID, len(device.UpgradeTasks))
	return nil
}

// getExistingUpgradeTask 获取现有的升级任务
func getExistingUpgradeTask(device *model.Aibox_device) (*entity.ResponseMessage, error) {
	common.Logger.Debugf("[升级任务] 开始获取设备现有升级任务: ID=%s", device.ID)

	if device.UpgradeTasks == "" {
		common.Logger.Debugf("[升级任务] 设备没有升级任务: ID=%s", device.ID)
		return nil, nil
	}

	common.Logger.Debugf("[升级任务] 解析升级任务JSON: 长度=%d", len(device.UpgradeTasks))
	var upgradeTasks []entity.ResponseMessage
	if err := json.Unmarshal([]byte(device.UpgradeTasks), &upgradeTasks); err != nil {
		common.Logger.Errorf("[升级任务] 解析升级任务失败: %v, 原始JSON=%s",
			err, device.UpgradeTasks)
		return nil, fmt.Errorf("解析升级任务失败: %v", err)
	}

	if len(upgradeTasks) == 0 {
		common.Logger.Debugf("[升级任务] 升级任务列表为空: ID=%s", device.ID)
		return nil, nil
	}

	common.Logger.Debugf("[升级任务] 找到升级任务: ID=%s, 数量=%d, 返回第一个任务",
		device.ID, len(upgradeTasks))
	upgradeTask := upgradeTasks[0]
	upgradeTasks = upgradeTasks[1:]
	common.Logger.Infof("[升级任务] 返回升级任务: 设备=%s, 文件=%v, 版本=%v",
		device.ID, upgradeTask.Data["filename"], upgradeTask.Data["version"])

	common.Logger.Debugf("[升级任务] 更新剩余任务列表: 剩余数量=%d", len(upgradeTasks))
	if err := updateDeviceUpgradeTasks(device, upgradeTasks); err != nil {
		common.Logger.Errorf("[升级任务] 更新剩余升级任务失败: %v", err)
		return nil, err
	}

	return &upgradeTask, nil
}

// updateDeviceStatus 更新设备状态
func updateDeviceStatus(device *model.Aibox_device) error {
	common.Logger.Debugf("[设备状态] 开始更新设备状态: ID=%s", device.ID)
	err := common.DbUpsert[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		*device,
		model.Aibox_deviceTableInfo.Name,
		"id",
	)
	if err != nil {
		common.Logger.Errorf("[设备状态] 更新设备状态失败: %v", err)
		return err
	}
	common.Logger.Debugf("[设备状态] 设备状态更新成功: ID=%s", device.ID)
	return nil
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
	err = createOrUpdateEvent(message, eventTime)
	if err != nil {
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
func createOrUpdateEvent(message *entity.EventMessage, eventTime time.Time) error {
	dn := message.BoxID + "-" + message.EventType
	if message.DN != "" {
		dn = message.DN
	}
	levelInt := parseEventLevel(message.EventLevel)

	isActive := (message.Status != "" && message.Status == "1")

	existActive, err := common.DbGetOne[model.Aibox_event](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_eventTableInfo.Name,
		"dn="+dn+"&status=1",
	)
	if err != nil {
		common.Logger.Errorf("查询事件失败: %v", err)
		return err
	}

	if isActive {
		if existActive != nil {
			existActive.Status = 0
			existActive.UpdatedTime = common.LocalTime(time.Now())
			if err := saveEvent(*existActive); err != nil {
				common.Logger.Errorf("更新已存在事件状态失败: %v", err)
				return err
			}
		}

		return saveEvent(model.Aibox_event{
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
		})
	} else {
		if existActive != nil {
			common.Logger.Infof("事件已存在: %s", dn)
			existActive.Status = 0
			existActive.UpdatedTime = common.LocalTime(time.Now())
			return saveEvent(*existActive)
		}
	}
	return nil
}

// saveEvent 保存事件
func saveEvent(event model.Aibox_event) error {
	if event.ID == "" {
		return nil
	}
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
	common.Logger.Debugf("[应用升级] 开始检查应用升级: 当前版本=%s", deviceBuildTime)

	updates, err := getLatestUpdateInfo()
	if err != nil {
		common.Logger.Errorf("[应用升级] 获取软件更新信息失败: %v", err)
		return nil, false
	}

	if len(updates) == 0 {
		common.Logger.Debugf("[应用升级] 没有找到可用的软件更新")
		return nil, false
	}

	latestUpdate := updates[0]
	common.Logger.Debugf("[应用升级] 找到最新应用版本: 版本=%s, 文件=%s",
		latestUpdate.Version, latestUpdate.FileName)

	if deviceBuildTime == "" {
		common.Logger.Warnf("[应用升级] 设备版本信息为空，无法比较")
		return nil, false
	}

	needsUpdate, err := compareVersions(deviceBuildTime, latestUpdate.Version)
	if err != nil {
		common.Logger.Warnf("[应用升级] 版本比较失败: %v", err)
		return nil, false
	}

	if needsUpdate {
		common.Logger.Infof("[应用升级] 发现需要更新: 设备版本=%s, 最新版本=%s",
			deviceBuildTime, latestUpdate.Version)
		return &latestUpdate, true
	}

	common.Logger.Debugf("[应用升级] 设备版本已是最新: %s", deviceBuildTime)
	return nil, false
}

// getLatestUpdateInfo 获取最新的更新信息
func getLatestUpdateInfo() ([]model.Aibox_update_info, error) {
	common.Logger.Debugf("[应用升级] 查询数据库获取最新应用版本")
	updates, err := common.DbQuery[model.Aibox_update_info](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_update_infoTableInfo.Name,
		"type=1&status=1&_order=-updated_time",
	)
	if err != nil {
		common.Logger.Errorf("[应用升级] 数据库查询应用版本失败: %v", err)
		return nil, err
	}
	common.Logger.Debugf("[应用升级] 查询到应用版本数量: %d", len(updates))
	return updates, nil
}

// compareVersions 比较版本号
func compareVersions(deviceVersion, updateVersion string) (bool, error) {
	common.Logger.Debugf("[版本比较] 比较版本: 设备=%s, 更新=%s", deviceVersion, updateVersion)

	deviceTime, err := time.Parse("2006-01-02_15:04:05", deviceVersion)
	if err != nil {
		common.Logger.Errorf("[版本比较] 解析设备版本时间失败: %v", err)
		return false, fmt.Errorf("解析设备版本时间失败: %v", err)
	}

	updateTime, err := time.Parse("2006-01-02_15:04:05", updateVersion)
	if err != nil {
		common.Logger.Errorf("[版本比较] 解析更新版本时间失败: %v", err)
		return false, fmt.Errorf("解析更新版本时间失败: %v", err)
	}

	result := updateTime.After(deviceTime)
	common.Logger.Debugf("[版本比较] 比较结果: 是否需要更新=%v, 设备时间=%s, 更新时间=%s",
		result, deviceTime.Format(time.RFC3339), updateTime.Format(time.RFC3339))
	return result, nil
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
