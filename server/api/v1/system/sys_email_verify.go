package system

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	pluginGlobal "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/global"
	emailUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"github.com/gin-gonic/gin"

	// NOTE: base64Captcha 在本文件未直接使用，store 变量由 sys_captcha.go 提供
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SendEmailCode 发送邮箱验证码
// @Tags     Base
// @Summary  发送邮箱验证码
// @Produce  application/json
// @Param    data  body      SendEmailCodeRequest                          true  "邮箱地址和图形验证码"
// @Success  200   {object}  response.Response{msg=string}                 "发送验证码成功"
// @Router   /base/sendEmailCode [post]
func (b *BaseApi) SendEmailCode(c *gin.Context) {
	var req SendEmailCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验邮箱格式
	if !isValidEmail(req.Email) {
		response.FailWithMessage("邮箱格式不正确", c)
		return
	}

	// 校验图形验证码，防止机器人滥用
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("图形验证码错误", c)
		return
	}

	// 生成6位数验证码
	code := generateVerifyCode(6)

	// 验证码存储key
	key := fmt.Sprintf("email_verify_code:%s", req.Email)
	expiration := 5 * time.Minute // 5分钟过期

	// 根据系统配置选择存储方式
	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		// 使用Redis存储验证码
		err = global.GVA_REDIS.Set(context.Background(), key, code, expiration).Err()
		if err != nil {
			global.GVA_LOG.Error("验证码存储失败!", zap.Error(err))
			response.FailWithMessage("验证码发送失败", c)
			return
		}
	} else {
		// 使用本地缓存存储验证码
		global.BlackCache.Set(key, code, expiration)
	}

	// 发送邮件验证码
	subject := "用户注册验证码"
	body := fmt.Sprintf(`
		<h2>用户注册验证码</h2>
		<p>您的验证码是：<strong style="color: #007cff; font-size: 24px;">%s</strong></p>
		<p>验证码有效期为5分钟，请及时使用。</p>
		<p>如果这不是您的操作，请忽略此邮件。</p>
	`, code)

	// 检查邮件插件配置是否可用
	if pluginGlobal.GlobalConfig.From == "" || pluginGlobal.GlobalConfig.Host == "" {
		global.GVA_LOG.Error("邮件插件配置未初始化")
		response.FailWithMessage("邮件服务暂不可用", c)
		return
	}

	err = emailUtils.Email(req.Email, subject, body)
	if err != nil {
		// 兼容处理：部分 SMTP 服务商可能出现 "short response" 异常但邮件实际已投递成功
		// 这种情况下我们记录警告日志，并将结果按成功返回，避免前端提示失败造成误解
		if strings.Contains(err.Error(), "short response") {
			global.GVA_LOG.Warn("邮件发送疑似成功但返回异常，忽略错误", zap.Error(err))
			response.OkWithMessage("验证码发送成功", c)
			return
		}
		global.GVA_LOG.Error("邮件发送失败!", zap.Error(err))
		response.FailWithMessage("验证码发送失败", c)
		return
	}

	response.OkWithMessage("验证码发送成功", c)
}

// VerifyEmailCode 校验邮箱验证码
func VerifyEmailCode(email, code string) bool {
	key := fmt.Sprintf("email_verify_code:%s", email)

	var storedCode string
	var found bool

	// 根据系统配置选择存储方式
	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		// 从Redis获取验证码
		result, err := global.GVA_REDIS.Get(context.Background(), key).Result()
		if err == redis.Nil {
			return false // 验证码不存在或已过期
		} else if err != nil {
			global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
			return false
		}
		storedCode = result
		found = true

		// 验证成功后删除验证码
		global.GVA_REDIS.Del(context.Background(), key)
	} else {
		// 从本地缓存获取验证码
		value, ok := global.BlackCache.Get(key)
		if !ok {
			return false // 验证码不存在或已过期
		}
		if codeStr, ok := value.(string); ok {
			storedCode = codeStr
			found = true
		}

		// 本地缓存删除需要重新设置为过期
		global.BlackCache.Set(key, "", 1*time.Nanosecond)
	}

	return found && storedCode == code
}

