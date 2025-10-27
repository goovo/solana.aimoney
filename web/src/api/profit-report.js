import service from '@/utils/request'

// @Summary 获取收益报表
// @Produce  application/json
// @Param data body {type:"string",startTime:"string",endTime:"string"} true "查询参数"
// @Success  200  {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /system/profit-report [post]
export const getProfitReport = (data) => {
  return service({
    url: '/system/profit-report',
    method: 'post',
    data
  })
}