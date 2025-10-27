package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FreqStrategy struct {
	global.GVA_MODEL
	Name       string `json:"Name" gorm:"uniqueIndex;comment:策略名"` // 策略名称
	FileName   string `json:"FileName"  gorm:"comment:文件名"`        // 策略文件名
	Status     string `json:"Status" gorm:"default:true;comment:状态"`
	Hyperopt   bool   `json:"Hyperopt" gorm:"default:false;comment:是否支持超参优化"`
	BuyParams  uint   `json:"BuyParams" gorm:"default:0;comment:Buy参数个数"`
	SellParams uint   `json:"SellParams" gorm:"default:0;comment:Sell参数个数"`
	AI         bool   `json:"AI" gorm:"default:false;comment:是否支持AI优化"`
	TimeFrame  string `json:"TimeFrame"  gorm:"comment:K线周期"`        // 正则提取: 30s/1m/5m/15m/1h/4h
	Direction  string `json:"Direction"  gorm:"comment:方向 more/less"` // 用户邮箱
}

func (FreqStrategy) TableName() string {
	return "sys_freq_strategy"
}
