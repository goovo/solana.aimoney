import service from '@/utils/request'
// @Tags SysUserAssets
// @Summary 创建用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAssets true "创建用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysUserAssets/createSysUserAssets [post]
export const createSysUserAssets = (data) => {
  return service({
    url: '/sysUserAssets/createSysUserAssets',
    method: 'post',
    data
  })
}

// @Tags SysUserAssets
// @Summary 删除用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAssets true "删除用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserAssets/deleteSysUserAssets [delete]
export const deleteSysUserAssets = (params) => {
  return service({
    url: '/sysUserAssets/deleteSysUserAssets',
    method: 'delete',
    params
  })
}

// @Tags SysUserAssets
// @Summary 批量删除用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUserAssets/deleteSysUserAssets [delete]
export const deleteSysUserAssetsByIds = (params) => {
  return service({
    url: '/sysUserAssets/deleteSysUserAssetsByIds',
    method: 'delete',
    params
  })
}

// @Tags SysUserAssets
// @Summary 更新用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysUserAssets true "更新用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUserAssets/updateSysUserAssets [put]
export const updateSysUserAssets = (data) => {
  return service({
    url: '/sysUserAssets/updateSysUserAssets',
    method: 'put',
    data
  })
}

// @Tags SysUserAssets
// @Summary 用id查询用户资产
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysUserAssets true "用id查询用户资产"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUserAssets/findSysUserAssets [get]
export const findSysUserAssets = (params) => {
  return service({
    url: '/sysUserAssets/findSysUserAssets',
    method: 'get',
    params
  })
}

// @Tags SysUserAssets
// @Summary 分页获取用户资产列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户资产列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUserAssets/getSysUserAssetsList [get]
export const getSysUserAssetsList = (params) => {
  return service({
    url: '/sysUserAssets/getSysUserAssetsList',
    method: 'get',
    params
  })
}

// @Tags SysUserAssets
// @Summary 不需要鉴权的用户资产接口
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.SysUserAssetsSearch true "分页获取用户资产列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysUserAssets/getSysUserAssetsPublic [get]
export const getSysUserAssetsPublic = () => {
  return service({
    url: '/sysUserAssets/getSysUserAssetsPublic',
    method: 'get',
  })
}
