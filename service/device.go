package service

import (
	"context"
	"fmt"
	"time"

	"aibox-service/model"

	"github.com/dapr-platform/common"
)

const (
	// 设备离线阈值，超过此时间未收到心跳则标记为离线
	DeviceOfflineThreshold = 3 * time.Minute
)

// InitDeviceService 初始化设备服务
func init() {
	common.Logger.Info("初始化设备服务...")
	go startDeviceMonitor()
	common.Logger.Info("设备服务初始化完成")
}

// startDeviceMonitor 启动设备监控协程
func startDeviceMonitor() {
	defer func() {
		if err := recover(); err != nil {
			common.Logger.Errorf("设备监控协程发生异常: %v", err)
			// 重启协程
			go startDeviceMonitor()
		}
	}()

	common.Logger.Info("启动设备监控协程")
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		checkOfflineDevices()
	}
}

// checkOfflineDevices 检查离线设备
func checkOfflineDevices() {
	common.Logger.Debug("开始检查离线设备")

	// 获取当前时间
	now := time.Now()
	offlineThreshold := now.Add(-DeviceOfflineThreshold)

	// 查询所有在线设备
	devices, err := common.DbQuery[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"status=1", // 只查询在线设备
	)

	if err != nil {
		common.Logger.Errorf("查询设备列表失败: %v", err)
		return
	}

	common.Logger.Debugf("找到 %d 个在线设备，检查心跳时间", len(devices))
	var offlineCount int

	// 遍历设备，检查心跳时间
	for _, device := range devices {
		// 将心跳时间转换为本地时间
		heartbeatTime := time.Time(device.LatestHeartBeatTime)

		// 如果心跳时间早于离线阈值，则标记为离线
		if heartbeatTime.Before(offlineThreshold) {
			common.Logger.Infof("设备 [%s] 心跳超时，最后心跳时间: %s，标记为离线",
				device.ID, heartbeatTime.Format(time.RFC3339))

			// 更新设备状态为离线
			device.Status = 0 // 离线
			device.UpdatedTime = common.LocalTime(now)
			device.UpdatedBy = "admin"

			err := common.DbUpsert[model.Aibox_device](
				context.Background(),
				common.GetDaprClient(),
				device,
				model.Aibox_deviceTableInfo.Name,
				"id",
			)

			if err != nil {
				common.Logger.Errorf("更新设备 [%s] 状态失败: %v", device.ID, err)
			} else {
				offlineCount++
			}
		}
	}

	if offlineCount > 0 {
		common.Logger.Infof("检测到 %d 台设备离线", offlineCount)
	} else {
		common.Logger.Debug("未检测到离线设备")
	}
}

// GetDeviceStatus 获取设备状态
func GetDeviceStatus(deviceId string) (string, error) {
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		fmt.Sprintf("id=%s", deviceId),
	)

	if err != nil {
		return "", err
	}

	if device == nil {
		return "", fmt.Errorf("设备 [%s] 不存在", deviceId)
	}

	// 根据状态码返回状态描述
	if device.Status == 1 {
		return "在线", nil
	} else {
		return "离线", nil
	}
}
