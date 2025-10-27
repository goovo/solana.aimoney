
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type SysUserAibotSearch struct{
      UserId  *int `json:"userId" form:"userId"` 
      AiBot  string `json:"aiBot" form:"aiBot"` 
    request.PageInfo
}
