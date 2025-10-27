import service from '@/utils/request'

// 获取加密货币数据
export const getCryptoData = () => {
  return service({
    url: '/user/crypto/getCryptoData',
    method: 'get'
  })
}