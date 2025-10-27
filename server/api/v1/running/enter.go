package running

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysUserRiskApi
	SysUserApiApi
	SysUserAssetsApi
	TradesApi
	BinanceApi
	SysUserAibotApi
}

var (
	sysUserRiskService   = service.ServiceGroupApp.RunningServiceGroup.SysUserRiskService
	sysUserApiService    = service.ServiceGroupApp.RunningServiceGroup.SysUserApiService
	sysUserAssetsService = service.ServiceGroupApp.RunningServiceGroup.SysUserAssetsService
	tradesService        = service.ServiceGroupApp.RunningServiceGroup.TradesService
	binanceService       = service.ServiceGroupApp.RunningServiceGroup.BinanceService
	sysUserAibotService  = service.ServiceGroupApp.RunningServiceGroup.SysUserAibotService
)
