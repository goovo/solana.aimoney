import service from '@/utils/request'
// @Tags SysFreqStrategy
// @Summary 创建sysFreqStrategy表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysFreqStrategy true "创建sysFreqStrategy表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysFreqStrategy/createSysFreqStrategy [post]
export const createSysFreqStrategy = (data) => {
  return service({
    url: '/sysFreqStrategy/createSysFreqStrategy',
    method: 'post',
    data
  })
}

// @Tags SysFreqStrategy
// @Summary 删除sysFreqStrategy表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysFreqStrategy true "删除sysFreqStrategy表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysFreqStrategy/deleteSysFreqStrategy [delete]
export const deleteSysFreqStrategy = (params) => {
  return service({
    url: '/sysFreqStrategy/deleteSysFreqStrategy',
    method: 'delete',
    params
  })
}

// @Tags SysFreqStrategy
// @Summary 批量删除sysFreqStrategy表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysFreqStrategy表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysFreqStrategy/deleteSysFreqStrategy [delete]
export const deleteSysFreqStrategyByIds = (params) => {
  return service({
    url: '/sysFreqStrategy/deleteSysFreqStrategyByIds',
    method: 'delete',
    params
  })
}

// @Tags SysFreqStrategy
// @Summary 更新sysFreqStrategy表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SysFreqStrategy true "更新sysFreqStrategy表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysFreqStrategy/updateSysFreqStrategy [put]
export const updateSysFreqStrategy = (data) => {
  return service({
    url: '/sysFreqStrategy/updateSysFreqStrategy',
    method: 'put',
    data
  })
}

// @Tags SysFreqStrategy
// @Summary 用id查询sysFreqStrategy表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SysFreqStrategy true "用id查询sysFreqStrategy表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysFreqStrategy/findSysFreqStrategy [get]
export const findSysFreqStrategy = (params) => {
  return service({
    url: '/sysFreqStrategy/findSysFreqStrategy',
    method: 'get',
    params
  })
}

// @Tags SysFreqStrategy
// @Summary 分页获取sysFreqStrategy表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysFreqStrategy表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysFreqStrategy/getSysFreqStrategyList [get]
export const getSysFreqStrategyList = (params) => {
  return service({
    url: '/sysFreqStrategy/getSysFreqStrategyList',
    method: 'get',
    params
  })
}

// @Tags SysFreqStrategy
// @Summary 不需要鉴权的sysFreqStrategy表接口
// @Accept application/json
// @Produce application/json
// @Param data query strategyReq.SysFreqStrategySearch true "分页获取sysFreqStrategy表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysFreqStrategy/getSysFreqStrategyPublic [get]
export const getSysFreqStrategyPublic = () => {
  return service({
    url: '/sysFreqStrategy/getSysFreqStrategyPublic',
    method: 'get',
  })
}
