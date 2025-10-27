package system

import (
	"bufio"
	"bytes"
	"context"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

type FreqStrategyService struct{}

// SyncListStrategies 主入口：执行命令 → 解析 → 入库
func (s *FreqStrategyService) SyncListStrategies(ctx context.Context) {
	logger := global.GVA_LOG.Named("timer.freqStrategy")

	out, err := runFreqtradeListStrategies(ctx)
	if err != nil {
		logger.Error("freqtrade 命令失败", zap.Error(err), zap.ByteString("out", out))
		return
	}

	rows := parseCliOutput(string(out))
	for _, r := range rows {
		// 按 Name 判重
		var cnt int64
		global.GVA_DB.WithContext(ctx).Model(&system.FreqStrategy{}).
			Where("name = ?", r.Name).Count(&cnt)
		if cnt > 0 {
			continue
		}
		if err := global.GVA_DB.WithContext(ctx).Create(&r).Error; err != nil {
			logger.Error("写入策略失败", zap.Error(err))
		}
	}
	logger.Info("同步策略完成", zap.Int("new", len(rows)))
}

// -------------- 私有解析函数 --------------
var (
	fileRe     = regexp.MustCompile(`^(ai|s)_(\d+[smhd])_(more|less)_strategy\.py$`)
	headerLine = "┏━━━━━━━━━━━━━━━━━━━━━━┳" // 用于定位表头
)

func parseCliOutput(raw string) []system.FreqStrategy {
	var list []system.FreqStrategy
	sc := bufio.NewScanner(bytes.NewBufferString(raw))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.HasPrefix(line, "│") {
			continue // 跳过装饰线
		}
		// 去掉首尾的 │
		line = strings.Trim(line, "│")
		cols := strings.Split(line, "│")
		if len(cols) < 6 {
			continue
		}
		name := strings.TrimSpace(cols[0])
		file := strings.TrimSpace(cols[1])
		status := strings.TrimSpace(cols[2])
		hyper := strings.TrimSpace(cols[3]) == "Yes"
		buy, _ := strconv.Atoi(strings.TrimSpace(cols[4]))
		sell, _ := strconv.Atoi(strings.TrimSpace(cols[5]))

		ai, tf, dir := parseFileName(file)
		list = append(list, system.FreqStrategy{
			Name:       name,
			FileName:   file,
			Status:     status,
			Hyperopt:   hyper,
			BuyParams:  uint(buy),
			SellParams: uint(sell),
			AI:         ai,
			TimeFrame:  tf,
			Direction:  dir,
		})
	}
	return list
}

func parseFileName(fname string) (ai bool, tf, dir string) {
	m := fileRe.FindStringSubmatch(fname)
	if len(m) != 4 {
		return false, "", ""
	}
	ai = (m[1] == "ai")
	tf = m[2]
	dir = m[3]
	return
}

// -------------- 执行命令并捕获输出 --------------
func runFreqtradeListStrategies(ctx context.Context) ([]byte, error) {
	// 1. 构造一个带虚拟环境的 shell 指令
	//    注意：\n 用来保证三条语句依次执行
	shCmd := `
cd /trader/freqtrade
source .venv/bin/activate
freqtrade list-strategies
`

	// 2. 用 bash -c 启动，并把 stdout/err 都抓出来
	cmd := exec.CommandContext(ctx, "bash", "-c", shCmd)
	cmd.Dir = "/trader/freqtrade" // 工作目录
	cmd.Env = append(os.Environ(),
		"PATH=/trader/freqtrade/.venv/bin:"+os.Getenv("PATH"), // 确保虚拟环境的 PATH 在最前面
	)
	// 3. 捕获输出
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
