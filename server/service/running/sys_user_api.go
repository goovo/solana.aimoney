package running

import (
	"context"
	"errors"
	"strconv"
	"strings"

	// 新增：用于 OKX/Bybit 签名和请求的标准库包
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"
	"math"

	// 新增：本地文件系统与操作系统判断
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	futures "github.com/adshao/go-binance/v2/futures"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/running"
	runningReq "github.com/flipped-aurora/gin-vue-admin/server/model/running/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysUserApiService struct{}

// CreateSysUserApi 创建用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) CreateSysUserApi(ctx context.Context, sysUserApi *running.SysUserApi) (err error) {
	err = global.GVA_DB.Create(sysUserApi).Error
	return err
}

// DeleteSysUserApi 删除用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) DeleteSysUserApi(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&running.SysUserApi{}, "id = ?", id).Error
	return err
}

// DeleteSysUserApiByIds 批量删除用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) DeleteSysUserApiByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]running.SysUserApi{}, "id in ?", ids).Error
	return err
}

// UpdateSysUserApi 更新用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) UpdateSysUserApi(ctx context.Context, sysUserApi running.SysUserApi) (err error) {
	err = global.GVA_DB.Model(&running.SysUserApi{}).Where("id = ?", sysUserApi.Id).Updates(&sysUserApi).Error
	return err
}

// GetSysUserApi 根据id获取用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) GetSysUserApi(ctx context.Context, id string) (sysUserApi running.SysUserApi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysUserApi).Error
	return
}

// GetSysUserApiInfoList 分页获取用户APIs记录
// Author [yourname](https://github.com/yourname)
func (sysUserApiService *SysUserApiService) GetSysUserApiInfoList(ctx context.Context, info runningReq.SysUserApiSearch) (list []running.SysUserApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&running.SysUserApi{})

	var sysUserApis []running.SysUserApi
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Order("id desc").Limit(limit).Offset(offset)
	}

	err = db.Find(&sysUserApis).Error
	return sysUserApis, total, err
}

func (sysUserApiService *SysUserApiService) GetSysUserApiInfoListWithUid(ctx context.Context, info runningReq.SysUserApiSearch, uid uint) (list []running.SysUserApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db，并添加用户ID过滤条件，确保只查询当前用户的数据
	db := global.GVA_DB.Model(&running.SysUserApi{}).Where("userId = ?", uid)
	var sysUserApis []running.SysUserApi
	// 如果有条件搜索 下方会自动创建搜索语句

	// 统计总数时同样应用 userId 条件
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Order("id desc").Limit(limit).Offset(offset)
	}

	// 查询列表数据，已应用 userId 条件
	err = db.Find(&sysUserApis).Error
	return sysUserApis, total, err
}

