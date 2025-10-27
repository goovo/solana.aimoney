package running

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type BinanceApi struct{}

// GetAccountBalance
// @Tags Binance
// @Summary 获取Binance账户余额和权限
// @Produce  application/json
// @Success 200 {object} response.Response{data=binance.Account} "账户信息"
// @Router /binance/account [get]
func (b *BinanceApi) GetAccountBalance(c *gin.Context) {
	apiKey := "bL1XWq8F62JSOnBCWOZ1Pe86yMBsga4RujSQtbMzgrF9txmMlQNGdEtnqklcOgng"
	secretKey := "RAPCUOydyE5E3xuiRxIaBaCbVTYYvnguckh28QQDNhZceI4n88CSby5i1rpQ1TRW"

	account, err := binanceService.GetBinanceAccountInfo(apiKey, secretKey)
	if err != nil {
		response.FailWithMessage("获取账户信息失败", c)
		return
	}

	response.OkWithData(account, c)
}

func (b *BinanceApi) GetApiRes(c *gin.Context) {
	apiKey := "fecNTPlkb7fO787DRGbLrWNnF6DWOU9GGIgkbhcvGHNn3Vb91e1qfCg6HkiBA0Sg"
	secretKey := "LAMvskAf5hTLcUK2nxXRwu26KGVPWJmyvASevmfeOsa10lMXB0bsE3tWx2NLZPK9"

	apiRes, err := binanceService.GetApiResInfo(apiKey, secretKey)
	if err != nil {
		response.FailWithMessage("获取账户API Res信息失败", c)
		return
	}

	response.OkWithData(apiRes, c)
}