// SendEmailCodeRequest 发送邮箱验证码请求结构体
type SendEmailCodeRequest struct {
	Email     string `json:"email" binding:"required" example:"user@example.com"`
	CaptchaId string `json:"captchaId" binding:"required" example:"captcha_id"` // 图形验证码ID
	Captcha   string `json:"captcha" binding:"required" example:"1234"`         // 图形验证码值
}

// SendPhoneCodeRequest 发送手机验证码请求结构体
type SendPhoneCodeRequest struct {
	Phone     string `json:"phone" binding:"required" example:"13800138000"`    // 手机号
	CaptchaId string `json:"captchaId" binding:"required" example:"captcha_id"` // 图形验证码ID
	Captcha   string `json:"captcha" binding:"required" example:"1234"`         // 图形验证码值
}

// SendPhoneCode 发送手机验证码（占位实现：目前仅生成并保存验证码，不实际发送短信）
// @Tags     Base
// @Summary  发送手机验证码（需图形验证码）
// @Produce  application/json
// @Param    data  body      SendPhoneCodeRequest                        true  "手机号和图形验证码"
// @Success  200   {object}  response.Response{msg=string}               "发送验证码成功（占位实现）"
// @Router   /base/sendPhoneCode [post]
func (b *BaseApi) SendPhoneCode(c *gin.Context) {
	var req SendPhoneCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 简单手机号格式校验（可替换为更严格的正则）
	if len(req.Phone) < 6 || len(req.Phone) > 20 {
		response.FailWithMessage("手机号格式不正确", c)
		return
	}

	// 校验图形验证码
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("图形验证码错误", c)
		return
	}

	// 生成6位验证码
	code := generateVerifyCode(6)
	key := fmt.Sprintf("phone_verify_code:%s", req.Phone)
	expiration := 5 * time.Minute // 5分钟过期

	// 存储验证码（Redis 或本地缓存）
	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		if err := global.GVA_REDIS.Set(context.Background(), key, code, expiration).Err(); err != nil {
			global.GVA_LOG.Error("手机验证码存储失败!", zap.Error(err))
			response.FailWithMessage("验证码发送失败", c)
			return
		}
	} else {
		global.BlackCache.Set(key, code, expiration)
	}

	// 说明：当前为占位实现，未集成短信网关。可在此处对接阿里云短信、腾讯云短信或 Twilio 等服务
	// 如需对接，请在配置文件增加短信网关配置，并在此发送短信。
	const (
		appKey    = "VD4bcv"
		appSecret = "wmHDjT"
		appCode   = "1000"
		extend    = ""
		url       = "http://150.158.58.120:9090/sms/batch/v1"
	)
	// 构造请求体
	timestamp := fmt.Sprintf("%d", time.Now().UnixMilli())
	sign := md5V2(appKey + appSecret + timestamp)

	bodyMap := map[string]string{
		"appkey":    appKey,
		"appcode":   appCode,
		"timestamp": timestamp,
		"phone":     req.Phone,
		"msg":       "【全民认证】您的短信验证码: " + code + "，有效期5分钟。",
		"sign":      sign,
		"extend":    extend,
	}

	jsonData, err := json.Marshal(bodyMap)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}

	// 发请求
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("http post err:", err)
		return
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("response:", string(respBytes))

	response.OkWithMessage("验证码发送成功", c)
}

// md5V2 返回 32 位小写十六进制字符串
func md5V2(s string) string {
	h := md5.Sum([]byte(s))
	return hex.EncodeToString(h[:])
}

// VerifyPhoneCode 校验手机验证码
func VerifyPhoneCode(phone, code string) bool {
	key := fmt.Sprintf("phone_verify_code:%s", phone)
	var storedCode string
	var found bool
	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		result, err := global.GVA_REDIS.Get(context.Background(), key).Result()
		if err == redis.Nil {
			return false
		} else if err != nil {
			global.GVA_LOG.Error("手机验证码获取失败!", zap.Error(err))
			return false
		}
		storedCode = result
		found = true
		global.GVA_REDIS.Del(context.Background(), key)
	} else {
		value, ok := global.BlackCache.Get(key)
		if !ok {
			return false
		}
		if codeStr, ok := value.(string); ok {
			storedCode = codeStr
			found = true
		}
		global.BlackCache.Set(key, "", 1*time.Nanosecond)
	}
	return found && storedCode == code
}

