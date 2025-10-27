package system

import (
	"fmt"
	mathrand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	// 新增：引入 running 包，访问 sys_user_api、trades 模型
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	key := c.ClientIP()
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc && !store.Verify(l.CaptchaId, l.Captcha, true) {
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("验证码错误", c)
		return
	}

	u := &system.SysUser{Username: l.Username, Password: l.Password}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	}
	if user.Enable != 1 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		return
	}
	b.TokenNext(c, *user)
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
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

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "密码, 角色ID, 手机/邮箱验证码"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 如果是公开注册路由，则需要校验验证码，避免机器人滥用
	if c.FullPath() == "/base/register" {
		// 优先校验手机号验证码（默认使用手机号注册）
		if r.Phone != "" && r.PhoneCode != "" {
			// 校验手机验证码
			if !VerifyPhoneCode(r.Phone, r.PhoneCode) {
				response.FailWithMessage("手机验证码错误或已过期", c)
				return
			}
		} else if r.Email != "" && r.EmailCode != "" {
			// 备选：校验邮箱验证码
			if !VerifyEmailCode(r.Email, r.EmailCode) {
				response.FailWithMessage("邮箱验证码错误或已过期", c)
				return
			}
		} else {
			// 两种验证码都未提供时的错误提示
			response.FailWithMessage("请提供手机验证码或邮箱验证码进行校验", c)
			return
		}
	}

	var authorities []system.SysAuthority
	for _, v := range r.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}

	// 自动生成用户名：优先使用手机号，其次是邮箱
	username := generateUsername(r.Phone, r.Email)

	// 生成唯一的邀请码
	inviteCode, err := generateUniqueInviteCode()
	if err != nil {
		global.GVA_LOG.Error("生成邀请码失败!", zap.Error(err))
		response.FailWithMessage("生成邀请码失败", c)
		return
	}

	user := &system.SysUser{
		Username:     username,
		NickName:     generateNickname(r.Phone, r.Email),
		Password:     r.Password,
		HeaderImg:    r.HeaderImg,
		AuthorityId:  r.AuthorityId,
		Authorities:  authorities,
		Enable:       r.Enable,
		Phone:        r.Phone,
		Email:        r.Email,
		InviteCode:   inviteCode,
		ReferrerCode: r.ReferrerCode,
	}

	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// generateUsername 根据手机号或邮箱生成用户名
func generateUsername(phone, email string) string {
	if phone != "" {
		// 使用手机号后8位作为用户名
		if len(phone) >= 8 {
			return "u_" + phone[len(phone)-8:]
		}
		return "u_" + phone
	}

	if email != "" {
		// 使用邮箱@前的部分作为用户名
		parts := strings.Split(email, "@")
		if len(parts) > 0 {
			return "u_" + parts[0]
		}
		return "u_" + email
	}

	// 生成随机用户名
	return fmt.Sprintf("user_%d", time.Now().Unix())
}

// generateNickname 生成昵称
func generateNickname(phone, email string) string {
	if phone != "" {
		// 使用手机号中间4位替换为*作为昵称
		if len(phone) == 11 {
			return phone[:3] + "****" + phone[7:]
		}
		return phone
	}

	if email != "" {
		// 使用邮箱@前的部分作为昵称
		parts := strings.Split(email, "@")
		if len(parts) > 0 {
			return parts[0]
		}
		return email
	}

	return "新用户"
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo systemReq.GetUserList
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// SetUserAuthority
// @Tags      SysUser
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthority [post]
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	claims.AuthorityId = sua.AuthorityId
	token, err := utils.NewJWT().CreateToken(*claims)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Header("new-token", token)
	c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
	utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	response.OkWithMessage("修改成功", c)
}

