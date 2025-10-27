package strategy

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type KeyvaluestoreRouter struct {}

// InitKeyvaluestoreRouter 初始化 keyvaluestore表 路由信息
func (s *KeyvaluestoreRouter) InitKeyvaluestoreRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	keyvaluestoreRouter := Router.Group("keyvaluestore").Use(middleware.OperationRecord())
	keyvaluestoreRouterWithoutRecord := Router.Group("keyvaluestore")
	keyvaluestoreRouterWithoutAuth := PublicRouter.Group("keyvaluestore")
	{
		keyvaluestoreRouter.POST("createKeyvaluestore", keyvaluestoreApi.CreateKeyvaluestore)   // 新建keyvaluestore表
		keyvaluestoreRouter.DELETE("deleteKeyvaluestore", keyvaluestoreApi.DeleteKeyvaluestore) // 删除keyvaluestore表
		keyvaluestoreRouter.DELETE("deleteKeyvaluestoreByIds", keyvaluestoreApi.DeleteKeyvaluestoreByIds) // 批量删除keyvaluestore表
		keyvaluestoreRouter.PUT("updateKeyvaluestore", keyvaluestoreApi.UpdateKeyvaluestore)    // 更新keyvaluestore表
	}
	{
		keyvaluestoreRouterWithoutRecord.GET("findKeyvaluestore", keyvaluestoreApi.FindKeyvaluestore)        // 根据ID获取keyvaluestore表
		keyvaluestoreRouterWithoutRecord.GET("getKeyvaluestoreList", keyvaluestoreApi.GetKeyvaluestoreList)  // 获取keyvaluestore表列表
	}
	{
	    keyvaluestoreRouterWithoutAuth.GET("getKeyvaluestorePublic", keyvaluestoreApi.GetKeyvaluestorePublic)  // keyvaluestore表开放接口
	}
}
