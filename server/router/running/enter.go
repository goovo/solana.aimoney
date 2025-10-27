package running

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SysUserRiskRouter
	SysUserApiRouter
	SysUserAssetsRouter
	TradesRouter
	BinanceRouter
	SysUserAibotRouter
}

var (
	sysUserRiskApi   = api.ApiGroupApp.RunningApiGroup.SysUserRiskApi
	sysUserApiApi    = api.ApiGroupApp.RunningApiGroup.SysUserApiApi
	sysUserAssetsApi = api.ApiGroupApp.RunningApiGroup.SysUserAssetsApi
	tradesApi        = api.ApiGroupApp.RunningApiGroup.TradesApi
	binanceApi       = api.ApiGroupApp.RunningApiGroup.BinanceApi
	sysUserAibotApi  = api.ApiGroupApp.RunningApiGroup.SysUserAibotApi
)
