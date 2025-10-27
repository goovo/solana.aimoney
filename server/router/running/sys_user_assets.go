package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUserAssetsRouter struct {}

// InitSysUserAssetsRouter 初始化 用户资产 路由信息
func (s *SysUserAssetsRouter) InitSysUserAssetsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysUserAssetsRouter := Router.Group("sysUserAssets").Use(middleware.OperationRecord())
	sysUserAssetsRouterWithoutRecord := Router.Group("sysUserAssets")
	sysUserAssetsRouterWithoutAuth := PublicRouter.Group("sysUserAssets")
	{
		sysUserAssetsRouter.POST("createSysUserAssets", sysUserAssetsApi.CreateSysUserAssets)   // 新建用户资产
		sysUserAssetsRouter.DELETE("deleteSysUserAssets", sysUserAssetsApi.DeleteSysUserAssets) // 删除用户资产
		sysUserAssetsRouter.DELETE("deleteSysUserAssetsByIds", sysUserAssetsApi.DeleteSysUserAssetsByIds) // 批量删除用户资产
		sysUserAssetsRouter.PUT("updateSysUserAssets", sysUserAssetsApi.UpdateSysUserAssets)    // 更新用户资产
	}
	{
		sysUserAssetsRouterWithoutRecord.GET("findSysUserAssets", sysUserAssetsApi.FindSysUserAssets)        // 根据ID获取用户资产
		sysUserAssetsRouterWithoutRecord.GET("getSysUserAssetsList", sysUserAssetsApi.GetSysUserAssetsList)  // 获取用户资产列表
	}
	{
	    sysUserAssetsRouterWithoutAuth.GET("getSysUserAssetsPublic", sysUserAssetsApi.GetSysUserAssetsPublic)  // 用户资产开放接口
	}
}
