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


Table: v_aibox_active_event_stats
[ 0] level                                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] level_name                                     TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] event_count                                    INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "level": 82,    "level_name": "mOvMBbVYbqJWQJWZkRjgyhvAb",    "event_count": 73}


Comments
-------------------------------------
[ 0] Warning table: v_aibox_active_event_stats does not have a primary key defined, setting col position 1 level as primary key
Warning table: v_aibox_active_event_stats primary key column level is nullable column, setting it as NOT NULL




*/

var (
	Aibox_active_event_stats_FIELD_NAME_level = "level"

	Aibox_active_event_stats_FIELD_NAME_level_name = "level_name"

	Aibox_active_event_stats_FIELD_NAME_event_count = "event_count"
)

// Aibox_active_event_stats struct is a row record of the v_aibox_active_event_stats table in the  database
type Aibox_active_event_stats struct {
	Level int32 `json:"level"` //事件级别

	LevelName string `json:"level_name"` //事件级别名称

	EventCount int32 `json:"event_count"` //事件数量

}

var Aibox_active_event_statsTableInfo = &TableInfo{
	Name: "v_aibox_active_event_stats",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "level",
			Comment: `事件级别`,
			Notes: `Warning table: v_aibox_active_event_stats does not have a primary key defined, setting col position 1 level as primary key
Warning table: v_aibox_active_event_stats primary key column level is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Level",
			GoFieldType:        "int32",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "event_count",
			Comment:            `事件数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "EventCount",
			GoFieldType:        "int32",
			JSONFieldName:      "event_count",
			ProtobufFieldName:  "event_count",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Aibox_active_event_stats) TableName() string {
	return "v_aibox_active_event_stats"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Aibox_active_event_stats) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Aibox_active_event_stats) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Aibox_active_event_stats) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Aibox_active_event_stats) TableInfo() *TableInfo {
	return Aibox_active_event_statsTableInfo
}
