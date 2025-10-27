package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	PublicGroup := Router.Group("base")
	{
		PublicGroup.POST("login", baseApi.Login)
		PublicGroup.POST("captcha", baseApi.Captcha)
		// 新增公开注册接口：无需鉴权，通过基础路由对外开放
		// 复用 BaseApi.Register 的实现，保持与 /user/admin_register 一致的注册逻辑
		PublicGroup.POST("register", baseApi.Register)
		PublicGroup.POST("sendEmailCode", baseApi.SendEmailCode) // 发送邮箱验证码，无需鉴权
		PublicGroup.POST("checkUsername", baseApi.CheckUsername) // 检查用户名是否已存在，无需鉴权
		PublicGroup.POST("sendPhoneCode", baseApi.SendPhoneCode) // 发送手机验证码，无需鉴权
		PublicGroup.POST("checkPhone", baseApi.CheckPhone) // 检查手机号是否已注册，无需鉴权
		PublicGroup.POST("checkEmail", baseApi.CheckEmail) // 检查邮箱是否已注册，无需鉴权
	}
	return PublicGroup
}
