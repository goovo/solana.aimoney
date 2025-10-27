package system

import (
	"github.com/gin-gonic/gin"
)

type PhantomRouter struct{}

// InitPhantomRouter 初始化Phantom钱包路由（公开路由）
func (s *PhantomRouter) InitPhantomRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	phantomRouter := Router.Group("phantom")
	web3WalletApi := apiGroupApp.Web3WalletApi
	{
		phantomRouter.POST("getNonce", web3WalletApi.GetNonce)   // 获取Nonce
		phantomRouter.POST("login", web3WalletApi.PhantomLogin)  // Phantom登录
	}
	return phantomRouter
}

// InitWeb3WalletRouter 初始化Web3钱包管理路由（需要鉴权）
func (s *PhantomRouter) InitWeb3WalletRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	web3Router := Router.Group("web3")
	web3WalletApi := apiGroupApp.Web3WalletApi
	{
		web3Router.POST("bindWallet", web3WalletApi.BindWallet)       // 绑定钱包
		web3Router.DELETE("unbindWallet", web3WalletApi.UnbindWallet) // 解绑钱包
	}
	return web3Router
}