func (sysUserApiService *SysUserApiService) GetSysUserApiPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// BatchCheckAndSyncApis 批量检查API（每5分钟执行）并同步状态与USDT合约可用资金
// 逻辑：
// 1) 查询 sys_user_api 中 status != 1 的记录；
// 2) 根据 exchange 分发到对应交易所校验（当前支持 binance/okx/bybit：检查是否具备 USDT 永续合约交易权限，并读取 USDT 可用余额——OKX/Bybit 侧采用“能成功查询余额即视为可用”的近似判断，后续可细化为更严格的权限校验）；
// 3) 将校验结果写回 sys_user_api.status（1=正常，2=无权限，3=错误），资金写入/更新 sys_user_assets.assets（按 api_id 关联）。
// 4) 当判定状态为正常(1)时：按操作系统选择工作目录、复制模板配置并写入该 api 的 key/secret/exchange
func (sysUserApiService *SysUserApiService) BatchCheckAndSyncApis(ctx context.Context) {
	global.GVA_LOG.Info("BatchCheckAndSyncApis 启动")

	var apis []running.SysUserApi
	// 仅处理状态不为1的记录；若需要连同NULL一起处理，可自行扩展
	if err := global.GVA_DB.Where("status <> ?", 1).Find(&apis).Error; err != nil {
		global.GVA_LOG.Error("查询 sys_user_api 失败", zap.Error(err))
		return
	}

	for _, api := range apis {
		exch := strings.ToLower(derefString(api.Exchange))
		key := derefString(api.Key)
		secret := derefString(api.Secret)
		passphrase := derefString(api.Passwd) // OKX 需要 passphrase

		if api.Id == nil {
			global.GVA_LOG.Warn("sys_user_api 缺少主键Id，跳过")
			continue
		}

		var (
			newStatus = 3 // 默认错误
			usdtAvail = 0.0
		)

		switch exch {
		case "binance", "币安":
			st, bal, err := checkBinanceFutures(ctx, key, secret)
			if err != nil {
				// 将明显的权限类错误判定为2，无权限；其它网络/未知错误为3
				msg := err.Error()
				if strings.Contains(msg, "permission") || strings.Contains(msg, "Invalid API") || strings.Contains(msg, "API-key") || strings.Contains(msg, "403") || strings.Contains(msg, "401") {
					newStatus = 2
				} else {
					newStatus = 3
				}
				global.GVA_LOG.Error("Binance 期货校验失败", zap.Error(err), zap.Int("apiId", derefInt(api.Id)))
			} else {
				newStatus = st
				usdtAvail = bal
			}
		case "okx", "okex":
			// OKX 使用 V5 接口，需要 passphrase；此处调用账户余额接口校验有效性并读取 USDT 可用余额
			st, bal, err := checkOkxFutures(ctx, key, secret, passphrase)
			if err != nil {
				msg := err.Error()
				if strings.Contains(strings.ToLower(msg), "invalid") || strings.Contains(strings.ToLower(msg), "permission") || strings.Contains(msg, "401") || strings.Contains(msg, "403") {
					newStatus = 2
				} else {
					newStatus = 3
				}
				global.GVA_LOG.Error("OKX 期货校验失败", zap.Error(err), zap.Int("apiId", derefInt(api.Id)))
			} else {
				newStatus = st
				usdtAvail = bal
			}
		case "bybit":
			// Bybit 使用 V5 接口，采用钱包余额查询判断可用性并读取 USDT 可用余额
			st, bal, err := checkBybitFutures(ctx, key, secret)
			if err != nil {
				msg := err.Error()
				if strings.Contains(strings.ToLower(msg), "invalid") || strings.Contains(strings.ToLower(msg), "permission") || strings.Contains(msg, "401") || strings.Contains(msg, "403") {
					newStatus = 2
				} else {
					newStatus = 3
				}
				global.GVA_LOG.Error("Bybit 期货校验失败", zap.Error(err), zap.Int("apiId", derefInt(api.Id)))
			} else {
				newStatus = st
				usdtAvail = bal
			}
		default:
			// 其他交易所暂未实现
			newStatus = 3
			global.GVA_LOG.Warn("未实现的交易所，标记为错误", zap.String("exchange", exch), zap.Int("apiId", derefInt(api.Id)))
		}

		// 预先定义 config 目标路径变量，便于在状态回写时一并更新 config_file 字段
		var dstPath string
		// configFileRel 存储相对路径（去除工作目录前缀），用于写入数据库
		var configFileRel string

		// 如果状态为正常，则执行文件复制与配置写入
		if newStatus == 1 {
			// 1) 根据操作系统选择工作目录
			workDir := determineWorkDir()
			// 2) 复制 user_data/config.json 到 user_data/users/config_{uid}_{exchange}.json
			uidVal := derefInt(api.UserId)
			exNameRaw := derefString(api.Exchange)
			// 文件名中的 exchange 保持与数据库一致，不进行大小写转换
			// exNameForFile := strings.ToLower(exNameRaw)
			srcPath := ""
			if exNameRaw == "OKX" {
				srcPath = filepath.Join(workDir, "user_data", "config_okx.json")

			} else {
				srcPath = filepath.Join(workDir, "user_data", "config_binance.json")
			}

			dstDir := filepath.Join(workDir, "user_data", "users")
			dstPath = filepath.Join(dstDir, fmt.Sprintf("config_%d_%s.json", uidVal, exNameRaw))
			// 计算相对路径（去除工作目录前缀），用于数据库字段 config_file
			if rel, err := filepath.Rel(workDir, dstPath); err == nil {
				configFileRel = rel
			} else {
				// 兜底：相对路径计算失败则回退为原始路径（一般不会发生）
				configFileRel = dstPath
			}

			if err := copyFile(srcPath, dstPath); err != nil {
				// 复制失败仅记录错误，不影响数据库状态写回
				global.GVA_LOG.Error("复制模板配置失败", zap.Error(err), zap.String("src", srcPath), zap.String("dst", dstPath))
			} else {
				// 3) 更新目标配置中的 exchange.name/key/secret
				if err := updateFreqtradeConfigExchange(dstPath, exNameRaw, key, secret, passphrase); err != nil {
					global.GVA_LOG.Error("更新配置 exchange 字段失败", zap.Error(err), zap.String("dst", dstPath))
				} else {
					global.GVA_LOG.Info("已生成专属配置", zap.String("path", dstPath))
				}
			}
		}

		// 回写 sys_user_api.status 和 config_file（存相对路径；当 newStatus != 1 时，configFileRel 为空字符串）
		if err := global.GVA_DB.Model(&running.SysUserApi{}).Where("id = ?", api.Id).Updates(map[string]interface{}{
			"status":      newStatus,
			"config_file": configFileRel,
		}).Error; err != nil {
			global.GVA_LOG.Error("更新 sys_user_api.status/config_file 失败", zap.Error(err), zap.Int("apiId", derefInt(api.Id)))
		}

		// 同步/落库 sys_user_assets（按 api_id 关联）
		apiId := derefInt(api.Id)
		var asset running.SysUserAssets
		if err := global.GVA_DB.Where("api_id = ?", apiId).First(&asset).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 不存在则创建
				uid := derefInt(api.UserId)
				ex := derefString(api.Exchange)
				bal := usdtAvail
				newAsset := running.SysUserAssets{
					UserId:   &uid,
					Exchange: &ex,
					ApiId:    &apiId,
					Assets:   &bal,
				}
				if err := global.GVA_DB.Create(&newAsset).Error; err != nil {
					global.GVA_LOG.Error("创建 sys_user_assets 失败", zap.Error(err), zap.Int("apiId", apiId))
				}
			} else {
				global.GVA_LOG.Error("查询 sys_user_assets 失败", zap.Error(err), zap.Int("apiId", apiId))
			}
		} else {
			// 已存在则仅更新 assets 字段
			if err := global.GVA_DB.Model(&running.SysUserAssets{}).Where("api_id = ?", apiId).Update("assets", usdtAvail).Error; err != nil {
				global.GVA_LOG.Error("更新 sys_user_assets.assets 失败", zap.Error(err), zap.Int("apiId", apiId))
			}
		}
	}

	global.GVA_LOG.Info("BatchCheckAndSyncApis 完成")
}