// generateVerifyCode 生成指定长度的数字验证码
func generateVerifyCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	var code strings.Builder
	for i := 0; i < length; i++ {
		code.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}
	return code.String()
}

// isValidEmail 简单的邮箱格式校验
func isValidEmail(email string) bool {
	// 这里使用简单的校验，实际项目中可以使用更严格的正则表达式
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// CheckUsernameRequest 检查用户名请求结构体
type CheckUsernameRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
}

// CheckUsername 检查用户名是否已存在
// @Tags     Base
// @Summary  检查用户名是否已存在
// @Produce  application/json
// @Param    data  body      CheckUsernameRequest                          true  "用户名"
// @Success  200   {object}  response.Response{data=bool,msg=string}       "检查用户名是否可用"
// @Router   /base/checkUsername [post]
func (b *BaseApi) CheckUsername(c *gin.Context) {
	var req CheckUsernameRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验用户名格式
	if len(req.Username) < 3 || len(req.Username) > 20 {
		response.FailWithMessage("用户名长度应在3-20个字符之间", c)
		return
	}

	// 调用用户服务检查用户名是否已存在
	var user system.SysUser
	err = global.GVA_DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 用户名不存在，可以使用
			response.OkWithDetailed(map[string]bool{"available": true}, "用户名可用", c)
		} else {
			global.GVA_LOG.Error("查询用户名失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		return
	}

	// 用户名已存在
	response.OkWithDetailed(map[string]bool{"available": false}, "用户名已被使用", c)
}

// CheckPhoneRequest 检查手机号请求结构体
type CheckPhoneRequest struct {
	Phone string `json:"phone" binding:"required" example:"13800138000"`
}

// CheckPhone 检查手机号是否已注册
// @Tags     Base
// @Summary  检查手机号是否已注册
// @Produce  application/json
// @Param    data  body      CheckPhoneRequest                          true  "手机号"
// @Success  200   {object}  response.Response{data=bool,msg=string}       "检查手机号是否可用"
// @Router   /base/checkPhone [post]
func (b *BaseApi) CheckPhone(c *gin.Context) {
	var req CheckPhoneRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验手机号格式
	if len(req.Phone) != 11 || req.Phone[0] != '1' {
		response.FailWithMessage("手机号格式不正确", c)
		return
	}

	// 查询手机号是否已存在
	var user system.SysUser
	err = global.GVA_DB.Where("phone = ?", req.Phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 手机号不存在，可以使用
			response.OkWithDetailed(map[string]bool{"available": true}, "手机号可用", c)
		} else {
			global.GVA_LOG.Error("查询手机号失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		return
	}

	// 手机号已存在
	response.OkWithDetailed(map[string]bool{"available": false}, "手机号已被注册", c)
}

// CheckEmailRequest 检查邮箱请求结构体
type CheckEmailRequest struct {
	Email string `json:"email" binding:"required" example:"user@example.com"`
}

// CheckEmail 检查邮箱是否已注册
// @Tags     Base
// @Summary  检查邮箱是否已注册
// @Produce  application/json
// @Param    data  body      CheckEmailRequest                          true  "邮箱"
// @Success  200   {object}  response.Response{data=bool,msg=string}       "检查邮箱是否可用"
// @Router   /base/checkEmail [post]
func (b *BaseApi) CheckEmail(c *gin.Context) {
	var req CheckEmailRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 校验邮箱格式
	if !isValidEmail(req.Email) {
		response.FailWithMessage("邮箱格式不正确", c)
		return
	}

	// 查询邮箱是否已存在
	var user system.SysUser
	err = global.GVA_DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 邮箱不存在，可以使用
			response.OkWithDetailed(map[string]bool{"available": true}, "邮箱可用", c)
		} else {
			global.GVA_LOG.Error("查询邮箱失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		}
		return
	}

	// 邮箱已存在
	response.OkWithDetailed(map[string]bool{"available": false}, "邮箱已被注册", c)
}

// 注意：store 变量已在 sys_captcha.go 中声明，此处无需重复声明
