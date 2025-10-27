package system

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CryptoApi struct{}

var cryptoCron *cron.Cron

// CryptoData 加密货币数据结构
type CryptoData struct {
	ID                uint      `json:"id"`
	Rank              int       `json:"rank"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Icon              string    `json:"icon"`
	Price             float64   `json:"price"`
	Change24h         float64   `json:"change24h"`
	MarketCap         float64   `json:"marketCap"`
	Volume24h         float64   `json:"volume24h"`
	CirculatingSupply float64   `json:"circulatingSupply"`
	TrendData        []float64 `json:"trendData" gorm:"-"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// CryptoStatistics 统计数据结构
type CryptoStatistics struct {
	TotalMarketCap    float64 `json:"totalMarketCap"`
	TotalMarketCapChange float64 `json:"totalMarketCapChange"`
	TotalVolume       float64 `json:"totalVolume"`
	TotalVolumeChange float64 `json:"totalVolumeChange"`
	BTCDominance      float64 `json:"btcDominance"`
	BTCDominanceChange float64 `json:"btcDominanceChange"`
}

// GetCryptoData
// @Tags     Crypto
// @Summary  获取加密货币数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success  200 {object} response.Response{data=map[string]interface{},msg=string} "获取加密货币数据"
// @Router   /user/crypto/getCryptoData [get]
func (c *CryptoApi) GetCryptoData(ctx *gin.Context) {
	// 获取最新的20条数据
	var cryptoList []system.SysCrypto
	err := global.GVA_DB.Order("`rank` asc").Limit(20).Find(&cryptoList).Error
	if err != nil {
		global.GVA_LOG.Error("获取加密货币数据失败!", zap.Error(err))
		response.FailWithMessage("获取数据失败", ctx)
		return
	}

	// 转换为前端需要的格式
	var resultList []CryptoData
	for _, crypto := range cryptoList {
		resultList = append(resultList, CryptoData{
			ID:                crypto.ID,
			Rank:              crypto.Rank,
			Name:              crypto.Name,
			Symbol:            crypto.Symbol,
			Icon:              crypto.Icon,
			Price:             crypto.Price,
			Change24h:         crypto.Change24h,
			MarketCap:         crypto.MarketCap,
			Volume24h:         crypto.Volume24h,
			CirculatingSupply: crypto.CirculatingSupply,
			TrendData:         generateTrendData(),
			CreatedAt:         crypto.CreatedAt,
			UpdatedAt:         crypto.UpdatedAt,
		})
	}

	// 获取统计数据
	statistics, err := getCryptoStatistics()
	if err != nil {
		global.GVA_LOG.Error("获取统计数据失败!", zap.Error(err))
		response.FailWithMessage("获取统计数据失败", ctx)
		return
	}

	response.OkWithData(gin.H{
		"cryptoList": resultList,
		"statistics": statistics,
	}, ctx)
}

// StartCryptoScheduler 启动加密货币数据定时任务
func (c *CryptoApi) StartCryptoScheduler() {
	if cryptoCron != nil {
		cryptoCron.Stop()
	}

	cryptoCron = cron.New()
	
	// 每5分钟执行一次
	_, err := cryptoCron.AddFunc("*/5 * * * *", func() {
		global.GVA_LOG.Info("开始抓取aicoin数据")
		if err := fetchAndUpdateAICoinData(); err != nil {
			global.GVA_LOG.Error("抓取aicoin数据失败", zap.Error(err))
		}
	})

	if err != nil {
		global.GVA_LOG.Error("启动定时任务失败", zap.Error(err))
		return
	}

	cryptoCron.Start()
	global.GVA_LOG.Info("加密货币数据定时任务已启动")

	// 立即执行一次
	go func() {
		time.Sleep(5 * time.Second)
		if err := fetchAndUpdateAICoinData(); err != nil {
			global.GVA_LOG.Error("首次抓取aicoin数据失败", zap.Error(err))
		}
	}()
}

// StopCryptoScheduler 停止定时任务
func (c *CryptoApi) StopCryptoScheduler() {
	if cryptoCron != nil {
		cryptoCron.Stop()
		global.GVA_LOG.Info("加密货币数据定时任务已停止")
	}
}

// fetchAndUpdateAICoinData 从aicoin抓取数据
func fetchAndUpdateAICoinData() error {
	url := "https://www.aicoin.com/zh-Hans/currencies/all/1/desc"
	
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头，模拟浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("请求返回状态码: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("解析HTML失败: %v", err)
	}

	var cryptoList []system.SysCrypto
	
	// 解析表格数据
	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		if i >= 20 { // 只取前20条数据
			return
		}

		crypto := parseCryptoRow(s)
		if crypto != nil {
			cryptoList = append(cryptoList, *crypto)
		}
	})

	if len(cryptoList) == 0 {
		return fmt.Errorf("未解析到任何数据")
	}

	// 清空旧数据
	if err := global.GVA_DB.Exec("TRUNCATE TABLE sys_cryptos").Error; err != nil {
		global.GVA_LOG.Error("清空旧数据失败", zap.Error(err))
	}

	// 批量插入新数据
	if err := global.GVA_DB.CreateInBatches(cryptoList, 100).Error; err != nil {
		return fmt.Errorf("批量插入数据失败: %v", err)
	}

	global.GVA_LOG.Info("成功更新加密货币数据", zap.Int("数量", len(cryptoList)))
	return nil
}