// checkBinanceFutures 校验是否具备 USDT 永续合约交易权限，并返回 USDT 合约账户可用余额
// 返回：status(1=正常可交易; 2=无权限; 3=错误)、usdtAvailable(USDT可用余额)
func checkBinanceFutures(ctx context.Context, apiKey, secret string) (int, float64, error) {
	if apiKey == "" || secret == "" {
		return 3, 0, errors.New("apiKey/secret 为空")
	}

	client := futures.NewClient(apiKey, secret)
	// 通过查询合约账户信息来判断是否有期货权限
	acct, err := client.NewGetAccountService().Do(ctx)
	if err != nil {
		return 0, 0, err
	}

	// 能成功获取账户信息，视为具备期货权限
	// 提取 USDT 的可用余额（AvailableBalance）
	usdt := 0.0
	for _, a := range acct.Assets {
		if strings.EqualFold(a.Asset, "USDT") {
			// AvailableBalance 是字符串，需要转 float64
			if a.AvailableBalance != "" {
				if v, e := strconv.ParseFloat(a.AvailableBalance, 64); e == nil {
					usdt = v
				}
			}
			break
		}
	}
	
	// 新增：若 USDT 可用资金超过 20 美金，尝试进行一次 10 美金的期货下单，并在 5 秒后立即平仓，用于验证 API 是否具备交易权限
	// 约定返回：
	// - 若下单阶段明确无交易权限（如接口返回权限错误），返回 2, usdt, nil
	// - 若能成功下单且成功平仓，返回 1, usdt, nil
	// - 其他异常场景，保持原有 3 表示错误
	if usdt > 50 {
		symbol := "ETHUSDT" // Binance 合约交易对，ETH/USDT 永续合约
		// 1) 将杠杆设置为 1 倍
		if _, err := client.NewChangeLeverageService().Symbol(symbol).Leverage(1).Do(ctx); err != nil {
			// 大概率为权限问题或接口不可用，按无交易权限处理
			global.GVA_LOG.Warn("Binance 设置杠杆失败，可能无交易权限", zap.Error(err))
			return 2, usdt, nil
		}
		
		// 2) 获取参考价格（使用最近 1 根 1m K线的收盘价），用于计算 10 美金对应的下单数量
		klines, err := client.NewKlinesService().Symbol(symbol).Interval("1m").Limit(1).Do(ctx)
		if err != nil || len(klines) == 0 || klines[0] == nil {
			global.GVA_LOG.Error("Binance 获取K线失败", zap.Error(err))
			return 3, usdt, err
		}
		price, err := strconv.ParseFloat(klines[0].Close, 64)
		if err != nil || price <= 0 {
			global.GVA_LOG.Error("Binance K线收盘价解析失败", zap.Error(err), zap.String("close", klines[0].Close))
			return 3, usdt, err
		}
		
		// 3) 计算数量（按 10 USDT 计算），并做简单的小数位控制以适配常见步进
		qtyF := 50.0 / price
		// 简单处理成 3 位小数，避免过长精度（更严谨可获取交易规则过滤步进）
		qty := fmt.Sprintf("%.3f", qtyF)
		
		// 4) 市价做多开仓（非双向持仓模式下，直接 BUY 即可）
		_, err = client.NewCreateOrderService().
			Symbol(symbol).
			Side(futures.SideTypeBuy).
			Type(futures.OrderTypeMarket).
			Quantity(qty).
			Do(ctx)
		if err != nil {
			// 下单失败，若为权限类错误，按无交易权限返回；此处统一按 2 处理并记录详细日志
			global.GVA_LOG.Warn("Binance 市价做多下单失败，可能无交易权限或资金/步进不足", zap.Error(err), zap.String("qty", qty))
			return 2, usdt, nil
		}
		global.GVA_LOG.Info("Binance 市价做多下单成功(验证权限)", zap.String("symbol", symbol), zap.String("qty", qty))
		
		// 5) 持仓 5 秒后，市价平仓（ReduceOnly=true），尝试用同样数量卖出以平掉刚开仓的多单
		time.Sleep(5 * time.Second)
		_, err = client.NewCreateOrderService().
			Symbol(symbol).
			Side(futures.SideTypeSell).
			Type(futures.OrderTypeMarket).
			Quantity(qty).
			ReduceOnly(true).
			Do(ctx)
		if err != nil {
			global.GVA_LOG.Error("Binance 市价平仓失败", zap.Error(err), zap.String("qty", qty))
			return 3, usdt, err
		}
		global.GVA_LOG.Info("Binance 市价平仓成功(权限验证通过)", zap.String("symbol", symbol), zap.String("qty", qty))
		return 1, usdt, nil
	}
	return 1, usdt, nil
}

