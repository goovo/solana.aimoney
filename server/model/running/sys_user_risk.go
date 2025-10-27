// 自动生成模板SysUserRisk
package running

import (
	"time"
)

// 用户风险等级 结构体  SysUserRisk
type SysUserRisk struct {
	Id        *int       `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`    //id字段
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"` //createdAt字段
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"` //updatedAt字段
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"` //deletedAt字段
	// 将 userId 声明为唯一索引，确保同一用户只有一条风险记录，便于 ON DUPLICATE KEY UPDATE 生效
	UserId *int   `json:"userId" form:"userId" gorm:"comment:用户UID;column:userId;size:20;uniqueIndex:uniq_sys_user_risk_userId"` //用户UID（唯一）
	Risk   string `json:"risk" form:"risk" gorm:"default:low;comment:风险等级;column:risk;type:enum('low','medium','high');"`        //风险等级
}

// TableName 用户风险等级 SysUserRisk自定义表名 sys_user_risk
func (SysUserRisk) TableName() string {
	return "sys_user_risk"
}
