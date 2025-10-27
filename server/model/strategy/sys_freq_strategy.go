
// 自动生成模板SysFreqStrategy
package strategy
import (
	"time"
)

// sysFreqStrategy表 结构体  SysFreqStrategy
type SysFreqStrategy struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  Name  *string `json:"name" form:"name" gorm:"comment:策略名;column:name;size:191;"`  //策略名
  FileName  *string `json:"fileName" form:"fileName" gorm:"comment:文件名;column:file_name;size:191;"`  //文件名
  Status  *string `json:"status" form:"status" gorm:"comment:状态;column:status;size:191;"`  //状态
  Hyperopt  *bool `json:"hyperopt" form:"hyperopt" gorm:"comment:是否支持超参优化;column:hyperopt;"`  //是否支持超参优化
  BuyParams  *int `json:"buyParams" form:"buyParams" gorm:"comment:Buy参数个数;column:buy_params;size:20;"`  //Buy参数个数
  SellParams  *int `json:"sellParams" form:"sellParams" gorm:"comment:Sell参数个数;column:sell_params;size:20;"`  //Sell参数个数
  Ai  *bool `json:"ai" form:"ai" gorm:"comment:是否支持AI优化;column:ai;"`  //是否支持AI优化
  TimeFrame  *string `json:"timeFrame" form:"timeFrame" gorm:"comment:K线周期;column:time_frame;size:191;"`  //K线周期
  Direction  *string `json:"direction" form:"direction" gorm:"comment:方向 more/less;column:direction;size:191;"`  //方向 more/less
}


// TableName sysFreqStrategy表 SysFreqStrategy自定义表名 sys_freq_strategy
func (SysFreqStrategy) TableName() string {
    return "sys_freq_strategy"
}





