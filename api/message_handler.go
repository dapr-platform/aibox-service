package api

import (
	"encoding/json"
	"io"
	"net/http"

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
		http.Error(w, "读取请求失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 首先解析为基础消息，确定消息类型
	var baseMsg entity.BaseMessage
	if err := json.Unmarshal(body, &baseMsg); err != nil {
		common.Logger.Errorf("解析基础消息失败: %v", err)
		http.Error(w, "无法解析消息格式", http.StatusBadRequest)
		return
	}

	// 根据消息类型解析完整消息
	var message entity.Message
	switch baseMsg.Type {
	case entity.MessageTypeHeartbeat:
		var heartbeatMsg entity.HeartbeatMessage
		if err := json.Unmarshal(body, &heartbeatMsg); err != nil {
			common.Logger.Errorf("解析心跳消息失败: %v", err)
			http.Error(w, "解析心跳消息失败", http.StatusBadRequest)
			return
		}
		message = &heartbeatMsg

	case entity.MessageTypeEvent:
		var eventMsg entity.EventMessage
		if err := json.Unmarshal(body, &eventMsg); err != nil {
			common.Logger.Errorf("解析事件消息失败: %v", err)
			http.Error(w, "解析事件消息失败", http.StatusBadRequest)
			return
		}
		message = &eventMsg

	default:
		common.Logger.Warnf("未知消息类型: %s", baseMsg.Type)
		http.Error(w, "未知消息类型", http.StatusBadRequest)
		return
	}

	// 处理消息
	go service.ProcessMessage(message)

	// 返回成功
	response := entity.ResponseMessage{
		Action: "ack",
		Data: map[string]interface{}{
			"status": "success",
		},
	}

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		common.Logger.Errorf("返回响应失败: %v", err)
		http.Error(w, "返回响应失败", http.StatusInternalServerError)
	}
}