// checkOkxFutures 校验 OKX USDT 合约可用性并返回 USDT 可用余额
// 说明：OKX 采用 V5 API，签名为 base64(hmac_sha256(timestamp + method + requestPath + body, secret))；
// 这里采用查询账户余额接口（/api/v5/account/balance?ccy=USDT）来判断 API 有效性和读取余额。
func checkOkxFutures(ctx context.Context, apiKey, secret, passphrase string) (int, float64, error) {
	if apiKey == "" || secret == "" || passphrase == "" {
		return 3, 0, errors.New("apiKey/secret/passphrase 为空")
	}

	// 生成 OKX 要求的时间戳，UTC，带毫秒
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	method := "GET"
	requestPath := "/api/v5/account/balance"
	query := "ccy=USDT"
	pathWithQuery := requestPath + "?" + query

	// 计算签名：base64(hmac_sha256(timestamp + method + requestPath(+query) + body, secret))
	prehash := timestamp + method + pathWithQuery
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(prehash))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	url := "https://www.okx.com" + pathWithQuery
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return 3, 0, err
	}

	// 设置 OKX 认证头
	req.Header.Set("OK-ACCESS-KEY", apiKey)
	req.Header.Set("OK-ACCESS-SIGN", sign)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 3, 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Data []struct {
			Details []struct {
				Ccy      string `json:"ccy"`
				AvailBal string `json:"availBal"`
			} `json:"details"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 3, 0, err
	}

	// code != "0" 表示失败；尝试根据错误信息归类为权限问题或其它错误
	if data.Code != "0" {
		msg := strings.ToLower(data.Msg)
		if strings.Contains(msg, "invalid") || strings.Contains(msg, "permission") || resp.StatusCode == 401 || resp.StatusCode == 403 {
			return 2, 0, errors.New("okx 权限或认证失败: " + data.Msg)
		}
		return 3, 0, errors.New("okx 请求失败: " + data.Msg)
	}

	// 提取 USDT 可用余额
	usdt := 0.0
	if len(data.Data) > 0 {
		for _, d := range data.Data[0].Details {
			if strings.EqualFold(d.Ccy, "USDT") && d.AvailBal != "" {
				if v, e := strconv.ParseFloat(d.AvailBal, 64); e == nil {
					usdt = v
				}
				break
			}
		}
	}
	// 如果余额不足 20 美金，则不进行下单校验，直接返回余额
	if usdt <= 50 {
		return 1, usdt, nil
	}

	// 余额充足，进行一次最小风险的交易权限校验：
	// 使用 10 USDT，1 倍杠杆在 ETH-USDT-SWAP 上市价做多，5 秒后市价平仓。
	symbol := "ETH-USDT-SWAP"
	budget := 50.0

	// 1) 获取当前市价（last）
	tickerURL := "https://www.okx.com/api/v5/market/ticker?instId=" + symbol
	tReq, err := http.NewRequestWithContext(ctx, "GET", tickerURL, nil)
	if err != nil {
		return 3, usdt, err
	}
	tResp, err := http.DefaultClient.Do(tReq)
	if err != nil {
		return 3, usdt, err
	}
	defer tResp.Body.Close()
	var tdata struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Data []struct {
			Last string `json:"last"`
		} `json:"data"`
	}
	if err := json.NewDecoder(tResp.Body).Decode(&tdata); err != nil {
		return 3, usdt, err
	}
	if tdata.Code != "0" || len(tdata.Data) == 0 {
		return 3, usdt, errors.New("获取 OKX 行情失败: " + tdata.Msg)
	}
	price, err := strconv.ParseFloat(tdata.Data[0].Last, 64)
	if err != nil || price <= 0 {
		return 3, usdt, errors.New("解析 OKX 行情价格失败")
	}
	global.GVA_LOG.Info("OKX 行情价格", zap.Float64("price", price))

	// 2) 获取合约面值(ctVal)、最小/步进下单量(minSz/lotSz) 以计算下单 sz（张数）
	insURL := "https://www.okx.com/api/v5/public/instruments?instType=SWAP&instId=" + symbol
	insReq, err := http.NewRequestWithContext(ctx, "GET", insURL, nil)
	if err != nil {
		return 3, usdt, err
	}
	insResp, err := http.DefaultClient.Do(insReq)
	if err != nil {
		return 3, usdt, err
	}
	defer insResp.Body.Close()
	var idata struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
		Data []struct {
			InstId string `json:"instId"`
			CtVal  string `json:"ctVal"`
			LotSz  string `json:"lotSz"`
			MinSz  string `json:"minSz"`
		} `json:"data"`
	}
	if err := json.NewDecoder(insResp.Body).Decode(&idata); err != nil {
		return 3, usdt, err
	}
	if idata.Code != "0" || len(idata.Data) == 0 {
		return 3, usdt, errors.New("获取合约参数失败: " + idata.Msg)
	}
	ctVal, _ := strconv.ParseFloat(idata.Data[0].CtVal, 64) // 面值(每张合约对应的标的币数量)
	lotSz, _ := strconv.ParseFloat(idata.Data[0].LotSz, 64) // 下单步进
	minSz, _ := strconv.ParseFloat(idata.Data[0].MinSz, 64) // 最小下单量
	if ctVal <= 0 {
		// 兜底：若未获取到面值，假设面值为 0.01
		ctVal = 0.01
	}
	if lotSz <= 0 {
		lotSz = 1
	}
	if minSz <= 0 {
		minSz = lotSz
	}
	// 预算 10 USDT，在 1x 杠杆下，张数约等于 预算 / (价格 * 面值)
	// 再按步进取整，且不低于最小下单量
	nFloat := budget / (price * ctVal)
	steps := math.Floor(nFloat/lotSz)
	if steps < 1 {
		steps = 1
	}
	sz := steps * lotSz
	var szStr string
	if math.Mod(sz, 1) == 0 {
		szStr = strconv.FormatInt(int64(sz), 10)
	} else {
		szStr = strconv.FormatFloat(sz, 'f', 6, 64)
	}
	global.GVA_LOG.Info("OKX 计划下单参数", zap.Float64("ctVal", ctVal), zap.Float64("lotSz", lotSz), zap.Float64("minSz", minSz), zap.String("sz", szStr))

	// 3) 设置杠杆为 1（全仓）
	{
		postPath := "/api/v5/account/set-leverage"
		body := map[string]string{
			"instId": symbol,
			"lever": "1",
			"mgnMode": "cross",
		}
		bs, _ := json.Marshal(body)
		ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		prehash := ts + "POST" + postPath + string(bs)
		m := hmac.New(sha256.New, []byte(secret))
		m.Write([]byte(prehash))
		sign := base64.StdEncoding.EncodeToString(m.Sum(nil))

		req, err := http.NewRequestWithContext(ctx, "POST", "https://www.okx.com"+postPath, strings.NewReader(string(bs)))
		if err != nil {
			return 3, usdt, err
		}
		req.Header.Set("OK-ACCESS-KEY", apiKey)
		req.Header.Set("OK-ACCESS-SIGN", sign)
		req.Header.Set("OK-ACCESS-TIMESTAMP", ts)
		req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return 3, usdt, err
		}
		defer resp.Body.Close()
		var r struct {
			Code string `json:"code"`
			Msg  string `json:"msg"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&r)
		if r.Code != "0" {
			msg := strings.ToLower(r.Msg)
			if strings.Contains(msg, "invalid") || strings.Contains(msg, "permission") || resp.StatusCode == 401 || resp.StatusCode == 403 || r.Code == "50026" {
				return 2, usdt, errors.New("设置杠杆失败，可能无交易权限: " + r.Msg)
			}
			// 非权限类错误，记录后继续尝试下单
			global.GVA_LOG.Warn("设置杠杆失败", zap.String("code", r.Code), zap.String("msg", r.Msg))
		}
	}

	// 4) 市价开多
	openOrder := func(usePosSide bool) (string, string, int, error) {
		postPath := "/api/v5/trade/order"
		body := map[string]string{
			"instId": symbol,
			"tdMode": "cross",
			"side":   "buy",
			"ordType": "market",
			"sz":     szStr,
		}
		if usePosSide {
			body["posSide"] = "long"
		}
		bs, _ := json.Marshal(body)
		ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		prehash := ts + "POST" + postPath + string(bs)
		m := hmac.New(sha256.New, []byte(secret))
		m.Write([]byte(prehash))
		sign := base64.StdEncoding.EncodeToString(m.Sum(nil))

		req, err := http.NewRequestWithContext(ctx, "POST", "https://www.okx.com"+postPath, strings.NewReader(string(bs)))
		if err != nil {
			return "", "", 0, err
		}
		req.Header.Set("OK-ACCESS-KEY", apiKey)
		req.Header.Set("OK-ACCESS-SIGN", sign)
		req.Header.Set("OK-ACCESS-TIMESTAMP", ts)
		req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", "", resp.StatusCode, err
		}
		defer resp.Body.Close()
		var r struct {
			Code string `json:"code"`
			Msg  string `json:"msg"`
			Data []struct {
				OrdId string `json:"ordId"`
			} `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return "", "", resp.StatusCode, err
		}
		if r.Code != "0" {
			return r.Code, r.Msg, resp.StatusCode, errors.New(r.Msg)
		}
		ordId := ""
		if len(r.Data) > 0 {
			ordId = r.Data[0].OrdId
		}
		return r.Code, ordId, resp.StatusCode, nil
	}

	code, ordId, status, err := openOrder(true)
	if err != nil {
		lc := strings.ToLower(err.Error())
		if code == "51000" && strings.Contains(lc, "posside") {
			// 重试：不传 posSide
			code, ordId, status, err = openOrder(false)
		}
	}
	if err != nil {
		if status == 401 || status == 403 || code == "50026" || strings.Contains(strings.ToLower(err.Error()), "invalid") || strings.Contains(strings.ToLower(err.Error()), "permission") {
			return 2, usdt, errors.New("OKX 开仓失败，无交易权限或认证错误: " + err.Error())
		}
		return 3, usdt, errors.New("OKX 开仓失败: " + err.Error())
	}
	global.GVA_LOG.Info("OKX 开仓成功", zap.String("ordId", ordId), zap.String("sz", szStr))

	// 5) 等待 5 秒后市价平仓（只减仓）
	time.Sleep(5 * time.Second)
	closeOrder := func(usePosSide bool) (string, string, int, error) {
		postPath := "/api/v5/trade/order"
		body := map[string]string{
			"instId":     symbol,
			"tdMode":     "cross",
			"side":       "sell",
			"ordType":    "market",
			"sz":         szStr,
			"reduceOnly": "true",
		}
		if usePosSide {
			body["posSide"] = "long"
		}
		bs, _ := json.Marshal(body)
		ts := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
		prehash := ts + "POST" + postPath + string(bs)
		m := hmac.New(sha256.New, []byte(secret))
		m.Write([]byte(prehash))
		sign := base64.StdEncoding.EncodeToString(m.Sum(nil))

		req, err := http.NewRequestWithContext(ctx, "POST", "https://www.okx.com"+postPath, strings.NewReader(string(bs)))
		if err != nil {
			return "", "", 0, err
		}
		req.Header.Set("OK-ACCESS-KEY", apiKey)
		req.Header.Set("OK-ACCESS-SIGN", sign)
		req.Header.Set("OK-ACCESS-TIMESTAMP", ts)
		req.Header.Set("OK-ACCESS-PASSPHRASE", passphrase)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", "", resp.StatusCode, err
		}
		defer resp.Body.Close()
		var r struct {
			Code string `json:"code"`
			Msg  string `json:"msg"`
			Data []struct {
				OrdId string `json:"ordId"`
			} `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return "", "", resp.StatusCode, err
		}
		if r.Code != "0" {
			return r.Code, r.Msg, resp.StatusCode, errors.New(r.Msg)
		}
		ordId := ""
		if len(r.Data) > 0 {
			ordId = r.Data[0].OrdId
		}
		return r.Code, ordId, resp.StatusCode, nil
	}

	code, ordId, status, err = closeOrder(true)
	if err != nil {
		lc := strings.ToLower(err.Error())
		if code == "51000" && strings.Contains(lc, "posside") {
			code, ordId, status, err = closeOrder(false)
		}
	}
	if err != nil {
		if status == 401 || status == 403 || code == "50026" || strings.Contains(strings.ToLower(err.Error()), "invalid") || strings.Contains(strings.ToLower(err.Error()), "permission") {
			return 2, usdt, errors.New("OKX 平仓失败，无交易权限或认证错误: " + err.Error())
		}
		return 3, usdt, errors.New("OKX 平仓失败: " + err.Error())
	}
	global.GVA_LOG.Info("OKX 平仓成功", zap.String("ordId", ordId), zap.String("sz", szStr))

	return 1, usdt, nil
}

