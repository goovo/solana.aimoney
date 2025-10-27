import service from '@/utils/request'


// @Tags APIs
// @Summary get User of CEX with API Info
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /binance/getUserInfo [get]
export const getApiResInfo = (userId) => {
  return service({
    url: '/binance/apiRes/'+userId,
    method: 'get'
  })
}

export const getAccount = (userId) => {
  return service({
    url: '/binance/account/'+userId,
    method: 'get'
  })
}