// SetUserAuthorities
// @Tags      SysUser
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /user/setUserAuthorities [post]
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authorityID := utils.GetUserAuthorityId(c)
	err = userService.SetUserAuthorities(authorityID, sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// DeleteUser
// @Tags      SysUser
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 无法删除自己。", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(user.AuthorityIds) != 0 {
		authorityID := utils.GetUserAuthorityId(c)
		err = userService.SetUserAuthorities(authorityID, user.ID, user.AuthorityIds)
		if err != nil {
			global.GVA_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /user/SetSelfInfo [put]
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		GVA_MODEL: global.GVA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfSetting
// @Tags      SysUser
// @Summary   设置用户配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      map[string]interface{}  true  "用户配置数据"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户配置"
// @Router    /user/SetSelfSetting [put]
func (b *BaseApi) SetSelfSetting(c *gin.Context) {
	var req common.JSONMap
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = userService.SetSelfSetting(req, utils.GetUserID(c))
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}

// GetUser
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /user/getUser [get]
func (b *BaseApi) GetUser(c *gin.Context) {
	// 1. 参数校验
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		response.FailWithMessage("用户ID不能为空 or 转换失败", c)
		return
	}

	ReqUser, err := userService.FindUserById(userID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}

// GetRisk
// @Tags      SysUser
// @Summary   通过 userId 获取最新的投资风险等级
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     userId  path      int  true  "用户ID"
// @Success   200 {object} response.Response{data=map[string]interface{},msg=string} "查询成功"
// @Router    /user/getRisk/{userId} [get]
func (b *BaseApi) GetRisk(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	// 中文注释：从路径参数获取 userId 并进行合法性校验
	userID := int(ReqUser.ID)
	if userID <= 0 {
		response.FailWithMessage("用户ID无效", c)
		return
	}
	// 中文注释：调用 running 模块的服务层方法，按 userId 查询最新一条风险等级记录
	record, err := service.ServiceGroupApp.RunningServiceGroup.SysUserRiskService.GetSysUserRiskByUserId(c.Request.Context(), userID)
	if err != nil {
		global.GVA_LOG.Error("查询用户风险等级失败", zap.Error(err))
		response.FailWithMessage("未查询到风险等级", c)
		return
	}

	// 中文注释：仅返回必要字段，避免泄露不必要信息
	response.OkWithDetailed(gin.H{
		"userId":    userID,
		"risk":      record.Risk,
		"updatedAt": record.UpdatedAt, // 中文注释：更新时间，可能为 null
	}, "查询成功", c)
}

// ResetPassword
// @Tags      SysUser
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      system.SysUser                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /user/resetPassword [post]
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var rps systemReq.ResetPassword
	err := c.ShouldBindJSON(&rps)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(rps.ID, rps.Password)
	if err != nil {
		global.GVA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

func init() {
	// 初始化随机数生成器
	mathrand.Seed(time.Now().UnixNano())
}

// generateInviteCode 生成邀请码
func generateInviteCode() string {
	// 定义字符集：24个字母（大小写）+ 10个数字 = 62个字符
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const codeLength = 6

	// 生成6位随机邀请码
	code := make([]byte, codeLength)
	for i := 0; i < codeLength; i++ {
		// 从字符集中随机选择一个字符
		code[i] = charset[mathrand.Intn(len(charset))]
	}

	return string(code)
}

// generateUniqueInviteCode 生成唯一的邀请码
func generateUniqueInviteCode() (string, error) {
	const maxRetries = 10 // 最大重试次数，避免无限循环
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const codeLength = 6

	for i := 0; i < maxRetries; i++ {
		// 生成邀请码
		code := make([]byte, codeLength)
		for j := 0; j < codeLength; j++ {
			code[j] = charset[mathrand.Intn(len(charset))]
		}

		inviteCode := string(code)

		// 检查是否已存在
		exists, err := userService.CheckInviteCodeExists(inviteCode)
		if err != nil {
			return "", err
		}

		// 如果不存在，返回该邀请码
		if !exists {
			return inviteCode, nil
		}

		// 如果存在，继续下一次循环生成新的邀请码
	}

	// 如果重试次数用完仍未生成唯一邀请码，返回错误
	return "", fmt.Errorf("无法生成唯一的邀请码，已达到最大重试次数")
}

// GetStat
// @Tags     User
// @Summary  获取当前登录用户交易统计
// @Produce  application/json
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string}  "返回 user: { apiCnt, tradeCnt, getCnt, getTotal }"
// @Router   /user/getStat [get]
func (b *BaseApi) GetStat(c *gin.Context) {
	// 从 JWT 中获取当前登录用户ID
	uid := utils.GetUserID(c)

	// 统计1：当前用户开通API数量（status=1）
	var apiCnt int64
	if err := global.GVA_DB.Model(&running.SysUserApi{}).Where("userId = ? AND status = 1", uid).Count(&apiCnt).Error; err != nil {
		global.GVA_LOG.Error("统计用户API数量失败", zap.Error(err), zap.Uint("uid", uid))
		response.FailWithMessage("统计用户API数量失败: "+err.Error(), c)
		return
	}

	// 统计2：交易笔数（freq 库 trades，已平仓 is_open=0）
	var tradeCnt int64
	if err := global.GetGlobalDBByDBName("freq").Model(&running.Trades{}).Where("user_id = ? AND is_open = false", uid).Count(&tradeCnt).Error; err != nil {
		global.GVA_LOG.Error("统计交易笔数失败", zap.Error(err), zap.Uint("uid", uid))
		response.FailWithMessage("统计交易笔数失败: "+err.Error(), c)
		return
	}

	// 统计3：赚钱交易笔数（realized_profit > 0）
	var getCnt int64
	if err := global.GetGlobalDBByDBName("freq").Model(&running.Trades{}).Where("user_id = ? AND is_open = false AND realized_profit > 0", uid).Count(&getCnt).Error; err != nil {
		global.GVA_LOG.Error("统计盈利交易笔数失败", zap.Error(err), zap.Uint("uid", uid))
		response.FailWithMessage("统计盈利交易笔数失败: "+err.Error(), c)
		return
	}

	// 统计4：总收益（sum(realized_profit)，空值按0处理）
	var getTotal float64
	if err := global.GetGlobalDBByDBName("freq").Model(&running.Trades{}).Select("COALESCE(SUM(realized_profit), 0)").Where("user_id = ? AND is_open = false", uid).Scan(&getTotal).Error; err != nil {
		global.GVA_LOG.Error("统计总收益失败", zap.Error(err), zap.Uint("uid", uid))
		response.FailWithMessage("统计总收益失败: "+err.Error(), c)
		return
	}

	// 返回统一结构，前端可用 user.apiCnt/user.tradeCnt/user.getCnt/user.getTotal
	data := gin.H{
		"user": gin.H{
			"apiCnt":   apiCnt,
			"tradeCnt": tradeCnt,
			"getCnt":   getCnt,
			"getTotal": getTotal,
		},
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetStep
// @Tags     User
// @Summary  获取当前登录用户的关键步骤时间（注册/风评/API审核/首单完成）
// @Produce  application/json
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string}  "返回 user: { registerTime, riskTime, apiTime, tradeTime }"
// @Router   /user/getStep [get]
func (b *BaseApi) GetStep(c *gin.Context) {
	// 中文注释：从 JWT 中获取当前登录用户ID
	uid := utils.GetUserID(c)

	// 工具函数：格式化时间为 YYYY-MM-DD（指针安全，兼容 *time.Time 和 nil）
	format := func(t *time.Time) string {
		if t == nil || t.IsZero() {
			return ""
		}
		return t.Format("2006-01-02")
	}

	// 1. 注册时间：优先使用用户模型的 CreatedAt（一定存在），保证返回 YYYY-MM-DD
	var user system.SysUser
	var registerTime string
	if err := global.GVA_DB.Select("created_at").Where("id = ?", uid).First(&user).Error; err == nil {
		registerTime = format(&user.CreatedAt) // 兼容指针签名
	} else {
		// 兜底，出现异常则置空并仅记录日志
		global.GVA_LOG.Warn("查询注册时间失败", zap.Error(err), zap.Uint("uid", uid))
	}

	// 2. 风险等级评估时间：从 sys_user_risk 表取最新一条记录的 created_at
	var risk running.SysUserRisk
	var riskTime string
	if err := global.GVA_DB.Select("created_at").Where("userId = ?", uid).Order("id desc").First(&risk).Error; err == nil {
		riskTime = format(risk.CreatedAt) // risk.CreatedAt 为 *time.Time
	} else {
		if err.Error() != "record not found" {
			global.GVA_LOG.Warn("查询风险评估时间失败", zap.Error(err), zap.Uint("uid", uid))
		}
	}

	// 3. API 审核时间：从 sys_user_api 表筛选 status=1，取最新一条的 updated_at
	var api running.SysUserApi
	var apiTime string
	if err := global.GVA_DB.Select("updated_at").Where("userId = ? AND status = 1", uid).Order("updated_at desc").First(&api).Error; err == nil {
		apiTime = format(api.UpdatedAt) // api.UpdatedAt 为 *time.Time
	} else {
		if err.Error() != "record not found" {
			global.GVA_LOG.Warn("查询API审核时间失败", zap.Error(err), zap.Uint("uid", uid))
		}
	}

	// 4. 第一笔交易完成时间：freq.trades 表 is_open=0 按 id 升序取第一条的 close_date
	var trade running.Trades
	var tradeTime string
	if err := global.GetGlobalDBByDBName("freq").Select("close_date").Where("user_id = ? AND is_open = 0", uid).Order("id asc").First(&trade).Error; err == nil {
		if trade.CloseDate != nil {
			tradeTime = trade.CloseDate.Format("2006-01-02")
		}
	} else {
		if err.Error() != "record not found" {
			global.GVA_LOG.Warn("查询首单完成时间失败", zap.Error(err), zap.Uint("uid", uid))
		}
	}

	// 统一返回结构
	data := gin.H{
		"user": gin.H{
			"registerTime": registerTime, // 注册时间（YYYY-MM-DD）
			"riskTime":     riskTime,     // 风评完成时间（YYYY-MM-DD），可能为空
			"apiTime":      apiTime,      // API 审核通过时间（YYYY-MM-DD），可能为空
			"tradeTime":    tradeTime,    // 第一笔交易完成时间（YYYY-MM-DD），可能为空
		},
	}
	response.OkWithDetailed(data, "获取成功", c)
}
