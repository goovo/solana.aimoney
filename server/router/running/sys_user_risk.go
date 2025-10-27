package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUserRiskRouter struct {}

// InitSysUserRiskRouter 初始化 用户风险等级 路由信息
func (s *SysUserRiskRouter) InitSysUserRiskRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysUserRiskRouter := Router.Group("sysUserRisk").Use(middleware.OperationRecord())
	sysUserRiskRouterWithoutRecord := Router.Group("sysUserRisk")
	sysUserRiskRouterWithoutAuth := PublicRouter.Group("sysUserRisk")
	{
		sysUserRiskRouter.POST("setUserRisk", sysUserRiskApi.SetUserRisk)               // 设置用户风险等级
		sysUserRiskRouter.POST("createSysUserRisk", sysUserRiskApi.CreateSysUserRisk)   // 新建用户风险等级
		sysUserRiskRouter.DELETE("deleteSysUserRisk", sysUserRiskApi.DeleteSysUserRisk) // 删除用户风险等级
		sysUserRiskRouter.DELETE("deleteSysUserRiskByIds", sysUserRiskApi.DeleteSysUserRiskByIds) // 批量删除用户风险等级
		sysUserRiskRouter.PUT("updateSysUserRisk", sysUserRiskApi.UpdateSysUserRisk)    // 更新用户风险等级
	}
	{
		sysUserRiskRouterWithoutRecord.GET("findSysUserRisk", sysUserRiskApi.FindSysUserRisk)        // 根据ID获取用户风险等级
		sysUserRiskRouterWithoutRecord.GET("getSysUserRiskList", sysUserRiskApi.GetSysUserRiskList)  // 获取用户风险等级列表
	}
	{
	    sysUserRiskRouterWithoutAuth.GET("getSysUserRiskPublic", sysUserRiskApi.GetSysUserRiskPublic)  // 用户风险等级开放接口
	}
}
