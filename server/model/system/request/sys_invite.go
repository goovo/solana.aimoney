package request

// InviteRequest 获取推荐信息请求
type InviteRequest struct {
}

// MyReferralsRequest 获取我的推荐列表请求
type MyReferralsRequest struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"` // 搜索关键词
}