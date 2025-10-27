package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	"github.com/flipped-aurora/gin-vue-admin/server/model/strategy"
	"gorm.io/gorm"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(strategy.SysFreqStrategy{}, strategy.Keyvaluestore{}, running.SysUserRisk{}, running.SysUserApi{}, running.SysUserAssets{}, running.Trades{}, // 中文注释：自动迁移业务表结构（会依据 gorm tag 创建/更新索引）
		// 中文注释：为 sys_user_risk.userId 建立唯一索引，并清理历史重复数据，确保后续 ON DUPLICATE KEY UPDATE 能正确触发
		// ensureSysUserRiskUnique 确保 sys_user_risk.userId 唯一，并创建唯一索引
		// 1）清理历史重复数据：对每个重复的 userId 仅保留最新一条（按 id 倒序）
		// 2）创建唯一索引（若不存在），索引名与模型 tag 一致：uniq_sys_user_risk_userId
		// 中文注释：找出存在重复的 userId 列表
		// 中文注释：对每个重复的 userId，仅保留最新的记录（id 最大），其余删除
		// 兼容指针类型的主键 id
		// 中文注释：创建唯一索引（若不存在）。这里直接按唯一索引名创建，确保为唯一索引
		running.SysUserAibot{})
	if err != nil {
		return err
	}
	if err := ensureSysUserRiskUnique(db); err != nil {
		return err
	}
	return nil
}

func ensureSysUserRiskUnique(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var uids []int
		if err := tx.Raw("SELECT userId FROM sys_user_risk WHERE userId IS NOT NULL GROUP BY userId HAVING COUNT(*) > 1").Scan(&uids).Error; err != nil {
			return err
		}
		for _, uid := range uids {
			var keep running.SysUserRisk
			if err := tx.Where("userId = ?", uid).Order("id DESC").First(&keep).Error; err != nil {
				return err
			}
			keepID := 0
			if keep.Id != nil {
				keepID = *keep.Id
			}
			if err := tx.Where("userId = ? AND id <> ?", uid, keepID).Delete(&running.SysUserRisk{}).Error; err != nil {
				return err
			}
		}
		if !tx.Migrator().HasIndex(&running.SysUserRisk{}, "uniq_sys_user_risk_userId") {
			if err := tx.Migrator().CreateIndex(&running.SysUserRisk{}, "uniq_sys_user_risk_userId"); err != nil {
				return err
			}
		}
		return nil
	})
}
