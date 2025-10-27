package system

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SystemApi struct{}

// GetSystemConfig
// @Tags      System
// @Summary   获取配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysConfigResponse,msg=string}  "获取配置文件内容,返回包括系统配置"
// @Router    /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
}

// SetSystemConfig
// @Tags      System
// @Summary   设置配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      system.System                   true  "设置配置文件内容"
// @Success   200   {object}  response.Response{data=string}  "设置配置文件内容"
// @Router    /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	err := c.ShouldBindJSON(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = systemConfigService.SetSystemConfig(sys)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// ReloadSystem
// @Tags      System
// @Summary   重载系统
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "重载系统"
// @Router    /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	// 触发系统重载事件
	err := utils.GlobalSystemEvents.TriggerReload()
	if err != nil {
		global.GVA_LOG.Error("重载系统失败!", zap.Error(err))
		response.FailWithMessage("重载系统失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("重载系统成功", c)
}

// GetServerInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	server, err := systemConfigService.GetServerInfo()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}

// GetProfitReport
// @Tags     System
// @Summary  获取收益报表
// @Produce  application/json
// @Param    data  body      ProfitReportRequest  true  "查询参数"
// @Success  200   {object}  response.Response{data=ProfitReportResponse,msg=string}  "获取收益报表成功"
// @Router   /system/profit-report [post]
func (s *SystemApi) GetProfitReport(c *gin.Context) {
	var req ProfitReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取当前用户信息
	claims := utils.GetUserInfo(c)
	if claims == nil {
		response.FailWithMessage("获取用户信息失败", c)
		return
	}

	// 根据类型计算时间范围
	startTime, endTime, err := calculateTimeRange(req.Type, req.StartTime, req.EndTime)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 权限检查：只有管理员可以使用userId参数查询其他用户数据
	var targetUserId uint
	isAdmin := claims.AuthorityId != 888

	if req.UserId != "" {
		// 检查是否为管理员（authority_id != 888）
		if !isAdmin {
			response.FailWithMessage("普通用户无权查询其他用户数据", c)
			return
		}

		// 解析用户ID
		parsedUserId, err := strconv.ParseUint(req.UserId, 10, 32)
		if err != nil {
			response.FailWithMessage("用户ID格式错误", c)
			return
		}
		targetUserId = uint(parsedUserId)

		global.GVA_LOG.Info("管理员查询其他用户收益报表",
			zap.Uint("adminId", claims.BaseClaims.ID),
			zap.Uint("targetUserId", targetUserId))
	} else {
		// 没有指定用户ID时的处理
		if isAdmin {
			// 管理员没有指定用户ID，可以查询所有用户的数据
			targetUserId = 0 // 0 表示查询所有用户
		} else {
			// 普通用户没有指定用户ID，只能查询自己的数据
			targetUserId = claims.BaseClaims.ID
		}
	}

	// 记录查询参数（调试用）
	global.GVA_LOG.Info("查询收益报表",
		zap.String("type", req.Type),
		zap.String("startTime", startTime),
		zap.String("endTime", endTime),
		zap.Uint("authorityId", claims.AuthorityId),
		zap.Uint("currentUserId", claims.BaseClaims.ID),
		zap.String("requestUserId", req.UserId), // 原始请求中的userId
		zap.Uint("targetUserId", targetUserId),
		zap.Bool("isAdmin", isAdmin))

	// 查询数据 - 传入目标用户ID和管理员权限标识
	summary, items, err := getProfitDataFromDB(startTime, endTime, targetUserId, isAdmin)
	if err != nil {
		global.GVA_LOG.Error("查询收益数据失败", zap.Error(err))
		response.FailWithMessage("查询收益数据失败: "+err.Error(), c)
		return
	}

	// 记录查询结果
	global.GVA_LOG.Info("查询收益报表成功",
		zap.Int64("totalTrades", summary.TotalTrades),
		zap.Float64("totalProfit", summary.TotalProfit))

	response.OkWithDetailed(ProfitReportResponse{
		Summary: summary,
		Items:   items,
	}, "查询成功", c)
}

// ProfitReportItem 收益报表数据项
type ProfitReportItem struct {
	ID             int64   `json:"id"`
	UserId         int64   `json:"userId"`
	Exchange       string  `json:"exchange"`
	Pair           string  `json:"pair"`
	OpenDate       string  `json:"openDate"`
	OpenRate       float64 `json:"openRate"`
	Amount         float64 `json:"amount"`
	CloseDate      string  `json:"closeDate"`
	CloseRate      float64 `json:"closeRate"`
	RealizedProfit float64 `json:"realizedProfit"`
}

// ProfitReportSummary 收益报表汇总数据
type ProfitReportSummary struct {
	TotalProfit      float64 `json:"totalProfit"`
	TotalTrades      int64   `json:"totalTrades"`
	ProfitableTrades int64   `json:"profitableTrades"`
	LosingTrades     int64   `json:"losingTrades"`
	WinRate          float64 `json:"winRate"`
	AvgProfit        float64 `json:"avgProfit"`
	MaxProfit        float64 `json:"maxProfit"`
	MaxLoss          float64 `json:"maxLoss"`
}

// ProfitReportRequest 收益报表请求参数
type ProfitReportRequest struct {
	StartTime string `form:"startTime" json:"startTime"`
	EndTime   string `form:"endTime" json:"endTime"`
	Type      string `form:"type" json:"type"`     // today, week, month, all
	UserId    string `form:"userId" json:"userId"` // 管理员查询指定用户ID
}

// ProfitReportResponse 收益报表响应
type ProfitReportResponse struct {
	Summary ProfitReportSummary `json:"summary"`
	Items   []ProfitReportItem  `json:"items"`
}

// calculateTimeRange 根据类型计算时间范围
func calculateTimeRange(timeType, startTime, endTime string) (string, string, error) {
	now := time.Now()

	switch timeType {
	case "today":
		// 当天
		startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
		return startOfDay.Format("2006-01-02 15:04:05"), endOfDay.Format("2006-01-02 15:04:05"), nil

	case "week":
		// 本周
		weekday := now.Weekday()
		if weekday == time.Sunday {
			weekday = 7
		}
		startOfWeek := now.AddDate(0, 0, -int(weekday)+1)
		startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
		endOfWeek := startOfWeek.AddDate(0, 0, 6)
		endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 999999999, endOfWeek.Location())
		return startOfWeek.Format("2006-01-02 15:04:05"), endOfWeek.Format("2006-01-02 15:04:05"), nil

	case "month":
		// 本月
		startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, -1)
		endOfMonth = time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, endOfMonth.Location())
		return startOfMonth.Format("2006-01-02 15:04:05"), endOfMonth.Format("2006-01-02 15:04:05"), nil

	case "custom":
		// 自定义时间
		if startTime == "" || endTime == "" {
			return "", "", fmt.Errorf("自定义时间需要提供开始和结束时间")
		}
		return startTime, endTime, nil

	default:
		// 全部
		return "1970-01-01 00:00:00", now.Format("2006-01-02 15:04:05"), nil
	}
}

