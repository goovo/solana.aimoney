package strategy

import (
	"bufio"
	"context"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/strategy"
	strategyReq "github.com/flipped-aurora/gin-vue-admin/server/model/strategy/request"
)

type KeyvaluestoreService struct{}

// CreateKeyvaluestore 创建keyvaluestore表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) CreateKeyvaluestore(ctx context.Context, keyvaluestore *strategy.Keyvaluestore) (err error) {
	err = global.GVA_DB.Create(keyvaluestore).Error
	return err
}

// DeleteKeyvaluestore 删除keyvaluestore表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) DeleteKeyvaluestore(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&strategy.Keyvaluestore{}, "id = ?", id).Error
	return err
}

// DeleteKeyvaluestoreByIds 批量删除keyvaluestore表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) DeleteKeyvaluestoreByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]strategy.Keyvaluestore{}, "id in ?", ids).Error
	return err
}

// UpdateKeyvaluestore 更新keyvaluestore表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) UpdateKeyvaluestore(ctx context.Context, keyvaluestore strategy.Keyvaluestore) (err error) {
	err = global.GVA_DB.Model(&strategy.Keyvaluestore{}).Where("id = ?", keyvaluestore.Id).Updates(&keyvaluestore).Error
	return err
}

// GetKeyvaluestore 根据id获取keyvaluestore表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) GetKeyvaluestore(ctx context.Context, id string) (keyvaluestore strategy.Keyvaluestore, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&keyvaluestore).Error
	return
}

// GetKeyvaluestoreInfoList 分页获取keyvaluestore表记录 del -- with Goovo 20250817 15:50
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) GetKeyvaluestoreInfoListDel(ctx context.Context, info strategyReq.KeyvaluestoreSearch) (list []strategy.Keyvaluestore, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&strategy.Keyvaluestore{})
	var keyvaluestores []strategy.Keyvaluestore
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&keyvaluestores).Error
	return keyvaluestores, total, err
}
func (keyvaluestoreService *KeyvaluestoreService) GetKeyvaluestorePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetKeyvaluestoreInfoList 分页获取 keyvaluestore 表记录
// Author [yourname](https://github.com/yourname)
func (keyvaluestoreService *KeyvaluestoreService) GetKeyvaluestoreInfoList(
	ctx context.Context,
	info strategyReq.KeyvaluestoreSearch,
) (list []strategy.Keyvaluestore, total int64, err error) {

	// ---------- 1. 执行 shell 拿到 ps 输出 ----------
	// 组合命令：cd 到目录 -> source 激活 -> ps
	cmdStr := `cd /trader/freqtrade && source .venv/bin/activate && ps -ef | grep freqtrade | grep trade`
	cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil && len(out) == 0 {
		// 没有查找到进程也算正常，不算 error
		return []strategy.Keyvaluestore{}, 0, nil
	}

	// ---------- 2. 正则提取 ----------
	// 正则：uid pid ppid ... command
	// command 部分：/path/python /path/freqtrade trade -c <config> --strategy <strategy> --dry-run
	re := regexp.MustCompile(`(?m)^\S+\s+(\d+)\s+\d+\s+.*?\s+(.+?freqtrade)\s+trade\s+-c\s+(\S+)\s+--strategy\s+(\S+)\s+(--dry-run\b)`)
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) != 6 {
			continue
		}
		pid, _ := strconv.ParseInt(matches[1], 10, 64)
		procName := strings.TrimSpace(matches[2])
		configFile := strings.TrimSpace(matches[3])
		strategyName := strings.TrimSpace(matches[4])
		dryRun := strings.TrimSpace(matches[5])

		list = append(list, strategy.Keyvaluestore{
			Pid:          pid,
			ProcessName:  procName,
			ConfigFile:   configFile,
			StrategyName: strategyName,
			DryRun:       dryRun,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}

	// ---------- 3. 分页 ----------
	total = int64(len(list))
	if info.PageSize <= 0 {
		return list, total, nil
	}
	start := info.PageSize * (info.Page - 1)
	if start >= int(total) {
		return []strategy.Keyvaluestore{}, total, nil
	}
	end := start + info.PageSize
	if end > int(total) {
		end = int(total)
	}
	return list[start:end], total, nil
}