// parseCryptoRow 解析单行数据
func parseCryptoRow(s *goquery.Selection) *system.SysCrypto {
	// 解析排名
	rankText := strings.TrimSpace(s.Find("td:nth-child(1)").Text())
	rank, err := strconv.Atoi(rankText)
	if err != nil {
		rank = 0
	}

	// 解析名称和图标
	nameElem := s.Find("td:nth-child(2) .currency-name")
	name := strings.TrimSpace(nameElem.Text())
	
	symbolElem := s.Find("td:nth-child(2) .currency-symbol")
	symbol := strings.TrimSpace(symbolElem.Text())

	// 获取图标URL
	iconElem := s.Find("td:nth-child(2) img")
	icon, _ := iconElem.Attr("src")
	if !strings.HasPrefix(icon, "http") {
		icon = "https://www.aicoin.com" + icon
	}

	// 解析价格
	priceText := strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	price := parsePrice(priceText)

	// 解析24小时涨跌幅
	changeText := strings.TrimSpace(s.Find("td:nth-child(4)").Text())
	change24h := parseChange(changeText)

	// 解析市值
	marketCapText := strings.TrimSpace(s.Find("td:nth-child(5)").Text())
	marketCap := parseMarketCap(marketCapText)

	// 解析24小时交易量
	volumeText := strings.TrimSpace(s.Find("td:nth-child(6)").Text())
	volume24h := parseVolume(volumeText)

	// 解析流通量
	supplyText := strings.TrimSpace(s.Find("td:nth-child(7)").Text())
	circulatingSupply := parseSupply(supplyText)

	if name == "" || symbol == "" {
		return nil
	}

	return &system.SysCrypto{
		Rank:              rank,
		Name:              name,
		Symbol:            symbol,
		Icon:              icon,
		Price:             price,
		Change24h:         change24h,
		MarketCap:         marketCap,
		Volume24h:         volume24h,
		CirculatingSupply: circulatingSupply,
	}
}

// parsePrice 解析价格
func parsePrice(priceText string) float64 {
	// 移除 $ 和 , 符号
	priceText = strings.ReplaceAll(priceText, "$", "")
	priceText = strings.ReplaceAll(priceText, ",", "")
	
	price, err := strconv.ParseFloat(strings.TrimSpace(priceText), 64)
	if err != nil {
		return 0
	}
	return price
}

