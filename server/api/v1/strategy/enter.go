package strategy

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SysFreqStrategyApi
	KeyvaluestoreApi
}

var (
	sysFreqStrategyService = service.ServiceGroupApp.StrategyServiceGroup.SysFreqStrategyService
	keyvaluestoreService   = service.ServiceGroupApp.StrategyServiceGroup.KeyvaluestoreService
)
