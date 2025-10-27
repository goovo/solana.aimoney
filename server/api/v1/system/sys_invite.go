package system

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InviteApi struct{}

// 使用全局 userService 变量
// 注意：userService 是在 enter.go 中定义的全局变量

// GetInviteInfo
// @Tags     Invite
// @Summary  获取推荐信息
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {object} response.Response{data=systemRes.InviteInfoResponse,msg=string} "获取推荐信息"
// @Router   /invite/getInviteInfo [get]
func (i *InviteApi) GetInviteInfo(c *gin.Context) {
	userID := utils.GetUserID(c)

	// 获取用户信息
	user, err := userService.FindUserById(int(userID))
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	// 如果用户没有推荐码，生成一个唯一的邀请码
	if user.InviteCode == "" {
		uniqueInviteCode, err := generateUniqueInviteCode()
		if err != nil {
			global.GVA_LOG.Error("生成唯一推荐码失败!", zap.Error(err))
			response.FailWithMessage("生成推荐码失败，请稍后重试", c)
			return
		}
		user.InviteCode = uniqueInviteCode
		err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", userID).Update("invite_code", user.InviteCode).Error
		if err != nil {
			global.GVA_LOG.Error("更新推荐码失败!", zap.Error(err))
			response.FailWithMessage("生成推荐码失败", c)
			return
		}
	}

	// 获取总推荐人数
	totalReferral, err := userService.GetUserReferralCount(user.InviteCode)
	if err != nil {
		global.GVA_LOG.Error("获取推荐人数失败!", zap.Error(err))
		response.FailWithMessage("获取推荐人数失败", c)
		return
	}

	// 构建推荐链接
	inviteLink := ""
	if global.GVA_CONFIG.System.Addr != 0 {
		// 使用当前请求的主机和配置的端口构建完整URL
		protocol := "http"
		if c.Request.TLS != nil {
			protocol = "https"
		}
		host := c.Request.Host
		if host == "" {
			host = fmt.Sprintf("localhost:%d", global.GVA_CONFIG.System.Addr)
		}
		inviteLink = fmt.Sprintf("%s://%s/#/register?ref=%s", protocol, host, user.InviteCode)
	}

	res := systemRes.InviteInfoResponse{
		InviteCode:    user.InviteCode,
		InviteLink:    inviteLink,
		TotalReferral: totalReferral,
	}

	response.OkWithDetailed(res, "获取成功", c)
}

// GetMyReferrals
// @Tags     Invite
// @Summary  获取我的推荐列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data  query     systemReq.MyReferralsRequest                                true  "分页参数"
// @Success  200   {object}  response.Response{data=systemRes.MyReferralsResponse,msg=string} "获取推荐列表"
// @Router   /invite/getMyReferrals [get]
func (i *InviteApi) GetMyReferrals(c *gin.Context) {
	var req systemReq.MyReferralsRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	userID := utils.GetUserID(c)

	// 获取用户信息
	user, err := userService.FindUserById(int(userID))
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	// 获取推荐列表
	list, total, err := userService.GetUserReferrals(user.InviteCode, req.Page, req.PageSize, req.Keyword)
	if err != nil {
		global.GVA_LOG.Error("获取推荐列表失败!", zap.Error(err))
		response.FailWithMessage("获取推荐列表失败", c)
		return
	}

	res := systemRes.MyReferralsResponse{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	response.OkWithDetailed(res, "获取成功", c)
}

// CheckReferrerCode
// @Tags     Base
// @Summary  检查推荐码有效性
// @Produce  application/json
// @Param    code  query     string  true  "推荐码"
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string} "检查结果"
// @Router   /base/checkReferrerCode [get]
func (i *InviteApi) CheckReferrerCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("推荐码不能为空", c)
		return
	}

	exists, err := userService.CheckInviteCodeExists(code)
	if err != nil {
		global.GVA_LOG.Error("检查推荐码失败!", zap.Error(err))
		response.FailWithMessage("检查推荐码失败", c)
		return
	}

	if !exists {
		response.FailWithMessage("推荐码无效", c)
		return
	}

	response.OkWithData(gin.H{"valid": true}, c)
}
