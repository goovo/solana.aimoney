package response

// PhantomNonceResponse Nonce响应
type PhantomNonceResponse struct {
	Nonce   string `json:"nonce"`   // 随机字符串
	Message string `json:"message"` // 需要签名的完整消息
}