// checkBybitFutures 校验 Bybit USDT 合约可用性并返回 USDT 可用余额
// 说明：Bybit V5 API 签名规则为 hex(hmac_sha256(timestamp + api_key + recv_window + query/body, secret))；
// 这里使用查询钱包余额接口（/v5/account/wallet-balance?accountType=UNIFIED&coin=USDT）。
func checkBybitFutures(ctx context.Context, apiKey, secret string) (int, float64, error) {
	// 基本参数校验
	if apiKey == "" || secret == "" {
		return 3, 0, errors.New("apiKey/secret 为空")
	}

	// 设置请求参数
	method := "GET"
	requestPath := "/v5/account/wallet-balance"
	query := "accountType=UNIFIED&coin=USDT" // 统一账户查询 USDT
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	recvWindow := "5000"

	// v5 签名串：timestamp + apiKey + recvWindow + queryString（GET）
	signPayload := timestamp + apiKey + recvWindow + query
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signPayload))
	sign := hex.EncodeToString(mac.Sum(nil))

	// 发起请求
	url := "https://api.bybit.com" + requestPath + "?" + query
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return 3, 0, err
	}

	// 设置 Bybit 认证头
	req.Header.Set("X-BAPI-API-KEY", apiKey)
	req.Header.Set("X-BAPI-TIMESTAMP", timestamp)
	req.Header.Set("X-BAPI-RECV-WINDOW", recvWindow)
	req.Header.Set("X-BAPI-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 3, 0, err
	}
	defer resp.Body.Close()

	// 解析响应
	var data struct {
		RetCode int    `json:"retCode"`
		RetMsg  string `json:"retMsg"`
		Result  struct {
			List []struct {
				AccountType string `json:"accountType"`
				Coin        []struct {
					Coin                string `json:"coin"`
					AvailableToWithdraw string `json:"availableToWithdraw"`
					WalletBalance       string `json:"walletBalance"`
				} `json:"coin"`
			} `json:"list"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 3, 0, err
	}

	// 错误码判断：权限或认证失败返回 2
	if data.RetCode != 0 {
		msg := strings.ToLower(data.RetMsg)
		if strings.Contains(msg, "invalid") || strings.Contains(msg, "permission") || resp.StatusCode == 401 || resp.StatusCode == 403 {
			return 2, 0, errors.New("bybit 权限或认证失败: " + data.RetMsg)
		}
		return 3, 0, errors.New("bybit 请求失败: " + data.RetMsg)
	}

	// 提取 USDT 可用余额（优先使用 availableToWithdraw）
	usdt := 0.0
	if len(data.Result.List) > 0 {
		for _, c := range data.Result.List[0].Coin {
			if strings.EqualFold(c.Coin, "USDT") {
				if c.AvailableToWithdraw != "" {
					if v, e := strconv.ParseFloat(c.AvailableToWithdraw, 64); e == nil {
						usdt = v
					}
				} else if c.WalletBalance != "" {
					if v, e := strconv.ParseFloat(c.WalletBalance, 64); e == nil {
						usdt = v
					}
				}
				break
			}
		}
	}

	// 返回余额
	return 1, usdt, nil
}

// determineWorkDir 根据操作系统选择工作目录
// macOS 使用 /Users/qinjialiang/Devs/ai/aibot，其他系统（如 Linux/Ubuntu）使用 /trader/freqtrade
func determineWorkDir() string {
	if runtime.GOOS == "darwin" {
		return "/Users/qinjialiang/Devs/ai/aibot"
	}
	return "/trader/freqtrade"
}

// copyFile 将 src 文件复制到 dst，若目标目录不存在则自动创建
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}
	// 默认 0644 权限；若能读取到源文件权限则沿用
	mode := os.FileMode(0o644)
	if fi, err := os.Stat(src); err == nil {
		mode = fi.Mode()
	}
	return os.WriteFile(dst, data, mode)
}

// updateFreqtradeConfigExchange 更新 freqtrade 配置文件中的 exchange.name/key/secret 字段
func updateFreqtradeConfigExchange(path, name, key, secret, passphrase string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	ex, ok := m["exchange"].(map[string]any)
	if !ok || ex == nil {
		ex = map[string]any{}
		m["exchange"] = ex
	}
	if name == "OKX" {
		ex["password"] = passphrase
		ex["name"] = "okx"
	} else {
		ex["name"] = name
	}

	ex["key"] = key
	ex["secret"] = secret

	out, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0o644)
}

// derefString 工具方法：在 *string 可能为 nil 时给出安全的零值
func derefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// derefInt 工具方法：在 *int 可能为 nil 时给出安全的零值
func derefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}
