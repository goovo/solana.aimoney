import service from '@/utils/request'
// @Tags SysUserRisk
// @Summary 创建用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserRisk true "创建用户风险等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysUserRisk/createSysUserRisk [post]
export const createSysUserRisk = (data) => {
  return service({
    url: '/sysUserRisk/createSysUserRisk',
    method: 'post',
    data
  })
}

// @Tags User
// @Summary 设置用户风险等级
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /sysUserRisk/setUserRisk [post]
export const setUserRisk = (data) => {
  return service({
    url: '/sysUserRisk/setUserRisk',
    method: 'post',
    data: data
  })
}


// @Tags SysUserRisk
// @Summary 删除用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserRisk true "删除用户风险等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserRisk/deleteSysUserRisk [delete]
export const deleteSysUserRisk = (params) => {
  return service({
    url: '/sysUserRisk/deleteSysUserRisk',
    method: 'delete',
    params
  })
}

// @Tags SysUserRisk
// @Summary 批量删除用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户风险等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserRisk/deleteSysUserRisk [delete]
export const deleteSysUserRiskByIds = (params) => {
  return service({
    url: '/sysUserRisk/deleteSysUserRiskByIds',
    method: 'delete',
    params
  })
}

// @Tags SysUserRisk
// @Summary 更新用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserRisk true "更新用户风险等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUserRisk/updateSysUserRisk [put]
export const updateSysUserRisk = (data) => {
  return service({
    url: '/sysUserRisk/updateSysUserRisk',
    method: 'put',
    data
  })
}

// @Tags SysUserRisk
// @Summary 用id查询用户风险等级
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysUserRisk true "用id查询用户风险等级"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUserRisk/findSysUserRisk [get]
export const findSysUserRisk = (params) => {
  return service({
    url: '/sysUserRisk/findSysUserRisk',
    method: 'get',
    params
  })
}

// @Tags SysUserRisk
// @Summary 分页获取用户风险等级列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户风险等级列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUserRisk/getSysUserRiskList [get]
export const getSysUserRiskList = (params) => {
  return service({
    url: '/sysUserRisk/getSysUserRiskList',
    method: 'get',
    params
  })
}

// @Tags SysUserRisk
// @Summary 不需要鉴权的用户风险等级接口
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserRiskSearch true "分页获取用户风险等级列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserRisk/getSysUserRiskPublic [get]
export const getSysUserRiskPublic = () => {
  return service({
    url: '/sysUserRisk/getSysUserRiskPublic',
    method: 'get',
  })
}
