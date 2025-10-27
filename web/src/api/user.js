import service from '@/utils/request'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = () => {
  return service({
    url: '/base/captcha',
    method: 'post'
  })
}

// @Summary 用户注册（公开注册）
// @Produce  application/json
// @Param data body {password:"string",email:"string",emailCode:"string",phone:"string",phoneCode:"string",authorityId:number,authorityIds:number[]} true "公开注册，支持手机或邮箱验证码注册"
// @Router /base/register [post]
export const publicRegister = (data) => {
  // 使用公开注册接口，避免与管理员创建用户接口冲突
  return service({
    url: '/base/register',
    method: 'post',
    data
  })
}

// @Summary 发送邮箱验证码（公开）
// @Produce  application/json
// @Param data body {email:"string",captchaId:"string",captcha:"string"} true "邮箱地址和图形验证码"
// @Router /base/sendEmailCode [post]
export const sendEmailCode = (data) => {
  // 公开发送邮箱验证码接口，需要图形验证码验证
  return service({
    url: '/base/sendEmailCode',
    method: 'post',
    data
  })
}

// @Summary 发送手机验证码（公开）
// @Produce  application/json
// @Param data body {phone:"string",captchaId:"string",captcha:"string"} true "手机号和图形验证码"
// @Router /base/sendPhoneCode [post]
export const sendPhoneCode = (data) => {
  // 公开发送手机验证码接口，需要图形验证码验证
  return service({
    url: '/base/sendPhoneCode',
    method: 'post',
    data
  })
}

// @Summary 检查用户名是否可用（公开）
// @Produce  application/json
// @Param data body {username:"string"} true "要检查的用户名"
// @Router /base/checkUsername [post]
export const checkUsername = (data) => {
  // 公开检查用户名可用性接口
  return service({
    url: '/base/checkUsername',
    method: 'post',
    data
  })
}

// @Summary 检查手机号是否可用（公开）
// @Produce  application/json
// @Param data body {phone:"string"} true "要检查的手机号"
// @Router /base/checkPhone [post]
export const checkPhone = (data) => {
  // 公开检查手机号可用性接口
  return service({
    url: '/base/checkPhone',
    method: 'post',
    data
  })
}

// @Summary 检查邮箱是否可用（公开）
// @Produce  application/json
// @Param data body {email:"string"} true "要检查的邮箱"
// @Router /base/checkEmail [post]
export const checkEmail = (data) => {
  // 公开检查邮箱可用性接口
  return service({
    url: '/base/checkEmail',
    method: 'post',
    data
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return service({
    url: '/user/admin_register',
    method: 'post',
    data: data
  })
}

// @Summary 修改密码
// @Produce  application/json
// @Param data body {username:"string",password:"string",newPassword:"string"}
// @Router /user/changePassword [post]
export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getUserList = (data) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
  return service({
    url: '/user/setUserAuthority',
    method: 'post',
    data: data
  })
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setSelfInfo [put]
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置自身界面配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置自身界面配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setSelfSetting [put]
export const setSelfSetting = (data) => {
  return service({
    url: '/user/setSelfSetting',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/user/getUserInfo',
    method: 'get'
  })
}

export const getUser = (userId) => {
  return service({
    url: '/user/getUser/'+userId,
    method: 'get'
  })
}

export const getRisk = () => {
  return service({
    url: '/user/getRisk',
    method: 'get'
  })
}


export const resetPassword = (data) => {
  return service({
    url: '/user/resetPassword',
    method: 'post',
    data: data
  })
}

// @Summary 获取个人统计信息（API数量、交易笔数、胜率、总收益）
// @Router /user/getStat [get]
export const getStat = () => {
  return service({
    url: '/user/getStat',
    method: 'get'
  })
}

// @Summary 获取用户关键步骤时间（注册/风评/API审核/首单）
// @Router /user/getStep [get]
export const getStep = () => {
  return service({
    url: '/user/getStep',
    method: 'get'
  })
}

