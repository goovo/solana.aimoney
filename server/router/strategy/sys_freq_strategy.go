package strategy

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysFreqStrategyRouter struct {}

// InitSysFreqStrategyRouter 初始化 sysFreqStrategy表 路由信息
func (s *SysFreqStrategyRouter) InitSysFreqStrategyRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysFreqStrategyRouter := Router.Group("sysFreqStrategy").Use(middleware.OperationRecord())
	sysFreqStrategyRouterWithoutRecord := Router.Group("sysFreqStrategy")
	sysFreqStrategyRouterWithoutAuth := PublicRouter.Group("sysFreqStrategy")
	{
		sysFreqStrategyRouter.POST("createSysFreqStrategy", sysFreqStrategyApi.CreateSysFreqStrategy)   // 新建sysFreqStrategy表
		sysFreqStrategyRouter.DELETE("deleteSysFreqStrategy", sysFreqStrategyApi.DeleteSysFreqStrategy) // 删除sysFreqStrategy表
		sysFreqStrategyRouter.DELETE("deleteSysFreqStrategyByIds", sysFreqStrategyApi.DeleteSysFreqStrategyByIds) // 批量删除sysFreqStrategy表
		sysFreqStrategyRouter.PUT("updateSysFreqStrategy", sysFreqStrategyApi.UpdateSysFreqStrategy)    // 更新sysFreqStrategy表
	}
	{
		sysFreqStrategyRouterWithoutRecord.GET("findSysFreqStrategy", sysFreqStrategyApi.FindSysFreqStrategy)        // 根据ID获取sysFreqStrategy表
		sysFreqStrategyRouterWithoutRecord.GET("getSysFreqStrategyList", sysFreqStrategyApi.GetSysFreqStrategyList)  // 获取sysFreqStrategy表列表
	}
	{
	    sysFreqStrategyRouterWithoutAuth.GET("getSysFreqStrategyPublic", sysFreqStrategyApi.GetSysFreqStrategyPublic)  // sysFreqStrategy表开放接口
	}
}
