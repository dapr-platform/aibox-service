package api

import (
	"encoding/json"
	"io"
	"net/http"

	"aibox-service/config"
	"aibox-service/entity"
	"aibox-service/service"

	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
)

func InitMessageRoute(r chi.Router) {
	r.Post(common.BASE_CONTEXT+"/message", ProcessMessageHandler)
}

// @Summary 处理消息
// @Description 处理消息
// @Tags 消息
// @Accept json
// @Produce json
// @Param message body entity.BaseMessage true "消息"
// @Success 200 {object} entity.ResponseMessage "成功"
// @Failure 400 {object} entity.ResponseMessage "失败"
// @Failure 500 {object} entity.ResponseMessage "失败"
// @Router /message [post]
func ProcessMessageHandler(w http.ResponseWriter, r *http.Request) {
	common.Logger.Debug("收到消息请求")

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.Logger.Errorf("读取请求体失败: %v", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	defer r.Body.Close()

	// 首先解析为基础消息，确定消息类型
	var baseMsg entity.BaseMessage
	if err := json.Unmarshal(body, &baseMsg); err != nil {
		common.Logger.Errorf("解析基础消息失败: %v", err)
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	// 根据消息类型解析完整消息
	var response entity.ResponseMessage

	switch baseMsg.Type {
	case entity.MessageTypeHeartbeat:
		var heartbeatMsg entity.HeartbeatMessage
		if err := json.Unmarshal(body, &heartbeatMsg); err != nil {
			common.Logger.Errorf("解析心跳消息失败: %v", err)
			common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
			return
		}

		// 同步处理心跳消息，并检查是否需要更新
		upgradeTask, err := service.ProcessHeartbeatMessage(&heartbeatMsg)
		if err != nil {
			common.Logger.Errorf("处理心跳消息失败: %v", err)
			common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
			return
		}
		if upgradeTask != nil {
			response = *upgradeTask
			common.Logger.Infof("设备需要更新，返回更新指令: %v", response)
			common.HttpResult(w, common.OK.WithData(response))
			return
		}

		// 设备不需要更新，返回普通确认
		response = entity.ResponseMessage{
			Action: "ack",
			Data: map[string]interface{}{
				"status": "success",
			},
		}
		if config.AUTO_UPGRADE {
			// 检查设备是否需要更新
			update, needUpdate := service.CheckDeviceUpdateNeeded(heartbeatMsg.BuildTime)
			if needUpdate && update != nil {
				// 设备需要更新，返回更新指令
				response = service.GetDeviceUpdateResponse(update, r)
				common.Logger.Infof("设备需要更新，返回更新指令: %v", response)
			} else {
				common.Logger.Infof("设备不需要更新")
			}
		}

	case entity.MessageTypeEvent:
		var eventMsg entity.EventMessage
		if err := json.Unmarshal(body, &eventMsg); err != nil {
			common.Logger.Errorf("解析事件消息失败: %v", err)
			common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
			return
		}

		// 异步处理事件消息
		go service.ProcessEventMessage(&eventMsg)

		// 返回确认
		response = entity.ResponseMessage{
			Action: "ack",
			Data: map[string]interface{}{
				"status": "success",
			},
		}

	default:
		common.Logger.Warnf("未知消息类型: %s", baseMsg.Type)
		common.HttpResult(w, common.ErrService.AppendMsg("未知消息类型"))
		return
	}

	common.HttpResult(w, common.OK.WithData(response))
}
