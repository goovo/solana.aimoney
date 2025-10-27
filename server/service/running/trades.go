package running

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
)

type TradesService struct{}

// CreateTrades 创建交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) CreateTrades(ctx context.Context, trades *running.Trades) (err error) {
	err = global.GetGlobalDBByDBName("freq").Create(trades).Error
	return err
}

// DeleteTrades 删除交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) DeleteTrades(ctx context.Context, id string) (err error) {
	err = global.GetGlobalDBByDBName("freq").Delete(&running.Trades{}, "id = ?", id).Error
	return err
}

// DeleteTradesByIds 批量删除交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) DeleteTradesByIds(ctx context.Context, ids []string) (err error) {
	err = global.GetGlobalDBByDBName("freq").Delete(&[]running.Trades{}, "id in ?", ids).Error
	return err
}

// UpdateTrades 更新交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) UpdateTrades(ctx context.Context, trades running.Trades) (err error) {
	err = global.GetGlobalDBByDBName("freq").Model(&running.Trades{}).Where("id = ?", trades.Id).Updates(&trades).Error
	return err
}

// GetTrades 根据id获取交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) GetTrades(ctx context.Context, id string) (trades running.Trades, err error) {
	err = global.GetGlobalDBByDBName("freq").Where("id = ?", id).First(&trades).Error
	return
}

// GetTradesInfoList 分页获取交易报表(模拟盘)记录
// Author [yourname](https://github.com/yourname)
func (tradesService *TradesService) GetTradesInfoList(ctx context.Context, info runningReq.TradesSearch) (list []running.Trades, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("freq").Model(&running.Trades{})
	var tradess []running.Trades
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&tradess).Error
	return tradess, total, err
}

func (tradesService *TradesService) GetTradesInfoListWithUid(ctx context.Context, info runningReq.TradesSearch, uid uint) (list []running.Trades, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("freq").Model(&running.Trades{}).Where("user_id = ?", uid)
	var tradess []running.Trades
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id DESC").Find(&tradess).Error
	return tradess, total, err
}

func (tradesService *TradesService) GetTradesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
