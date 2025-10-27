package running

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
	"go.uber.org/zap"
)

type SysUserAibotService struct{}

// CreateSysUserAibot 创建授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) CreateSysUserAibot(ctx context.Context, sysUserAibot *running.SysUserAibot) (err error) {
	err = global.GVA_DB.Create(sysUserAibot).Error
	return err
}

// DeleteSysUserAibot 删除授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) DeleteSysUserAibot(ctx context.Context, userId string) (err error) {
	err = global.GVA_DB.Delete(&running.SysUserAibot{}, "userId = ?", userId).Error
	return err
}

// DeleteSysUserAibotByIds 批量删除授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) DeleteSysUserAibotByIds(ctx context.Context, userIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]running.SysUserAibot{}, "userId in ?", userIds).Error
	return err
}

// UpdateSysUserAibot 更新授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) UpdateSysUserAibot(ctx context.Context, sysUserAibot running.SysUserAibot) (err error) {
	err = global.GVA_DB.Model(&running.SysUserAibot{}).Where("userId = ?", sysUserAibot.UserId).Updates(&sysUserAibot).Error
	return err
}

// GetSysUserAibot 根据userId获取授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) GetSysUserAibot(ctx context.Context, userId string) (sysUserAibot running.SysUserAibot, err error) {
	err = global.GVA_DB.Where("userId = ?", userId).First(&sysUserAibot).Error
	return
}

// GetSysUserAibotInfoList 分页获取授权交易记录
// Author [yourname](https://github.com/yourname)
func (sysUserAibotService *SysUserAibotService) GetSysUserAibotInfoList(ctx context.Context, info runningReq.SysUserAibotSearch) (list []running.SysUserAibot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&running.SysUserAibot{})
	var sysUserAibots []running.SysUserAibot
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.UserId != nil {
		db = db.Where("userId = ?", *info.UserId)
	}
	if info.AiBot != "" {
		db = db.Where("aiBot = ?", info.AiBot)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysUserAibots).Error
	return sysUserAibots, total, err
}

func (sysUserAibotService *SysUserAibotService) GetSysUserAibotInfoListWithUid(ctx context.Context, info runningReq.SysUserAibotSearch, uid int64) (list []running.SysUserAibot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&running.SysUserAibot{})
	var sysUserAibots []running.SysUserAibot
	// 如果有条件搜索 下方会自动创建搜索语句

	if uid != 0 {
		db = db.Where("userId = ?", uid)
	}
	if info.AiBot != "" {
		db = db.Where("aiBot = ?", info.AiBot)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysUserAibots).Error
	return sysUserAibots, total, err
}

func (sysUserAibotService *SysUserAibotService) GetSysUserAibotPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AutoRunFreqtradeBots 定时任务：每10分钟扫描可跑策略用户，按risk启动对应策略
// 中文说明：
// 1) 判断操作系统选择工作目录（macOS:/Users/qinjialiang/Devs/ai/aibot；其他:/trader/freqtrade），并激活虚拟环境
// 2) 执行联表查询，获取满足条件的用户列表（userId, config_file, risk）
// 3) 逐个用户检查是否已有freqtrade进程存在(通过 ps -ef | grep userId)，存在则跳过，不存在则按risk启动对应策略
// 4) 启动后等待10秒再处理下一个用户
func (sysUserAibotService *SysUserAibotService) AutoRunFreqtradeBots(ctx context.Context) {
	// 1. 根据操作系统选择工作目录
	workDir := determineWorkDir()

	// 2. 执行SQL（使用GORM构建等价查询，避免直写schema前缀导致兼容问题）
	type eligibleUser struct {
		UserId     uint   `gorm:"column:userId"`      // 用户ID
		ConfigFile string `gorm:"column:config_file"` // 配置文件相对路径
		Risk       string `gorm:"column:risk"`        // 风险等级：low/medium/high
	}

	var users []eligibleUser
	db := global.GVA_DB.WithContext(ctx).
		Table("sys_user_api AS a").
		Select("a.userId, a.config_file, r.risk").
		Joins("JOIN sys_user_risk AS r ON r.userId = a.userId").
		Joins("JOIN sys_user_aibot AS b ON b.userId = a.userId AND b.aiBot = 'allow'").
		Joins("JOIN sys_user_assets AS s ON s.userId = a.userId AND s.assets >= ?", 50).
		Where("a.status = ?", 1)

	if err := db.Scan(&users).Error; err != nil {
		global.GVA_LOG.Error("AutoRunFreqtradeBots 查询可跑策略用户失败", zap.Error(err))
		return
	}

	if len(users) == 0 {
		global.GVA_LOG.Info("AutoRunFreqtradeBots 无可跑策略用户")
		return
	}

	global.GVA_LOG.Info("AutoRunFreqtradeBots 开始处理用户", zap.Int("count", len(users)))

	// 3. 遍历用户，检查进程并按risk启动
	for _, u := range users {
		uidStr := strconv.FormatUint(uint64(u.UserId), 10)

		// 3.1 检查是否已有该用户的freqtrade进程在运行
		checkCmdStr := "ps -ef | grep freqtrade | grep trade | grep " + uidStr + " | grep -v grep"
		out, _ := exec.Command("bash", "-c", checkCmdStr).CombinedOutput() // 无匹配时非0，视为无进程
		if strings.TrimSpace(string(out)) != "" {
			// 已有进程在跑，跳过
			global.GVA_LOG.Debug("用户已有freqtrade进程，跳过", zap.Uint("userId", u.UserId))
			continue
		}

		// 3.2 按风险等级选择策略名
		var strategy string
		switch strings.ToLower(strings.TrimSpace(u.Risk)) {
		case "low":
			strategy = "S15more1Strategy"
		case "medium":
			strategy = "S15more3Strategy"
		case "high":
			strategy = "S15more5Strategy"
		default:
			// 未知risk时采用最保守策略
			strategy = "S15more1Strategy"
		}

		// 3.3 校验配置文件路径
		if strings.TrimSpace(u.ConfigFile) == "" {
			global.GVA_LOG.Warn("config_file 为空，无法启动freqtrade", zap.Uint("userId", u.UserId))
			continue
		}

		// 3.4 组装并启动命令（激活虚拟环境 + 后台运行freqtrade）
		// 说明：这里使用 bash -c 与 & 后台运行，不阻塞任务；标准输出与错误输出重定向到 /dev/null
		cmdStr := fmt.Sprintf("source .venv/bin/activate && freqtrade trade -c '%s' -s %s --user-id %s > /dev/null 2>&1 &", u.ConfigFile, strategy, uidStr)
		cmd := exec.Command("bash", "-c", cmdStr)
		cmd.Dir = workDir // 在工作目录下执行，config_file为相对路径

		if err := cmd.Start(); err != nil {
			global.GVA_LOG.Error("启动freqtrade失败", zap.Uint("userId", u.UserId), zap.String("strategy", strategy), zap.Error(err))
			continue
		}
		global.GVA_LOG.Info("已启动freqtrade", zap.Uint("userId", u.UserId), zap.String("strategy", strategy))

		// 3.5 启动后等待10秒再处理下一个用户
		time.Sleep(10 * time.Second)
	}

	global.GVA_LOG.Info("AutoRunFreqtradeBots 处理完成")
}
