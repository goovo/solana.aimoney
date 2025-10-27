package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TradesRouter struct {}

// InitTradesRouter 初始化 交易报表(模拟盘) 路由信息
func (s *TradesRouter) InitTradesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tradesRouter := Router.Group("trades").Use(middleware.OperationRecord())
	tradesRouterWithoutRecord := Router.Group("trades")
	tradesRouterWithoutAuth := PublicRouter.Group("trades")
	{
		tradesRouter.POST("createDryRun", tradesApi.CreateDryRun)   // 创建机器人(模拟盘)
		tradesRouter.POST("createTrades", tradesApi.CreateTrades)   // 新建交易报表(模拟盘)
		tradesRouter.DELETE("deleteTrades", tradesApi.DeleteTrades) // 删除交易报表(模拟盘)
		tradesRouter.DELETE("deleteTradesByIds", tradesApi.DeleteTradesByIds) // 批量删除交易报表(模拟盘)
		tradesRouter.PUT("updateTrades", tradesApi.UpdateTrades)    // 更新交易报表(模拟盘)
	}
	{
		tradesRouterWithoutRecord.GET("findTrades", tradesApi.FindTrades)        // 根据ID获取交易报表(模拟盘)
		tradesRouterWithoutRecord.GET("getTradesList", tradesApi.GetTradesList)  // 获取交易报表(模拟盘)列表
		tradesRouterWithoutRecord.GET("getDryRun", tradesApi.GetDryRun)          // 查询当前用户 dry-run 进程PID（无则0）
	}
	{
	    tradesRouterWithoutAuth.GET("getTradesPublic", tradesApi.GetTradesPublic)  // 交易报表(模拟盘)开放接口
	}
}
