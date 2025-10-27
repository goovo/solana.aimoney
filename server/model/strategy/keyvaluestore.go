// 自动生成模板Keyvaluestore
package strategy

import (
	"time"
)

// keyvaluestore表 结构体  Keyvaluestore
type Keyvaluestore struct {
	Id            *int       `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`                   //id字段
	Key           *string    `json:"key" form:"key" gorm:"column:key;size:25;"`                           //key字段
	ValueType     *string    `json:"valueType" form:"valueType" gorm:"column:value_type;size:20;"`        //valueType字段
	StringValue   *string    `json:"stringValue" form:"stringValue" gorm:"column:string_value;size:255;"` //stringValue字段
	DatetimeValue *time.Time `json:"datetimeValue" form:"datetimeValue" gorm:"column:datetime_value;"`    //datetimeValue字段
	FloatValue    *float64   `json:"floatValue" form:"floatValue" gorm:"column:float_value;"`             //floatValue字段
	IntValue      *int       `json:"intValue" form:"intValue" gorm:"column:int_value;size:10;"`           //intValue字段

	Pid          int64  `gorm:"-" json:"pid"` // gorm:"-" 告诉 GORM 不处理
	ProcessName  string `gorm:"-" json:"processName"`
	ConfigFile   string `gorm:"-" json:"configFile"`
	StrategyName string `gorm:"-" json:"strategyName"`
	DryRun       string `gorm:"-" json:"dryRun"`
}

// TableName keyvaluestore表 Keyvaluestore自定义表名 keyvaluestore
func (Keyvaluestore) TableName() string {
	return "keyvaluestore"
}
