package system

import (
	"time"
)

// SysCrypto 加密货币数据结构
type SysCrypto struct {
	ID                uint      `json:"id" gorm:"primarykey;comment:主键"`
	Rank              int       `json:"rank" gorm:"index;comment:排名"`
	Name              string    `json:"name" gorm:"type:varchar(100);comment:名称"`
	Symbol            string    `json:"symbol" gorm:"type:varchar(20);index;comment:符号"`
	Icon              string    `json:"icon" gorm:"type:varchar(500);comment:图标URL"`
	Price             float64   `json:"price" gorm:"type:decimal(20,8);comment:价格"`
	Change24h         float64   `json:"change24h" gorm:"type:decimal(10,4);comment:24小时涨跌幅"`
	MarketCap         float64   `json:"marketCap" gorm:"type:decimal(30,2);comment:市值"`
	Volume24h         float64   `json:"volume24h" gorm:"type:decimal(30,2);comment:24小时交易量"`
	CirculatingSupply float64   `json:"circulatingSupply" gorm:"type:decimal(30,2);comment:流通量"`
	CreatedAt         time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt         time.Time `json:"updatedAt" gorm:"comment:更新时间"`
}

func (SysCrypto) TableName() string {
	return "sys_cryptos"
}