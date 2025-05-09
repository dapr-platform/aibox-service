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


Table: v_aibox_update_info
[ 0] id                                             VARCHAR(36)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 36      default: []
[ 1] version                                        VARCHAR(64)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 64      default: []
[ 2] type                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] type_name                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 4] file_path                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] file_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] file_key                                       VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] description                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] status_name                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[11] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "eXZBHhsJuAnOVFGUVcYRcvBoi",    "version": "mVPTxWlXTyUxQGELyRfHkqxgs",    "type": 31,    "type_name": "qDlDbbgaOJBKcXWMiaWogDJhS",    "file_path": "nrXruxlhnZgrmBbYXCveQXAqJ",    "file_name": "NaGGJNDpAhgKAfHXxRUwuaJCW",    "file_key": "aPLPWtUZJxpTCBRSCvouWxdlA",    "description": "ruAkpQtLuYTZsWRHcgwaWadgN",    "status": 33,    "status_name": "SsAlVtRYdbPHbVaaWtiOwWeWC",    "created_time": 86,    "updated_time": 11}


Comments
-------------------------------------
[ 0] Warning table: v_aibox_update_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_update_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Aibox_update_info_FIELD_NAME_id = "id"

	Aibox_update_info_FIELD_NAME_version = "version"

	Aibox_update_info_FIELD_NAME_type = "type"

	Aibox_update_info_FIELD_NAME_type_name = "type_name"

	Aibox_update_info_FIELD_NAME_file_path = "file_path"

	Aibox_update_info_FIELD_NAME_file_name = "file_name"

	Aibox_update_info_FIELD_NAME_file_key = "file_key"

	Aibox_update_info_FIELD_NAME_description = "description"

	Aibox_update_info_FIELD_NAME_status = "status"

	Aibox_update_info_FIELD_NAME_status_name = "status_name"

	Aibox_update_info_FIELD_NAME_created_time = "created_time"

	Aibox_update_info_FIELD_NAME_updated_time = "updated_time"
)

// Aibox_update_info struct is a row record of the v_aibox_update_info table in the  database
type Aibox_update_info struct {
	ID string `json:"id"` //更新ID

	Version string `json:"version"` //版本号

	Type int32 `json:"type"` //更新类型

	TypeName string `json:"type_name"` //更新类型名称

	FilePath string `json:"file_path"` //文件存放路径

	FileName string `json:"file_name"` //文件名

	FileKey string `json:"file_key"` //文件key

	Description string `json:"description"` //更新描述

	Status int32 `json:"status"` //状态

	StatusName string `json:"status_name"` //状态名称

	CreatedTime common.LocalTime `json:"created_time"` //创建时间

	UpdatedTime common.LocalTime `json:"updated_time"` //更新时间

}

var Aibox_update_infoTableInfo = &TableInfo{
	Name: "v_aibox_update_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `更新ID`,
			Notes: `Warning table: v_aibox_update_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_aibox_update_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "version",
			Comment:            `版本号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       64,
			GoFieldName:        "Version",
			GoFieldType:        "string",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
			Comment:            `更新类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type_name",
			Comment:            `更新类型名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "TypeName",
			GoFieldType:        "string",
			JSONFieldName:      "type_name",
			ProtobufFieldName:  "type_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "file_path",
			Comment:            `文件存放路径`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "FilePath",
			GoFieldType:        "string",
			JSONFieldName:      "file_path",
			ProtobufFieldName:  "file_path",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "file_name",
			Comment:            `文件名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "FileName",
			GoFieldType:        "string",
			JSONFieldName:      "file_name",
			ProtobufFieldName:  "file_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "file_key",
			Comment:            `文件key`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "FileKey",
			GoFieldType:        "string",
			JSONFieldName:      "file_key",
			ProtobufFieldName:  "file_key",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "description",
			Comment:            `更新描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
			Comment:            `状态`,
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
			Comment:            `状态名称`,
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
	},
}

// TableName sets the insert table name for this struct type
func (a *Aibox_update_info) TableName() string {
	return "v_aibox_update_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Aibox_update_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Aibox_update_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Aibox_update_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Aibox_update_info) TableInfo() *TableInfo {
	return Aibox_update_infoTableInfo
}
