package entity

const (
	MessageTypeHeartbeat = "heartbeat"
	MessageTypeEvent     = "event"
)

// Message 消息接口
type Message interface {
	GetType() string
}

// BaseMessage 基础消息结构
type BaseMessage struct {
	Type string `json:"type"`
}

func (m BaseMessage) GetType() string {
	return m.Type
}

type HeartbeatMessage struct {
	BaseMessage
	ID        string `json:"id"`
	Time      string `json:"time"`
	BoxID     string `json:"box_id"`
	IP        string `json:"ip"`
	BoxName   string `json:"box_name"`
	BuildTime string `json:"build_time"`
}

type EventMessage struct {
	BaseMessage
	ID           string `json:"id"`
	BoxID        string `json:"box_id"`
	Time         string `json:"time"`
	EventType    string `json:"event_type"`
	EventLevel   string `json:"event_level"`
	EventMessage string `json:"event_message"`
	EventPicture string `json:"event_picture"`
	Status       string `json:"status"`
}

type ResponseMessage struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}
