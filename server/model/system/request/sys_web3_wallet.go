package request

// PhantomLoginRequest Phantom钱包登录请求
type PhantomLoginRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"` // 钱包地址
	Signature     string `json:"signature" binding:"required"`     // 签名
	Message       string `json:"message" binding:"required"`       // 签名的消息
}

// PhantomNonceRequest 获取Nonce请求
type PhantomNonceRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"` // 钱包地址
}

