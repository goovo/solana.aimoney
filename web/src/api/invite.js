import request from '@/utils/request'

// 获取推荐信息
export const getInviteInfo = () => {
  return request({
    url: '/user/invite/getInviteInfo',
    method: 'get'
  })
}

// 获取我的推荐列表
export const getMyReferrals = (params) => {
  return request({
    url: '/user/invite/getMyReferrals',
    method: 'get',
    params
  })
}

// 检查推荐码有效性
export const checkReferrerCode = (code) => {
  return request({
    url: '/base/checkReferrerCode',
    method: 'get',
    params: { code }
  })
}