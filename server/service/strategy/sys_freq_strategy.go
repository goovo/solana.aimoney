
package strategy

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/strategy"
    strategyReq "github.com/flipped-aurora/gin-vue-admin/server/model/strategy/request"
)

type SysFreqStrategyService struct {}
// CreateSysFreqStrategy 创建sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService) CreateSysFreqStrategy(ctx context.Context, sysFreqStrategy *strategy.SysFreqStrategy) (err error) {
	err = global.GVA_DB.Create(sysFreqStrategy).Error
	return err
}

// DeleteSysFreqStrategy 删除sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService)DeleteSysFreqStrategy(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&strategy.SysFreqStrategy{},"id = ?",id).Error
	return err
}

// DeleteSysFreqStrategyByIds 批量删除sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService)DeleteSysFreqStrategyByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]strategy.SysFreqStrategy{},"id in ?",ids).Error
	return err
}

// UpdateSysFreqStrategy 更新sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService)UpdateSysFreqStrategy(ctx context.Context, sysFreqStrategy strategy.SysFreqStrategy) (err error) {
	err = global.GVA_DB.Model(&strategy.SysFreqStrategy{}).Where("id = ?",sysFreqStrategy.Id).Updates(&sysFreqStrategy).Error
	return err
}

// GetSysFreqStrategy 根据id获取sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService)GetSysFreqStrategy(ctx context.Context, id string) (sysFreqStrategy strategy.SysFreqStrategy, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysFreqStrategy).Error
	return
}
// GetSysFreqStrategyInfoList 分页获取sysFreqStrategy表记录
// Author [yourname](https://github.com/yourname)
func (sysFreqStrategyService *SysFreqStrategyService)GetSysFreqStrategyInfoList(ctx context.Context, info strategyReq.SysFreqStrategySearch) (list []strategy.SysFreqStrategy, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&strategy.SysFreqStrategy{})
    var sysFreqStrategys []strategy.SysFreqStrategy
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&sysFreqStrategys).Error
	return  sysFreqStrategys, total, err
}
func (sysFreqStrategyService *SysFreqStrategyService)GetSysFreqStrategyPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
