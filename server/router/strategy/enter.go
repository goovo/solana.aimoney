package strategy

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SysFreqStrategyRouter
	KeyvaluestoreRouter
}

var (
	sysFreqStrategyApi = api.ApiGroupApp.StrategyApiGroup.SysFreqStrategyApi
	keyvaluestoreApi   = api.ApiGroupApp.StrategyApiGroup.KeyvaluestoreApi
)
