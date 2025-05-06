package service

import (
	"context"
	"strconv"
	"time"

	"aibox-service/entity"
	"aibox-service/model"

	"github.com/dapr-platform/common"
)

// ProcessMessage 处理收到的消息
func ProcessMessage(message entity.Message) {
	switch message.GetType() {
	case entity.MessageTypeHeartbeat:
		if heartbeatMsg, ok := message.(*entity.HeartbeatMessage); ok {
			processHeartbeatMessage(heartbeatMsg)
		} else {
			common.Logger.Errorf("消息类型转换失败: 期望HeartbeatMessage, 实际类型: %T", message)
		}
	case entity.MessageTypeEvent:
		if eventMsg, ok := message.(*entity.EventMessage); ok {
			processEventMessage(eventMsg)
		} else {
			common.Logger.Errorf("消息类型转换失败: 期望EventMessage, 实际类型: %T", message)
		}
	default:
		common.Logger.Warnf("未知消息类型: %s", message.GetType())
	}
}

// processHeartbeatMessage 处理设备心跳消息
func processHeartbeatMessage(message *entity.HeartbeatMessage) {
	common.Logger.Infof("收到心跳消息: 设备ID=%s, 时间=%s", message.BoxID, message.Time)

	// 检查设备是否存在
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"id='"+message.BoxID+"'",
	)

	// 解析心跳时间
	heartbeatTime, err := time.Parse(time.RFC3339, message.Time)
	if err != nil {
		heartbeatTime = time.Now()
		common.Logger.Warnf("解析心跳时间失败: %v, 使用当前时间", err)
	}

	if device == nil {
		// 设备不存在，创建新设备
		common.Logger.Infof("创建新设备: ID=%s", message.BoxID)
		newDevice := model.Aibox_device{
			ID:                  message.BoxID,
			CreatedBy:           "system",
			CreatedTime:         common.LocalTime(time.Now()),
			UpdatedBy:           "system",
			UpdatedTime:         common.LocalTime(time.Now()),
			Name:                "AI盒子-" + message.BoxID,
			IP:                  "", // 需要后续更新
			BuildTimeStr:        message.BuildTime,
			LatestHeartBeatTime: common.LocalTime(heartbeatTime),
			Status:              1, // 在线
		}

		err = common.DbUpsert[model.Aibox_device](
			context.Background(),
			common.GetDaprClient(),
			newDevice,
			model.Aibox_deviceTableInfo.Name,
			"id",
		)
		if err != nil {
			common.Logger.Errorf("创建设备失败: %v", err)
		}
	} else {
		// 更新设备心跳信息
		common.Logger.Debugf("更新设备心跳: ID=%s", message.BoxID)
		device.LatestHeartBeatTime = common.LocalTime(heartbeatTime)
		device.BuildTimeStr = message.BuildTime
		device.Status = 1 // 设置为在线
		device.UpdatedTime = common.LocalTime(time.Now())
		device.UpdatedBy = "system"

		err = common.DbUpsert[model.Aibox_device](
			context.Background(),
			common.GetDaprClient(),
			*device,
			model.Aibox_deviceTableInfo.Name,
			"id",
		)
		if err != nil {
			common.Logger.Errorf("更新设备心跳失败: %v", err)
		}
	}
}

// processEventMessage 处理设备事件消息
func processEventMessage(message *entity.EventMessage) {
	common.Logger.Infof("收到事件消息: 设备ID=%s, 事件类型=%s, 级别=%s",
		message.BoxID, message.EventType, message.EventLevel)

	// 确保设备存在
	device, err := common.DbGetOne[model.Aibox_device](
		context.Background(),
		common.GetDaprClient(),
		model.Aibox_deviceTableInfo.Name,
		"id='"+message.BoxID+"'",
	)

	if device == nil {
		common.Logger.Warnf("事件关联设备不存在: %s, 尝试创建设备", message.BoxID)
		// 创建设备记录
		newDevice := model.Aibox_device{
			ID:                  message.BoxID,
			CreatedBy:           "system",
			CreatedTime:         common.LocalTime(time.Now()),
			UpdatedBy:           "system",
			UpdatedTime:         common.LocalTime(time.Now()),
			Name:                "AI盒子-" + message.BoxID,
			IP:                  "", // 需要后续更新
			BuildTimeStr:        "",
			LatestHeartBeatTime: common.LocalTime(time.Now()),
			Status:              1, // 在线
		}

		err = common.DbUpsert[model.Aibox_device](
			context.Background(),
			common.GetDaprClient(),
			newDevice,
			model.Aibox_deviceTableInfo.Name,
			"id",
		)
		if err != nil {
			common.Logger.Errorf("创建设备失败: %v", err)
			return
		}
	}

	// 解析事件级别
	levelInt := parseEventLevel(message.EventLevel)

	// 解析事件时间
	eventTime, err := time.Parse(time.RFC3339, message.Time)
	if err != nil {
		eventTime = time.Now()
		common.Logger.Warnf("解析事件时间失败: %v, 使用当前时间", err)
	}

	// 创建事件记录
	event := model.Aibox_event{
		ID:          message.ID,
		CreatedBy:   "system",
		CreatedTime: common.LocalTime(eventTime),
		UpdatedBy:   "system",
		UpdatedTime: common.LocalTime(eventTime),
		Dn:          message.BoxID,
		Title:       formatEventTitle(message.EventType, message.EventLevel),
		DeviceID:    message.BoxID,
		Content:     message.EventMessage,
		Picstr:      message.EventPicture,
		Level:       int32(levelInt),
		Status:      1, // 活动
	}

	err = common.DbUpsert[model.Aibox_event](
		context.Background(),
		common.GetDaprClient(),
		event,
		model.Aibox_eventTableInfo.Name,
		"id",
	)
	if err != nil {
		common.Logger.Errorf("保存事件记录失败: %v", err)
	} else {
		common.Logger.Infof("成功保存事件: ID=%s, 级别=%d, 设备=%s",
			event.ID, event.Level, event.DeviceID)
	}
}

// parseEventLevel 解析事件级别字符串为整数
func parseEventLevel(levelStr string) int {
	// 尝试直接解析为整数
	level, err := strconv.Atoi(levelStr)
	if err == nil && level >= 1 && level <= 4 {
		return level
	}

	// 根据级别名称映射
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
		// 默认为警告级别
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
