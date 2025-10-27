package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Web3WalletApi struct{}

var web3WalletService = service.ServiceGroupApp.SystemServiceGroup.Web3WalletService

// GetNonce
// @Tags     Phantom
// @Summary  获取Nonce用于签名
// @Produce  application/json
// @Param    data  body      systemReq.PhantomNonceRequest                                true  "钱包地址"
// @Success  200   {object}  response.Response{data=systemRes.PhantomNonceResponse,msg=string}  "返回Nonce和签名消息"
// @Router   /phantom/getNonce [post]
func (w *Web3WalletApi) GetNonce(c *gin.Context) {
	var req systemReq.PhantomNonceRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 检查Redis是否初始化
	if global.GVA_REDIS == nil {
		global.GVA_LOG.Error("Redis未初始化，请检查config.yaml中use-redis配置")
		response.FailWithMessage("Redis服务未启用，请联系管理员配置", c)
		return
	}

	// 生成Nonce
	nonce := web3WalletService.GenerateNonce()

	// 构造需要签名的消息
	message := "欢迎登录 AiMoney.run！\n\n请签名此消息以验证您的身份。\n\nNonce: " + nonce + "\nTimestamp: " + time.Now().Format("2006-01-02 15:04:05")

	// 将Nonce存储到Redis（5分钟有效期）
	err = global.GVA_REDIS.Set(c.Request.Context(), "phantom:nonce:"+req.WalletAddress, nonce, 5*time.Minute).Err()
	if err != nil {
		global.GVA_LOG.Error("存储Nonce失败!", zap.Error(err))
		response.FailWithMessage("生成Nonce失败", c)
		return
	}

	response.OkWithDetailed(systemRes.PhantomNonceResponse{
		Nonce:   nonce,
		Message: message,
	}, "获取Nonce成功", c)
}

// PhantomLogin
// @Tags     Phantom
// @Summary  Phantom钱包登录
// @Produce  application/json
// @Param    data  body      systemReq.PhantomLoginRequest                                true  "钱包地址、签名、消息"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /phantom/login [post]
func (w *Web3WalletApi) PhantomLogin(c *gin.Context) {
	var req systemReq.PhantomLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证签名并获取/创建用户
	user, err := web3WalletService.VerifySignatureAndLogin(req.WalletAddress, req.Message, req.Signature)
	if err != nil {
		global.GVA_LOG.Error("Phantom登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	if user.Enable != 1 {
		global.GVA_LOG.Error("登录失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}

	// 签发JWT Token（复用系统的TokenNext逻辑）
	w.TokenNext(c, *user)
}

// TokenNext 登录以后签发jwt（复用BaseApi的TokenNext逻辑）
func (w *Web3WalletApi) TokenNext(c *gin.Context, user system.SysUser) {
	token, claims, err := utils.LoginToken(&user)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	if !global.GVA_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := utils.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := utils.SetRedisJWT(token, user.GetUsername()); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// BindWallet
// @Tags      Web3Wallet
// @Summary   绑定钱包到当前用户
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      systemReq.PhantomLoginRequest                true  "钱包地址、签名、消息"
// @Success   200   {object}  response.Response{msg=string}                "绑定成功"
// @Router    /web3/bindWallet [post]
func (w *Web3WalletApi) BindWallet(c *gin.Context) {
	var req systemReq.PhantomLoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 验证签名
	valid, err := web3WalletService.VerifyPhantomSignature(req.WalletAddress, req.Message, req.Signature)
	if err != nil || !valid {
		response.FailWithMessage("签名验证失败", c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	// 绑定钱包
	err = web3WalletService.BindWalletToUser(req.WalletAddress, userId)
	if err != nil {
		global.GVA_LOG.Error("绑定钱包失败!", zap.Error(err))
		response.FailWithMessage("绑定失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("绑定成功", c)
}

// UnbindWallet
// @Tags      Web3Wallet
// @Summary   解绑钱包
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     walletAddress  query     string  true  "钱包地址"
// @Success   200   {object}  response.Response{msg=string}  "解绑成功"
// @Router    /web3/unbindWallet [delete]
func (w *Web3WalletApi) UnbindWallet(c *gin.Context) {
	walletAddress := c.Query("walletAddress")
	if walletAddress == "" {
		response.FailWithMessage("钱包地址不能为空", c)
		return
	}

	// 获取当前用户ID
	userId := utils.GetUserID(c)

	// 解绑钱包
	err := web3WalletService.UnbindWallet(walletAddress, userId)
	if err != nil {
		global.GVA_LOG.Error("解绑钱包失败!", zap.Error(err))
		response.FailWithMessage("解绑失败", c)
		return
	}

	response.OkWithMessage("解绑成功", c)
}

