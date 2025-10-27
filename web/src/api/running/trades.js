import service from '@/utils/request'
// @Tags Trades
// @Summary 创建交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Trades true "创建交易报表(模拟盘)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /trades/createTrades [post]
export const createTrades = (data) => {
  return service({
    url: '/trades/createTrades',
    method: 'post',
    data
  })
}

// 为当前用户创建模拟机器人进程
export const createDryRun = (data) => {
  return service({
    url: '/trades/createDryRun',
    method: 'post',
    data
  })
}

// 获取当前用户模拟机器人进程id
export const getDryRun = () => {
  return service({
    url: '/trades/getDryRun',
    method: 'get',
  })
}

// @Tags Trades
// @Summary 删除交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Trades true "删除交易报表(模拟盘)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trades/deleteTrades [delete]
export const deleteTrades = (params) => {
  return service({
    url: '/trades/deleteTrades',
    method: 'delete',
    params
  })
}

// @Tags Trades
// @Summary 批量删除交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除交易报表(模拟盘)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trades/deleteTrades [delete]
export const deleteTradesByIds = (params) => {
  return service({
    url: '/trades/deleteTradesByIds',
    method: 'delete',
    params
  })
}

// @Tags Trades
// @Summary 更新交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Trades true "更新交易报表(模拟盘)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /trades/updateTrades [put]
export const updateTrades = (data) => {
  return service({
    url: '/trades/updateTrades',
    method: 'put',
    data
  })
}

// @Tags Trades
// @Summary 用id查询交易报表(模拟盘)
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Trades true "用id查询交易报表(模拟盘)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /trades/findTrades [get]
export const findTrades = (params) => {
  return service({
    url: '/trades/findTrades',
    method: 'get',
    params
  })
}

// @Tags Trades
// @Summary 分页获取交易报表(模拟盘)列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取交易报表(模拟盘)列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trades/getTradesList [get]
export const getTradesList = (params) => {
  return service({
    url: '/trades/getTradesList',
    method: 'get',
    params
  })
}

// @Tags Trades
// @Summary 不需要鉴权的交易报表(模拟盘)接口
// @Accept application/json
// @Produce application/json
// @Param data query runningReq.TradesSearch true "分页获取交易报表(模拟盘)列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /trades/getTradesPublic [get]
export const getTradesPublic = () => {
  return service({
    url: '/trades/getTradesPublic',
    method: 'get',
  })
}
