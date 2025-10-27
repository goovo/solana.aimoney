
// 自动生成模板SysUserApi
package running
import (
	"time"
)

// 用户APIs 结构体  SysUserApi
type SysUserApi struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  UserId  *int `json:"userId" form:"userId" gorm:"comment:用户UID;column:userId;size:20;"`  //用户UID
  Exchange  *string `json:"exchange" form:"exchange" gorm:"comment:交易所名称;column:exchange;size:25;"`  //交易所名称
  Key  *string `json:"key" form:"key" gorm:"comment:api key;column:key;size:120;"`  //api key
  Secret  *string `json:"secret" form:"secret" gorm:"comment:api Secret;column:secret;size:120;"`  //api Secret
  Passwd  *string `json:"passwd" form:"passwd" gorm:"comment:api密码，OKx等交易所有设置;column:passwd;size:20;"`  //api密码，OKx等交易所有设置
  Status  *int `json:"status" form:"status" gorm:"comment:api状态: 1正常，2无权限，3错误;column:status;size:10;"`  //api状态: 1正常，2无权限，3错误
}


// TableName 用户APIs SysUserApi自定义表名 sys_user_api
func (SysUserApi) TableName() string {
    return "sys_user_api"
}





