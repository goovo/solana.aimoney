
// 自动生成模板SysUserAssets
package running
import (
	"time"
)

// 用户资产 结构体  SysUserAssets
type SysUserAssets struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
  UserId  *int `json:"userId" form:"userId" gorm:"comment:用户UID;column:userId;size:20;"`  //用户UID
  Exchange  *string `json:"exchange" form:"exchange" gorm:"comment:交易所名称;column:exchange;size:25;"`  //交易所名称
  ApiId  *int `json:"apiId" form:"apiId" gorm:"comment:对应sys_user_api.id;column:api_id;size:10;"`  //对应sys_user_api.id
  Assets  *float64 `json:"assets" form:"assets" gorm:"comment:api资产;column:assets;"`  //api资产
}


// TableName 用户资产 SysUserAssets自定义表名 sys_user_assets
func (SysUserAssets) TableName() string {
    return "sys_user_assets"
}





