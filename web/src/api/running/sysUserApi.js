import service from '@/utils/request'
// @Tags SysUserApi
// @Summary 创建用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserApi true "创建用户APIs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysUserApi/createSysUserApi [post]
export const createSysUserApi = (data) => {
  return service({
    url: '/sysUserApi/createSysUserApi',
    method: 'post',
    data
  })
}

// @Tags SysUserApi
// @Summary 删除用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserApi true "删除用户APIs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserApi/deleteSysUserApi [delete]
export const deleteSysUserApi = (params) => {
  return service({
    url: '/sysUserApi/deleteSysUserApi',
    method: 'delete',
    params
  })
}

// @Tags SysUserApi
// @Summary 批量删除用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户APIs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserApi/deleteSysUserApi [delete]
export const deleteSysUserApiByIds = (params) => {
  return service({
    url: '/sysUserApi/deleteSysUserApiByIds',
    method: 'delete',
    params
  })
}

// @Tags SysUserApi
// @Summary 更新用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserApi true "更新用户APIs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUserApi/updateSysUserApi [put]
export const updateSysUserApi = (data) => {
  return service({
    url: '/sysUserApi/updateSysUserApi',
    method: 'put',
    data
  })
}

// @Tags SysUserApi
// @Summary 用id查询用户APIs
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysUserApi true "用id查询用户APIs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUserApi/findSysUserApi [get]
export const findSysUserApi = (params) => {
  return service({
    url: '/sysUserApi/findSysUserApi',
    method: 'get',
    params
  })
}

// @Tags SysUserApi
// @Summary 分页获取用户APIs列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户APIs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUserApi/getSysUserApiList [get]
export const getSysUserApiList = (params) => {
  return service({
    url: '/sysUserApi/getSysUserApiList',
    method: 'get',
    params
  })
}

export const getUserApiList = (params) => {
  return service({
    url: '/sysUserApi/getUserApiList',
    method: 'get',
    params
  })
}
// @Tags SysUserApi
// @Summary 不需要鉴权的用户APIs接口
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserApiSearch true "分页获取用户APIs列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserApi/getSysUserApiPublic [get]
export const getSysUserApiPublic = () => {
  return service({
    url: '/sysUserApi/getSysUserApiPublic',
    method: 'get',
  })
}