// getProfitDataFromDB 从数据库查询收益数据
func getProfitDataFromDB(startTime, endTime string, userId uint, isAdmin bool) (ProfitReportSummary, []ProfitReportItem, error) {
	// 从 GVA_DBList 中获取 freq 数据库连接
	var db *gorm.DB
	var ok bool

	if global.GVA_DBList == nil {
		return ProfitReportSummary{}, nil, fmt.Errorf("数据库连接列表未初始化")
	}

	db, ok = global.GVA_DBList["freq"]
	if !ok {
		return ProfitReportSummary{}, nil, fmt.Errorf("未找到freq数据库配置")
	}

	// 构建查询SQL
	var query string
	var args []interface{}

	// 判断是否为普通用户（authority_id = 888）
	if !isAdmin {
		// 普通用户只能查看自己的数据
		query = `
			SELECT 
				id, user_id, exchange, pair, open_date, open_rate, amount, 
				close_date, close_rate, realized_profit
			FROM trades 
			WHERE is_open = false 
			AND close_date >= ? 
			AND close_date <= ?
			AND user_id = ?
			ORDER BY close_date DESC
		`
		args = []interface{}{startTime, endTime, userId}
	} else {
		// 管理员可以查看所有数据，或者指定用户的数据
		if userId != 0 {
			// 管理员查询指定用户的数据
			query = `
				SELECT 
					id, user_id, exchange, pair, open_date, open_rate, amount, 
					close_date, close_rate, realized_profit
				FROM trades 
				WHERE is_open = false 
				AND close_date >= ? 
				AND close_date <= ?
				AND user_id = ?
				ORDER BY close_date DESC
			`
			args = []interface{}{startTime, endTime, userId}
		} else {
			// 管理员查询所有用户的数据
			query = `
				SELECT 
					id, user_id, exchange, pair, open_date, open_rate, amount, 
					close_date, close_rate, realized_profit
				FROM trades 
				WHERE is_open = false 
				AND close_date >= ? 
				AND close_date <= ?
				ORDER BY close_date DESC
			`
			args = []interface{}{startTime, endTime}
		}
	}

	// 使用原生 SQL 查询
	rows, err := db.Raw(query, args...).Rows()
	if err != nil {
		global.GVA_LOG.Error("查询收益数据失败", zap.Error(err))
		return ProfitReportSummary{}, nil, fmt.Errorf("查询数据失败: %v", err)
	}
	defer rows.Close()

	var items []ProfitReportItem
	var totalProfit float64
	var totalTrades int64
	var profitableTrades int64
	var losingTrades int64
	var maxProfit float64
	var maxLoss float64

	for rows.Next() {
		var item ProfitReportItem
		var userId int64
		var exchange, pair, openDate, closeDate sql.NullString
		var openRate, amount, closeRate, realizedProfit sql.NullFloat64

		err := rows.Scan(
			&item.ID, &userId, &exchange, &pair,
			&openDate, &openRate, &amount,
			&closeDate, &closeRate, &realizedProfit,
		)
		if err != nil {
			global.GVA_LOG.Error("扫描数据行失败", zap.Error(err))
			continue
		}

		// 处理空值
		item.UserId = userId
		item.Exchange = exchange.String
		item.Pair = pair.String
		item.OpenDate = openDate.String
		item.CloseDate = closeDate.String
		if openRate.Valid {
			item.OpenRate = openRate.Float64
		}
		if amount.Valid {
			item.Amount = amount.Float64
		}
		if closeRate.Valid {
			item.CloseRate = closeRate.Float64
		}
		if realizedProfit.Valid {
			item.RealizedProfit = realizedProfit.Float64
		}

		items = append(items, item)
		totalProfit += item.RealizedProfit
		totalTrades++

		if item.RealizedProfit > 0 {
			profitableTrades++
			if item.RealizedProfit > maxProfit {
				maxProfit = item.RealizedProfit
			}
		} else if item.RealizedProfit < 0 {
			losingTrades++
			if item.RealizedProfit < maxLoss {
				maxLoss = item.RealizedProfit
			}
		}
	}

	// 计算汇总数据
	var winRate float64
	var avgProfit float64
	if totalTrades > 0 {
		winRate = float64(profitableTrades) / float64(totalTrades) * 100
		avgProfit = totalProfit / float64(totalTrades)
	}

	summary := ProfitReportSummary{
		TotalProfit:      totalProfit,
		TotalTrades:      totalTrades,
		ProfitableTrades: profitableTrades,
		LosingTrades:     losingTrades,
		WinRate:          winRate,
		AvgProfit:        avgProfit,
		MaxProfit:        maxProfit,
		MaxLoss:          maxLoss,
	}

	return summary, items, nil
}
