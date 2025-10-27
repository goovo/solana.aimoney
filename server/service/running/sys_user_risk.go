package running

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type SysUserRiskService struct{}

// CreateSysUserRisk 创建用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) CreateSysUserRisk(ctx context.Context, sysUserRisk *running.SysUserRisk) (err error) {
	err = global.GVA_DB.Create(sysUserRisk).Error
	return err
}

// setUserRisk 设置用户风险等级记录
// method: 已经有风控记录，更新风险等级，没有则新增一条
func (sysUserRiskService *SysUserRiskService) SetUserRisk(ctx context.Context, sysUserRisk *running.SysUserRisk, userId int) (err error) {
	// 将基本类型 userId 转为指针，适配模型字段类型为 *int
	u := userId
	sysUserRisk.UserId = &u
	// 关键语句：MySQL8 用 `ON DUPLICATE KEY UPDATE`
	err = global.GVA_DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "userId"}}, // 唯一索引/主键冲突字段
		UpdateAll: true,                              // 冲突时更新所有字段（零值也更新）
	}).Create(sysUserRisk).Error

	if err != nil {
		global.GVA_LOG.Error("SaveOrUpdateByUserID failed", zap.Error(err))
		return err
	}
	return nil
}

// DeleteSysUserRisk 删除用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) DeleteSysUserRisk(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&running.SysUserRisk{}, "id = ?", id).Error
	return err
}

// DeleteSysUserRiskByIds 批量删除用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) DeleteSysUserRiskByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]running.SysUserRisk{}, "id in ?", ids).Error
	return err
}

// UpdateSysUserRisk 更新用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) UpdateSysUserRisk(ctx context.Context, sysUserRisk running.SysUserRisk) (err error) {
	err = global.GVA_DB.Model(&running.SysUserRisk{}).Where("id = ?", sysUserRisk.Id).Updates(&sysUserRisk).Error
	return err
}

// GetSysUserRisk 根据id获取用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) GetSysUserRisk(ctx context.Context, id string) (sysUserRisk running.SysUserRisk, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysUserRisk).Error
	return
}

// GetSysUserRiskInfoList 分页获取用户风险等级记录
// Author [yourname](https://github.com/yourname)
func (sysUserRiskService *SysUserRiskService) GetSysUserRiskInfoList(ctx context.Context, info runningReq.SysUserRiskSearch) (list []running.SysUserRisk, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&running.SysUserRisk{})
	var sysUserRisks []running.SysUserRisk
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&sysUserRisks).Error
	return sysUserRisks, total, err
}
func (sysUserRiskService *SysUserRiskService) GetSysUserRiskPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetSysUserRiskByUserId 通过 userId 获取最新的用户风险等级记录
// 中文注释：根据用户ID查询 sys_user_risk 表，按id倒序取最新一条记录
func (sysUserRiskService *SysUserRiskService) GetSysUserRiskByUserId(ctx context.Context, userId int) (record running.SysUserRisk, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("userId = ?", userId).Order("id desc").First(&record).Error
	return
}
