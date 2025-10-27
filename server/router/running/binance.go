package running

import (
	"github.com/gin-gonic/gin"
)

type BinanceRouter struct{}

func (s *BinanceRouter) InitBinanceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	binanceRouter := Router.Group("binance")
	{
		binanceRouter.GET("account/:userId", binanceApi.GetAccountBalance)
		binanceRouter.GET("apiRes/:userId", binanceApi.GetApiRes)
	}
}
