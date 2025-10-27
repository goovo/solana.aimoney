package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysWeb3Wallet Web3钱包绑定表
type SysWeb3Wallet struct {
	global.GVA_MODEL
	WalletAddress string     `json:"walletAddress" gorm:"type:varchar(255);uniqueIndex;comment:钱包地址"` // 钱包地址
	WalletType    string     `json:"walletType" gorm:"type:varchar(50);default:'phantom';comment:钱包类型"` // 钱包类型: phantom, metamask等
	UserId        uint       `json:"userId" gorm:"index;comment:关联的用户ID"` // 关联的用户ID
	User          SysUser    `json:"user" gorm:"foreignKey:UserId;comment:关联用户"` // 关联的用户
	LastLoginAt   *time.Time `json:"lastLoginAt" gorm:"comment:最后登录时间"` // 最后登录时间
}

// TableName Web3钱包表名
func (SysWeb3Wallet) TableName() string {
	return "sys_web3_wallets"
}

