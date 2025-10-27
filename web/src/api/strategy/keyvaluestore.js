import service from '@/utils/request'
// @Tags Keyvaluestore
// @Summary 创建keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Keyvaluestore true "创建keyvaluestore表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /keyvaluestore/createKeyvaluestore [post]
export const createKeyvaluestore = (data) => {
  return service({
    url: '/keyvaluestore/createKeyvaluestore',
    method: 'post',
    data
  })
}

// @Tags Keyvaluestore
// @Summary 删除keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Keyvaluestore true "删除keyvaluestore表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keyvaluestore/deleteKeyvaluestore [delete]
export const deleteKeyvaluestore = (params) => {
  return service({
    url: '/keyvaluestore/deleteKeyvaluestore',
    method: 'delete',
    params
  })
}

// @Tags Keyvaluestore
// @Summary 批量删除keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除keyvaluestore表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /keyvaluestore/deleteKeyvaluestore [delete]
export const deleteKeyvaluestoreByIds = (params) => {
  return service({
    url: '/keyvaluestore/deleteKeyvaluestoreByIds',
    method: 'delete',
    params
  })
}

// @Tags Keyvaluestore
// @Summary 更新keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Keyvaluestore true "更新keyvaluestore表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /keyvaluestore/updateKeyvaluestore [put]
export const updateKeyvaluestore = (data) => {
  return service({
    url: '/keyvaluestore/updateKeyvaluestore',
    method: 'put',
    data
  })
}

// @Tags Keyvaluestore
// @Summary 用id查询keyvaluestore表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Keyvaluestore true "用id查询keyvaluestore表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /keyvaluestore/findKeyvaluestore [get]
export const findKeyvaluestore = (params) => {
  return service({
    url: '/keyvaluestore/findKeyvaluestore',
    method: 'get',
    params
  })
}

// @Tags Keyvaluestore
// @Summary 分页获取keyvaluestore表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取keyvaluestore表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /keyvaluestore/getKeyvaluestoreList [get]
export const getKeyvaluestoreList = (params) => {
  return service({
    url: '/keyvaluestore/getKeyvaluestoreList',
    method: 'get',
    params
  })
}

// @Tags Keyvaluestore
// @Summary 不需要鉴权的keyvaluestore表接口
// @Accept application/json
// @Produce application/json
// @Param data query strategyReq.KeyvaluestoreSearch true "分页获取keyvaluestore表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /keyvaluestore/getKeyvaluestorePublic [get]
export const getKeyvaluestorePublic = () => {
  return service({
    url: '/keyvaluestore/getKeyvaluestorePublic',
    method: 'get',
  })
}
