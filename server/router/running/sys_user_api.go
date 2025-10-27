package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUserApiRouter struct {}

// InitSysUserApiRouter 初始化 用户APIs 路由信息
func (s *SysUserApiRouter) InitSysUserApiRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysUserApiRouter := Router.Group("sysUserApi").Use(middleware.OperationRecord())
	sysUserApiRouterWithoutRecord := Router.Group("sysUserApi")
	sysUserApiRouterWithoutAuth := PublicRouter.Group("sysUserApi")
	{
		sysUserApiRouter.POST("createSysUserApi", sysUserApiApi.CreateSysUserApi)   // 新建用户APIs
		sysUserApiRouter.DELETE("deleteSysUserApi", sysUserApiApi.DeleteSysUserApi) // 删除用户APIs
		sysUserApiRouter.DELETE("deleteSysUserApiByIds", sysUserApiApi.DeleteSysUserApiByIds) // 批量删除用户APIs
		sysUserApiRouter.PUT("updateSysUserApi", sysUserApiApi.UpdateSysUserApi)    // 更新用户APIs
	}
	{
		sysUserApiRouterWithoutRecord.GET("findSysUserApi", sysUserApiApi.FindSysUserApi)        // 根据ID获取用户APIs
		sysUserApiRouterWithoutRecord.GET("getSysUserApiList", sysUserApiApi.GetSysUserApiList)  // 获取用户APIs列表
		sysUserApiRouterWithoutRecord.GET("getUserApiList", sysUserApiApi.GetUserApiList) 
	}
	{
	    sysUserApiRouterWithoutAuth.GET("getSysUserApiPublic", sysUserApiApi.GetSysUserApiPublic)  // 用户APIs开放接口
	}
}
