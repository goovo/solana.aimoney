package system

import (
	"github.com/gin-gonic/gin"
)

type CryptoRouter struct{}

func (s *CryptoRouter) InitCryptoRouter(Router *gin.RouterGroup) {
	cryptoRouter := Router.Group("user")
	{
		cryptoRouter.GET("crypto/getCryptoData", cryptoApi.GetCryptoData)    // 获取加密货币数据
	}
}