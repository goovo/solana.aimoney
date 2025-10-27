package running

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
)

type SysUserAssetsService struct{}

// CreateSysUserAssets 创建用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) CreateSysUserAssets(ctx context.Context, sysUserAssets *running.SysUserAssets) (err error) {
	err = global.GVA_DB.Create(sysUserAssets).Error
	return err
}

// DeleteSysUserAssets 删除用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) DeleteSysUserAssets(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&running.SysUserAssets{}, "id = ?", id).Error
	return err
}

// DeleteSysUserAssetsByIds 批量删除用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) DeleteSysUserAssetsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]running.SysUserAssets{}, "id in ?", ids).Error
	return err
}

// UpdateSysUserAssets 更新用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) UpdateSysUserAssets(ctx context.Context, sysUserAssets running.SysUserAssets) (err error) {
	err = global.GVA_DB.Model(&running.SysUserAssets{}).Where("id = ?", sysUserAssets.Id).Updates(&sysUserAssets).Error
	return err
}

// GetSysUserAssets 根据id获取用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) GetSysUserAssets(ctx context.Context, id string) (sysUserAssets running.SysUserAssets, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysUserAssets).Error
	return
}

// GetSysUserAssetsInfoList 分页获取用户资产记录
// Author [yourname](https://github.com/yourname)
func (sysUserAssetsService *SysUserAssetsService) GetSysUserAssetsInfoList(ctx context.Context, info runningReq.SysUserAssetsSearch) (list []running.SysUserAssets, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&running.SysUserAssets{})
	var sysUserAssetss []running.SysUserAssets
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&sysUserAssetss).Error
	return sysUserAssetss, total, err
}
func (sysUserAssetsService *SysUserAssetsService) GetSysUserAssetsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
