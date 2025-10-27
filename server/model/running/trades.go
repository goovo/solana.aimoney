// 自动生成模板Trades
package running

import (
	"time"
)

// 交易报表(模拟盘) 结构体  Trades
type Trades struct {
	Id                 *int        `json:"id" form:"id" gorm:"primarykey;comment:序列;column:id;size:10;"`                                                                                            //序列
	Exchange           *string     `json:"exchange" form:"exchange" gorm:"comment:交易所名称;column:exchange;size:25;"`                                                                                  //交易所
	Pair               *string     `json:"pair" form:"pair" gorm:"comment:市场符号，如 “BTC/USDT:USDT”;column:pair;size:25;"`                                                                             //市场符号
	BaseCurrency       *string     `json:"baseCurrency" form:"baseCurrency" gorm:"comment:pair 中左侧资产代码（BTC）;column:base_currency;size:25;"`                                                         //左侧代码
	StakeCurrency      *string     `json:"stakeCurrency" form:"stakeCurrency" gorm:"comment:pair 中右侧资产代码（USDT）;column:stake_currency;size:25;"`                                                     //右侧代码
	IsOpen             PgBoolAsInt `json:"isOpen" form:"isOpen" gorm:"comment:1 持仓；0 平仓;column:is_open;"`                                                                                           //状态
	FeeOpen            *float64    `json:"feeOpen" form:"feeOpen" gorm:"comment:开仓费率（如 0.001）;column:fee_open;"`                                                                                    //开仓费率
	FeeOpenCost        *float64    `json:"feeOpenCost" form:"feeOpenCost" gorm:"comment:开仓手续费计价货币金额（fee_open × open_trade_value）;column:fee_open_cost;"`                                            //开仓手续费
	FeeOpenCurrency    *string     `json:"feeOpenCurrency" form:"feeOpenCurrency" gorm:"comment:开仓手续费的计价币种;column:fee_open_currency;size:25;"`                                                      //开仓手续费币种
	FeeClose           *float64    `json:"feeClose" form:"feeClose" gorm:"comment:平仓费率;column:fee_close;"`                                                                                          //平仓费率
	FeeCloseCost       *float64    `json:"feeCloseCost" form:"feeCloseCost" gorm:"comment:平仓手续费计价货币金额;column:fee_close_cost;"`                                                                      //平仓手续费
	FeeCloseCurrency   *string     `json:"feeCloseCurrency" form:"feeCloseCurrency" gorm:"comment:平仓手续费的计价币种;column:fee_close_currency;size:25;"`                                                   //平仓手续费币种
	OpenRate           *float64    `json:"openRate" form:"openRate" gorm:"comment:实际成交开仓价;column:open_rate;"`                                                                                       //实际开仓价
	OpenRateRequested  *float64    `json:"openRateRequested" form:"openRateRequested" gorm:"comment:挂单/限价单指定的开仓价;column:open_rate_requested;"`                                                      //限价单指定的开仓价
	OpenTradeValue     *float64    `json:"openTradeValue" form:"openTradeValue" gorm:"comment:建仓时的名义价值 = open_rate × amount;column:open_trade_value;"`                                              //开仓价值
	CloseRate          *float64    `json:"closeRate" form:"closeRate" gorm:"comment:实际成交平仓价;column:close_rate;"`                                                                                    //实际平仓价
	CloseRateRequested *float64    `json:"closeRateRequested" form:"closeRateRequested" gorm:"comment:平仓挂单/限价价;column:close_rate_requested;"`                                                       //限价平仓价
	RealizedProfit     *float64    `json:"realizedProfit" form:"realizedProfit" gorm:"comment:已实现的净利润（计价货币）;column:realized_profit;"`                                                               //净利润
	CloseProfit        *float64    `json:"closeProfit" form:"closeProfit" gorm:"comment:相对收益百分比 = realized_profit / stake_amount;column:close_profit;"`                                             //收益百分比
	CloseProfitAbs     *float64    `json:"closeProfitAbs" form:"closeProfitAbs" gorm:"comment:与 realized_profit 相同，保留做历史兼容;column:close_profit_abs;"`                                               //净利润
	StakeAmount        *float64    `json:"stakeAmount" form:"stakeAmount" gorm:"comment:建仓时投入的计价货币本金（已乘杠杆）;column:stake_amount;"`                                                                   //建仓本金(x杠杆)
	MaxStakeAmount     *float64    `json:"maxStakeAmount" form:"maxStakeAmount" gorm:"comment:策略允许的最大本金（动态计算，用于仓位管理）;column:max_stake_amount;"`                                                     //最大本金
	Amount             *float64    `json:"amount" form:"amount" gorm:"comment:成交后的 base_currency 数量（含杠杆）;column:amount;"`                                                                           //Base数量(x杠杆)
	AmountRequested    *float64    `json:"amountRequested" form:"amountRequested" gorm:"comment:原始下单数量;column:amount_requested;"`                                                                   //原始下单数量
	OpenDate           *time.Time  `json:"openDate" form:"openDate" gorm:"comment:开仓时间;column:open_date;"`                                                                                          //开仓时间
	CloseDate          *time.Time  `json:"closeDate" form:"closeDate" gorm:"comment:平仓时间;column:close_date;"`                                                                                       //平仓时间
	StopLoss           *float64    `json:"stopLoss" form:"stopLoss" gorm:"comment:当前止损触发价;column:stop_loss;"`                                                                                       //止损触发价
	StopLossPct        *float64    `json:"stopLossPct" form:"stopLossPct" gorm:"comment:当前止损相对开仓价的百分比;column:stop_loss_pct;"`                                                                       //止损相对开仓价的百分比
	InitialStopLoss    *float64    `json:"initialStopLoss" form:"initialStopLoss" gorm:"comment:首次设置的硬止损价;column:initial_stop_loss;"`                                                               //硬止损价
	InitialStopLossPct *float64    `json:"initialStopLossPct" form:"initialStopLossPct" gorm:"comment:首次设置的硬止损百分比;column:initial_stop_loss_pct;"`                                                   //硬止损百分比
	IsStopLossTrailing *bool       `json:"isStopLossTrailing" form:"isStopLossTrailing" gorm:"comment:1 表示已启用移动止损;column:is_stop_loss_trailing;"`                                                   //移动止损
	MaxRate            *float64    `json:"maxRate" form:"maxRate" gorm:"comment:持仓期间行情的最高;column:max_rate;"`                                                                                        //最高价
	MinRate            *float64    `json:"minRate" form:"minRate" gorm:"comment:持仓期间行情的最低成交价;column:min_rate;"`                                                                                     //最低价
	ExitReason         *string     `json:"exitReason" form:"exitReason" gorm:"comment:平仓触发原因：'roi' / 'stop_loss' / 'sell_signal' / 'force_exit' / 'emergency_sell' 等;column:exit_reason;size:255;"` //平仓原因
	ExitOrderStatus    *string     `json:"exitOrderStatus" form:"exitOrderStatus" gorm:"comment:平仓订单的最终状态（closed/canceled/expired）;column:exit_order_status;size:100;"`                             //平仓状态
	Strategy           *string     `json:"strategy" form:"strategy" gorm:"comment:产生该交易的策略名;column:strategy;size:100;"`                                                                             //策略名称
	EnterTag           *string     `json:"enterTag" form:"enterTag" gorm:"comment:enter_long/enter_short 信号标签;column:enter_tag;size:255;"`                                                          //信号标签
	Timeframe          *int        `json:"timeframe" form:"timeframe" gorm:"comment:K 线周期;column:timeframe;size:10;"`                                                                               // K 线周期
	TradingMode        string      `json:"tradingMode" form:"tradingMode" gorm:"comment:交易模式;column:trading_mode;type:varchar(20);check:trading_mode in ('SPOT','MARGIN','FUTURES')"`               //交易模式
	AmountPrecision    *float64    `json:"amountPrecision" form:"amountPrecision" gorm:"comment:交易对允许的最小数量;column:amount_precision;"`                                                               //最小数量
	PricePrecision     *float64    `json:"pricePrecision" form:"pricePrecision" gorm:"comment:价格步长;column:price_precision;"`                                                                        //价格步长
	PrecisionMode      *int        `json:"precisionMode" form:"precisionMode" gorm:"comment:精度模式枚举值;column:precision_mode;size:10;"`                                                                //精度
	PrecisionModePrice *int        `json:"precisionModePrice" form:"precisionModePrice" gorm:"column:precision_mode_price;size:10;"`                                                                //精度值
	ContractSize       *float64    `json:"contractSize" form:"contractSize" gorm:"comment:仅期货/永续有效，合约乘数;column:contract_size;"`                                                                     //合约乘数
	Leverage           *float64    `json:"leverage" form:"leverage" gorm:"comment:杠杆倍数;column:leverage;"`                                                                                           //杠杆倍数
	IsShort            PgBoolAsInt `json:"isShort" form:"isShort" gorm:"comment:0=多头；1=空头;column:is_short;"`                                                                                        //0多１空
	LiquidationPrice   *float64    `json:"liquidationPrice" form:"liquidationPrice" gorm:"comment:当前仓位的预估强平价;column:liquidation_price;"`                                                            //预估强平价
	InterestRate       *float64    `json:"interestRate" form:"interestRate" gorm:"comment:资金费率快照（用于计算 funding_fee）;column:interest_rate;"`                                                          //资金费率快照
	FundingFees        *float64    `json:"fundingFees" form:"fundingFees" gorm:"comment:已支付/收到的累计资金费用;column:funding_fees;"`                                                                        //累计资金费用
	FundingFeeRunning  *float64    `json:"fundingFeeRunning" form:"fundingFeeRunning" gorm:"comment:当前持仓尚未结算的预估资金费用（实时更新）;column:funding_fee_running;"`                                             //预估资金费用
	RecordVersion      *int        `json:"recordVersion" form:"recordVersion" gorm:"comment:行级乐观锁，用于并发更新时检测冲突;column:record_version;size:10;"`                                                      //行级乐观锁
}

// TableName 交易报表(模拟盘) Trades自定义表名 trades
func (Trades) TableName() string {
	return "trades"
}
