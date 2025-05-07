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


Table: v_aibox_event_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] dn                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] title                                          VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] device_id                                      VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] content                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] picstr                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] level                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] level_name                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] status_name                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[11] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] device_name                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] device_ip                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] device_status                                  INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] device_status_name                             TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "lQjFlHKdaSkWBccBagRelVKAL",    "dn": "hBqDZmRivofvrsykoUQbHxUnj",    "title": "LhXYpCusVVALOfXMtbiBIbpDd",    "device_id": "rblcWXIttojhhLxcHIJvbLFoo",    "content": "rcuLrVVbDPRGMyFZUunDlCuln",    "picstr": "reEUWwBroiHhZmupkmTJEyGSu",    "level": 57,    "level_name": "grcJugSTmqrsZDPgXBsOeTJOZ",    "status": 80,    "status_name": "YEMDawamRGqgplvxIjdxnKfml",    "created_time": 67,    "updated_time": 84,    "device_name": "thDlAqjTgcdraSnOxZfAWtniH",    "device_ip": "drkxMIoJaUBgGNmRWhvHjRQOY",    "device_status": 41,    "device_status_name": "RkcYxobqWvlkxlpQsCrToenTL"}


Comments
-------------------------------------
[ 0] Warning table: v_aibox_event_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_event_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Aibox_event_info_FIELD_NAME_id = "id"

	Aibox_event_info_FIELD_NAME_dn = "dn"

	Aibox_event_info_FIELD_NAME_title = "title"

	Aibox_event_info_FIELD_NAME_device_id = "device_id"

	Aibox_event_info_FIELD_NAME_content = "content"

	Aibox_event_info_FIELD_NAME_picstr = "picstr"

	Aibox_event_info_FIELD_NAME_level = "level"

	Aibox_event_info_FIELD_NAME_level_name = "level_name"

	Aibox_event_info_FIELD_NAME_status = "status"

	Aibox_event_info_FIELD_NAME_status_name = "status_name"

	Aibox_event_info_FIELD_NAME_created_time = "created_time"

	Aibox_event_info_FIELD_NAME_updated_time = "updated_time"

	Aibox_event_info_FIELD_NAME_device_name = "device_name"

	Aibox_event_info_FIELD_NAME_device_ip = "device_ip"

	Aibox_event_info_FIELD_NAME_device_status = "device_status"

	Aibox_event_info_FIELD_NAME_device_status_name = "device_status_name"
)

// Aibox_event_info struct is a row record of the v_aibox_event_info table in the  database
type Aibox_event_info struct {
	ID string `json:"id"` //事件ID

	Dn string `json:"dn"` //设备编号

	Title string `json:"title"` //事件标题

	DeviceID string `json:"device_id"` //关联设备ID

	Content string `json:"content"` //事件内容

	Picstr string `json:"picstr"` //图片信息

	Level int32 `json:"level"` //事件级别(1:紧急, 2:严重, 3:轻微, 4:警告)

	LevelName string `json:"level_name"` //事件级别名称

	Status int32 `json:"status"` //事件状态(0:清除, 1:活动)

	StatusName string `json:"status_name"` //事件状态名称

	CreatedTime common.LocalTime `json:"created_time"` //创建时间

	UpdatedTime common.LocalTime `json:"updated_time"` //更新时间

	DeviceName string `json:"device_name"` //设备名称

	DeviceIP string `json:"device_ip"` //设备IP地址

	DeviceStatus int32 `json:"device_status"` //设备状态

	DeviceStatusName string `json:"device_status_name"` //设备状态名称

}

var Aibox_event_infoTableInfo = &TableInfo{
	Name: "v_aibox_event_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `事件ID`,
			Notes: `Warning table: v_aibox_event_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_event_info primary key column id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dn",
			Comment:            `设备编号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Dn",
			GoFieldType:        "string",
			JSONFieldName:      "dn",
			ProtobufFieldName:  "dn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "title",
			Comment:            `事件标题`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "device_id",
			Comment:            `关联设备ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "DeviceID",
			GoFieldType:        "string",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "content",
			Comment:            `事件内容`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Content",
			GoFieldType:        "string",
			JSONFieldName:      "content",
			ProtobufFieldName:  "content",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "picstr",
			Comment:            `图片信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Picstr",
			GoFieldType:        "string",
			JSONFieldName:      "picstr",
			ProtobufFieldName:  "picstr",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "level",
			Comment:            `事件级别(1:紧急, 2:严重, 3:轻微, 4:警告)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "level_name",
			Comment:            `事件级别名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "LevelName",
			GoFieldType:        "string",
			JSONFieldName:      "level_name",
			ProtobufFieldName:  "level_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
			Comment:            `事件状态(0:清除, 1:活动)`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status_name",
			Comment:            `事件状态名称`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "created_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "device_name",
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
			GoFieldName:        "DeviceName",
			GoFieldType:        "string",
			JSONFieldName:      "device_name",
			ProtobufFieldName:  "device_name",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "device_ip",
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
			GoFieldName:        "DeviceIP",
			GoFieldType:        "string",
			JSONFieldName:      "device_ip",
			ProtobufFieldName:  "device_ip",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "device_status",
			Comment:            `设备状态`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "DeviceStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "device_status",
			ProtobufFieldName:  "device_status",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "device_status_name",
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
			GoFieldName:        "DeviceStatusName",
			GoFieldType:        "string",
			JSONFieldName:      "device_status_name",
			ProtobufFieldName:  "device_status_name",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Aibox_event_info) TableName() string {
	return "v_aibox_event_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Aibox_event_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Aibox_event_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Aibox_event_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Aibox_event_info) TableInfo() *TableInfo {
	return Aibox_event_infoTableInfo
}
