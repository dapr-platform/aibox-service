package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: v_aibox_device_info
[ 0] id                                             VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] ip                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] build_time_str                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] latest_heart_beat_time                         TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] status_name                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] active_event_count                             INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 8] critical_event_count                           INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 9] major_event_count                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[10] minor_event_count                              INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[11] warning_event_count                            INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "MWndQmBHyIPvvGRbQGgMvrTma",    "name": "hhnggtRvpdCxoxtSTAsQFlULk",    "ip": "CtUsscCkbDLODwLrUyiLEUPZI",    "build_time_str": "cWATySaKWADAFuSZVFIltUDyw",    "latest_heart_beat_time": 50,    "status": 30,    "status_name": "JHqEFdKsWDbrMivgGESFwhsCi",    "active_event_count": 16,    "critical_event_count": 78,    "major_event_count": 80,    "minor_event_count": 17,    "warning_event_count": 99}


Comments
-------------------------------------
[ 0] Warning table: v_aibox_device_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_device_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Aibox_device_info_FIELD_NAME_id = "id"

	Aibox_device_info_FIELD_NAME_name = "name"

	Aibox_device_info_FIELD_NAME_ip = "ip"

	Aibox_device_info_FIELD_NAME_build_time_str = "build_time_str"

	Aibox_device_info_FIELD_NAME_latest_heart_beat_time = "latest_heart_beat_time"

	Aibox_device_info_FIELD_NAME_status = "status"

	Aibox_device_info_FIELD_NAME_status_name = "status_name"

	Aibox_device_info_FIELD_NAME_active_event_count = "active_event_count"

	Aibox_device_info_FIELD_NAME_critical_event_count = "critical_event_count"

	Aibox_device_info_FIELD_NAME_major_event_count = "major_event_count"

	Aibox_device_info_FIELD_NAME_minor_event_count = "minor_event_count"

	Aibox_device_info_FIELD_NAME_warning_event_count = "warning_event_count"
)

// Aibox_device_info struct is a row record of the v_aibox_device_info table in the  database
type Aibox_device_info struct {
	ID string `json:"id"` //设备ID

	Name string `json:"name"` //设备名称

	IP string `json:"ip"` //设备IP地址

	BuildTimeStr string `json:"build_time_str"` //设备构建时间

	LatestHeartBeatTime common.LocalTime `json:"latest_heart_beat_time"` //最近心跳时间

	Status int32 `json:"status"` //设备状态(0:离线，1:在线)

	StatusName string `json:"status_name"` //设备状态名称

	ActiveEventCount int32 `json:"active_event_count"` //活动事件总数

	CriticalEventCount int32 `json:"critical_event_count"` //紧急事件数

	MajorEventCount int32 `json:"major_event_count"` //严重事件数

	MinorEventCount int32 `json:"minor_event_count"` //轻微事件数

	WarningEventCount int32 `json:"warning_event_count"` //警告事件数

}

var Aibox_device_infoTableInfo = &TableInfo{
	Name: "v_aibox_device_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `设备ID`,
			Notes: `Warning table: v_aibox_device_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_device_info primary key column id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(36)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       36,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            `设备名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ip",
			Comment:            `设备IP地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "build_time_str",
			Comment:            `设备构建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "BuildTimeStr",
			GoFieldType:        "string",
			JSONFieldName:      "build_time_str",
			ProtobufFieldName:  "build_time_str",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "latest_heart_beat_time",
			Comment:            `最近心跳时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "LatestHeartBeatTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "latest_heart_beat_time",
			ProtobufFieldName:  "latest_heart_beat_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `设备状态(0:离线，1:在线)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status_name",
			Comment:            `设备状态名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "StatusName",
			GoFieldType:        "string",
			JSONFieldName:      "status_name",
			ProtobufFieldName:  "status_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "active_event_count",
			Comment:            `活动事件总数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ActiveEventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "active_event_count",
			ProtobufFieldName:  "active_event_count",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "critical_event_count",
			Comment:            `紧急事件数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "CriticalEventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "critical_event_count",
			ProtobufFieldName:  "critical_event_count",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "major_event_count",
			Comment:            `严重事件数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "MajorEventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "major_event_count",
			ProtobufFieldName:  "major_event_count",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "minor_event_count",
			Comment:            `轻微事件数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "MinorEventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "minor_event_count",
			ProtobufFieldName:  "minor_event_count",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "warning_event_count",
			Comment:            `警告事件数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "WarningEventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "warning_event_count",
			ProtobufFieldName:  "warning_event_count",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Aibox_device_info) TableName() string {
	return "v_aibox_device_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Aibox_device_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Aibox_device_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Aibox_device_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Aibox_device_info) TableInfo() *TableInfo {
	return Aibox_device_infoTableInfo
}
