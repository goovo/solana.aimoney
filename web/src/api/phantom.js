import service from '@/utils/request'

// @Tags Phantom
// @Summary 获取Nonce
// @Produce application/json
// @Param data body {walletAddress: string}
// @Success 200 {nonce: string, message: string}
// @Router /phantom/getNonce [post]
export const getNonce = (data) => {
  return service({
    url: '/phantom/getNonce',
    method: 'post',
    data
  })
}

// @Tags Phantom
// @Summary Phantom钱包登录
// @Produce application/json
// @Param data body {walletAddress: string, signature: string, message: string}
// @Success 200 {user: object, token: string, expiresAt: number}
// @Router /phantom/login [post]
export const phantomLogin = (data) => {
  return service({
    url: '/phantom/login',
    method: 'post',
    data
  })
}

// @Tags Web3Wallet
// @Summary 绑定钱包到当前用户
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body {walletAddress: string, signature: string, message: string}
// @Success 200 {msg: string}
// @Router /web3/bindWallet [post]
export const bindWallet = (data) => {
  return service({
    url: '/web3/bindWallet',
    method: 'post',
    data
  })
}

// @Tags Web3Wallet
// @Summary 解绑钱包
// @Security ApiKeyAuth
// @Produce application/json
// @Param walletAddress query string true "钱包地址"
// @Success 200 {msg: string}
// @Router /web3/unbindWallet [delete]
export const unbindWallet = (walletAddress) => {
  return service({
    url: '/web3/unbindWallet',
    method: 'delete',
    params: { walletAddress }
  })
}

