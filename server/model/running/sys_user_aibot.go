// 自动生成模板SysUserAibot
package running

import (
	"time"
)

// 授权交易 结构体  SysUserAibot
type SysUserAibot struct {
	UserId    *int       `json:"userId" form:"userId" gorm:"uniqueIndex;primarykey;comment:用户Id;column:userId;size:10;" binding:"required"`        //用户Id
	CreatedAt *time.Time `json:"createdAt" form:"createdAt" gorm:"comment:授权时间;column:created_at;"`                                                //授权时间
	UpdatedAt *time.Time `json:"updatedAt" form:"updatedAt" gorm:"comment:更新授权时间;column:updated_at;"`                                              //更新授权时间
	DeletedAt *time.Time `json:"deletedAt" form:"deletedAt" gorm:"comment:逻辑删除时间;column:deleted_at;"`                                              //逻辑删除时间
	AiBot     string     `json:"aiBot" form:"aiBot" gorm:"default:allow;comment:授权机器人;column:aiBot;type:enum('allow','deny');" binding:"required"` //授权机器人
}

// TableName 授权交易 SysUserAibot自定义表名 sys_user_aibot
func (SysUserAibot) TableName() string {
	return "sys_user_aibot"
}
