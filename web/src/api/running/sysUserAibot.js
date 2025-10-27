import service from '@/utils/request'
// @Tags SysUserAibot
// @Summary 创建授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAibot true "创建授权交易"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysUserAibot/createSysUserAibot [post]
export const createSysUserAibot = (data) => {
  return service({
    url: '/sysUserAibot/createSysUserAibot',
    method: 'post',
    data
  })
}

// @Tags SysUserAibot
// @Summary 删除授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAibot true "删除授权交易"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserAibot/deleteSysUserAibot [delete]
export const deleteSysUserAibot = (params) => {
  return service({
    url: '/sysUserAibot/deleteSysUserAibot',
    method: 'delete',
    params
  })
}

// @Tags SysUserAibot
// @Summary 批量删除授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除授权交易"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserAibot/deleteSysUserAibot [delete]
export const deleteSysUserAibotByIds = (params) => {
  return service({
    url: '/sysUserAibot/deleteSysUserAibotByIds',
    method: 'delete',
    params
  })
}

// @Tags SysUserAibot
// @Summary 更新授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAibot true "更新授权交易"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUserAibot/updateSysUserAibot [put]
export const updateSysUserAibot = (data) => {
  return service({
    url: '/sysUserAibot/updateSysUserAibot',
    method: 'put',
    data
  })
}

// @Tags SysUserAibot
// @Summary 用id查询授权交易
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysUserAibot true "用id查询授权交易"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUserAibot/findSysUserAibot [get]
export const findSysUserAibot = (params) => {
  return service({
    url: '/sysUserAibot/findSysUserAibot',
    method: 'get',
    params
  })
}

// @Tags SysUserAibot
// @Summary 分页获取授权交易列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取授权交易列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUserAibot/getSysUserAibotList [get]
export const getSysUserAibotList = (params) => {
  return service({
    url: '/sysUserAibot/getSysUserAibotList',
    method: 'get',
    params
  })
}

// @Tags SysUserAibot
// @Summary 不需要鉴权的授权交易接口
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserAibotSearch true "分页获取授权交易列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserAibot/getSysUserAibotPublic [get]
export const getSysUserAibotPublic = () => {
  return service({
    url: '/sysUserAibot/getSysUserAibotPublic',
    method: 'get',
  })
}
