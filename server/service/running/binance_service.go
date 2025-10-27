package running

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type BinanceService struct{}

// get AccountInfo
func (binanceService *BinanceService) GetBinanceAccountInfo(apiKey, secretKey string) (*binance.Account, error) {
	client := binance.NewClient(apiKey, secretKey)
	account, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		global.GVA_LOG.Error("Binance API 调用失败", zap.Error(err))
		return nil, err
	}
	return account, nil
}

// get apiRestrictions  on 2025.08.21 10:11
func (binanceService *BinanceService) GetApiResInfo(apiKey, secretKey string) (*binance.APIKeyPermission, error) {
	client := binance.NewClient(apiKey, secretKey)
	apiRes, err := client.NewGetAPIKeyPermission().Do(context.Background())
	if err != nil {
		global.GVA_LOG.Error("Binance API Res 调用失败", zap.Error(err))
		return nil, err
	}
	return apiRes, nil
}
