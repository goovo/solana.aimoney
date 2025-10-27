package response

import (
	"time"
)

// InviteInfoResponse 推荐信息响应
type InviteInfoResponse struct {
	InviteCode   string `json:"inviteCode"`   // 我的推荐码
	InviteLink   string `json:"inviteLink"`   // 推荐链接
	TotalReferral int64 `json:"totalReferral"` // 总推荐人数
}

// ReferralUserInfo 推荐用户信息
type ReferralUserInfo struct {
	ID           uint      `json:"id"`           // 用户ID
	NickName     string    `json:"nickName"`     // 用户昵称
	Username     string    `json:"username"`     // 用户名
	CreatedAt    time.Time `json:"createdAt"`    // 注册时间
	HasInvested  bool      `json:"hasInvested"`  // 是否已投资
	InvestTime   *time.Time `json:"investTime"`  // 投资时间
	Phone        string    `json:"phone"`        // 手机号（脱敏）
	Email        string    `json:"email"`        // 邮箱（脱敏）
}

// MyReferralsResponse 我的推荐列表响应
type MyReferralsResponse struct {
	List     []ReferralUserInfo `json:"list"`
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"pageSize"`
}