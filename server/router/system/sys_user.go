package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("user")
	{
		userRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		
		userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
		userRouter.PUT("setSelfSetting", baseApi.SetSelfSetting)          // 用户界面配置
	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
		userRouterWithoutRecord.GET("getUser/:userId", baseApi.GetUser)  // with userId获取user信息
		userRouterWithoutRecord.GET("getRisk", baseApi.GetRisk)          // 获取投资风险等级
		userRouterWithoutRecord.GET("getStat", baseApi.GetStat)          // 获取个人统计信息（API数量、交易笔数、胜率、总收益）
		userRouterWithoutRecord.GET("getStep", baseApi.GetStep)          // 获取用户关键步骤时间（注册/风评/API审核/首单）
	}
	
	// 推荐相关路由
	{
		userRouterWithoutRecord.GET("invite/getInviteInfo", inviteApi.GetInviteInfo)    // 获取推荐信息
		userRouterWithoutRecord.GET("invite/getMyReferrals", inviteApi.GetMyReferrals)  // 获取我的推荐列表
		userRouterWithoutRecord.GET("base/checkReferrerCode", inviteApi.CheckReferrerCode) // 检查推荐码有效性
	}
	
}
