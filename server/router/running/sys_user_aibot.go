package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUserAibotRouter struct {}

// InitSysUserAibotRouter 初始化 授权交易 路由信息
func (s *SysUserAibotRouter) InitSysUserAibotRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysUserAibotRouter := Router.Group("sysUserAibot").Use(middleware.OperationRecord())
	sysUserAibotRouterWithoutRecord := Router.Group("sysUserAibot")
	sysUserAibotRouterWithoutAuth := PublicRouter.Group("sysUserAibot")
	{
		sysUserAibotRouter.POST("createSysUserAibot", sysUserAibotApi.CreateSysUserAibot)   // 新建授权交易
		sysUserAibotRouter.DELETE("deleteSysUserAibot", sysUserAibotApi.DeleteSysUserAibot) // 删除授权交易
		sysUserAibotRouter.DELETE("deleteSysUserAibotByIds", sysUserAibotApi.DeleteSysUserAibotByIds) // 批量删除授权交易
		sysUserAibotRouter.PUT("updateSysUserAibot", sysUserAibotApi.UpdateSysUserAibot)    // 更新授权交易
	}
	{
		sysUserAibotRouterWithoutRecord.GET("findSysUserAibot", sysUserAibotApi.FindSysUserAibot)        // 根据ID获取授权交易
		sysUserAibotRouterWithoutRecord.GET("getSysUserAibotList", sysUserAibotApi.GetSysUserAibotList)  // 获取授权交易列表
	}
	{
	    sysUserAibotRouterWithoutAuth.GET("getSysUserAibotPublic", sysUserAibotApi.GetSysUserAibotPublic)  // 授权交易开放接口
	}
}