// parseChange 解析涨跌幅
func parseChange(changeText string) float64 {
	// 移除 % 符号
	changeText = strings.ReplaceAll(changeText, "%", "")
	changeText = strings.TrimSpace(changeText)
	
	// 处理正负号
	sign := 1.0
	if strings.HasPrefix(changeText, "-") {
		sign = -1.0
		changeText = strings.TrimPrefix(changeText, "-")
	} else if strings.HasPrefix(changeText, "+") {
		changeText = strings.TrimPrefix(changeText, "+")
	}
	
	change, err := strconv.ParseFloat(changeText, 64)
	if err != nil {
		return 0
	}
	return change * sign
}

// parseMarketCap 解析市值
func parseMarketCap(marketCapText string) float64 {
	// 移除 $ 和 , 符号
	marketCapText = strings.ReplaceAll(marketCapText, "$", "")
	marketCapText = strings.ReplaceAll(marketCapText, ",", "")
	marketCapText = strings.TrimSpace(marketCapText)
	
	// 处理单位
	multiplier := 1.0
	if strings.HasSuffix(marketCapText, "B") {
		multiplier = 1e9
		marketCapText = strings.TrimSuffix(marketCapText, "B")
	} else if strings.HasSuffix(marketCapText, "M") {
		multiplier = 1e6
		marketCapText = strings.TrimSuffix(marketCapText, "M")
	} else if strings.HasSuffix(marketCapText, "K") {
		multiplier = 1e3
		marketCapText = strings.TrimSuffix(marketCapText, "K")
	}
	
	marketCap, err := strconv.ParseFloat(marketCapText, 64)
	if err != nil {
		return 0
	}
	return marketCap * multiplier
}

// parseVolume 解析交易量
func parseVolume(volumeText string) float64 {
	return parseMarketCap(volumeText) // 同样的解析逻辑
}

// parseSupply 解析流通量
func parseSupply(supplyText string) float64 {
	// 提取数字部分
	re := regexp.MustCompile(`[\d,\.]+`)
	matches := re.FindString(supplyText)
	if matches == "" {
		return 0
	}
	
	// 移除 , 符号
	matches = strings.ReplaceAll(matches, ",", "")
	
	supply, err := strconv.ParseFloat(matches, 64)
	if err != nil {
		return 0
	}
	
	// 处理单位
	if strings.Contains(supplyText, "M") {
		supply *= 1e6
	} else if strings.Contains(supplyText, "B") {
		supply *= 1e9
	} else if strings.Contains(supplyText, "K") {
		supply *= 1e3
	}
	
	return supply
}

// getCryptoStatistics 获取统计数据
func getCryptoStatistics() (CryptoStatistics, error) {
	var stats CryptoStatistics
	
	// 计算总市值和总交易量
	var totalMarketCap, totalVolume float64
	var btcMarketCap float64
	
	var cryptos []system.SysCrypto
	err := global.GVA_DB.Find(&cryptos).Error
	if err != nil {
		return stats, err
	}
	
	for _, crypto := range cryptos {
		totalMarketCap += crypto.MarketCap
		totalVolume += crypto.Volume24h
		
		if strings.ToLower(crypto.Symbol) == "btc" {
			btcMarketCap = crypto.MarketCap
		}
	}
	
	if totalMarketCap > 0 {
		stats.BTCDominance = (btcMarketCap / totalMarketCap) * 100
	}
	
	stats.TotalMarketCap = totalMarketCap
	stats.TotalVolume = totalVolume
	
	// 计算变化率（这里简化处理，实际应该与历史数据比较）
	stats.TotalMarketCapChange = math.Round((rand.Float64()*10-5)*100) / 100
	stats.TotalVolumeChange = math.Round((rand.Float64()*20-10)*100) / 100
	stats.BTCDominanceChange = math.Round((rand.Float64()*2-1)*100) / 100
	
	return stats, nil
}

// generateTrendData 生成趋势数据
func generateTrendData() []float64 {
	data := make([]float64, 24)
	value := 50.0
	
	for i := 0; i < 24; i++ {
		value += (rand.Float64() - 0.5) * 10
		if value < 0 {
			value = 0
		}
		if value > 100 {
			value = 100
		}
		data[i] = value
	}
	
	return data
}