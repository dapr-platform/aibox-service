package model

import (
	"github.com/dapr-platform/common"
)

// Aibox_device AI盒子设备表
type Aibox_device struct {
	ID                  string           `json:"id"`
	CreatedBy           string           `json:"created_by"`
	CreatedTime         common.LocalTime `json:"created_time"`
	UpdatedBy           string           `json:"updated_by"`
	UpdatedTime         common.LocalTime `json:"updated_time"`
	Name                string           `json:"name"`
	IP                  string           `json:"ip"`
	BuildTimeStr        string           `json:"build_time_str"`
	LatestHeartBeatTime common.LocalTime `json:"latest_heart_beat_time"`
	Status              int32            `json:"status"` // 0:离线, 1:在线
}

// TableName 表名
func (m *Aibox_device) TableName() string {
	return "o_aibox_device"
}

// BeforeSave 保存前操作
func (m *Aibox_device) BeforeSave() error {
	return nil
}

// Prepare 准备操作
func (m *Aibox_device) Prepare() {
}

// Validate 验证操作
func (m *Aibox_device) Validate(action Action) error {
	return nil
}

// TableInfo 表信息
func (m *Aibox_device) TableInfo() *TableInfo {
	return Aibox_deviceTableInfo
}

// Aibox_event AI盒子事件表
type Aibox_event struct {
	ID          string           `json:"id"`
	CreatedBy   string           `json:"created_by"`
	CreatedTime common.LocalTime `json:"created_time"`
	UpdatedBy   string           `json:"updated_by"`
	UpdatedTime common.LocalTime `json:"updated_time"`
	Dn          string           `json:"dn"`
	Title       string           `json:"title"`
	DeviceID    string           `json:"device_id"`
	Content     string           `json:"content"`
	Picstr      string           `json:"picstr"`
	Level       int32            `json:"level"`  // 1:紧急, 2:严重, 3:轻微, 4:警告
	Status      int32            `json:"status"` // 0:清除, 1:活动
}

// TableName 表名
func (m *Aibox_event) TableName() string {
	return "o_aibox_event"
}

// BeforeSave 保存前操作
func (m *Aibox_event) BeforeSave() error {
	return nil
}

// Prepare 准备操作
func (m *Aibox_event) Prepare() {
}

// Validate 验证操作
func (m *Aibox_event) Validate(action Action) error {
	return nil
}

// TableInfo 表信息
func (m *Aibox_event) TableInfo() *TableInfo {
	return Aibox_eventTableInfo
}

// 表信息定义
var (
	Aibox_deviceTableInfo = &TableInfo{
		Name: "o_aibox_device",
		Columns: []*ColumnInfo{
			{Name: "id", GoFieldName: "ID", JSONFieldName: "id"},
			{Name: "created_by", GoFieldName: "CreatedBy", JSONFieldName: "created_by"},
			{Name: "created_time", GoFieldName: "CreatedTime", JSONFieldName: "created_time"},
			{Name: "updated_by", GoFieldName: "UpdatedBy", JSONFieldName: "updated_by"},
			{Name: "updated_time", GoFieldName: "UpdatedTime", JSONFieldName: "updated_time"},
			{Name: "name", GoFieldName: "Name", JSONFieldName: "name"},
			{Name: "ip", GoFieldName: "IP", JSONFieldName: "ip"},
			{Name: "build_time_str", GoFieldName: "BuildTimeStr", JSONFieldName: "build_time_str"},
			{Name: "latest_heart_beat_time", GoFieldName: "LatestHeartBeatTime", JSONFieldName: "latest_heart_beat_time"},
			{Name: "status", GoFieldName: "Status", JSONFieldName: "status"},
		},
	}

	Aibox_eventTableInfo = &TableInfo{
		Name: "o_aibox_event",
		Columns: []*ColumnInfo{
			{Name: "id", GoFieldName: "ID", JSONFieldName: "id"},
			{Name: "created_by", GoFieldName: "CreatedBy", JSONFieldName: "created_by"},
			{Name: "created_time", GoFieldName: "CreatedTime", JSONFieldName: "created_time"},
			{Name: "updated_by", GoFieldName: "UpdatedBy", JSONFieldName: "updated_by"},
			{Name: "updated_time", GoFieldName: "UpdatedTime", JSONFieldName: "updated_time"},
			{Name: "dn", GoFieldName: "Dn", JSONFieldName: "dn"},
			{Name: "title", GoFieldName: "Title", JSONFieldName: "title"},
			{Name: "device_id", GoFieldName: "DeviceID", JSONFieldName: "device_id"},
			{Name: "content", GoFieldName: "Content", JSONFieldName: "content"},
			{Name: "picstr", GoFieldName: "Picstr", JSONFieldName: "picstr"},
			{Name: "level", GoFieldName: "Level", JSONFieldName: "level"},
			{Name: "status", GoFieldName: "Status", JSONFieldName: "status"},
		},
	}
)
