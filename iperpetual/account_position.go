// Position (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-position)
package iperpetual

import (
	"time"

	"github.com/ginarea/gobybit/transport"
)

// My Position (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-myposition)
type GetPosition struct {
	Symbol *string `param:"symbol"`
}

func (this GetPosition) Do(client *Client) ([]PositionItem, bool) {
	if this.Symbol == nil {
		return Get[[]PositionItem](client, "position/list", this)
	}
	r, ok := Get[PositionItem](client, "position/list", this)
	return []PositionItem{r}, ok
}

type PositionBase struct {
	ID                  int               `json:"id"`
	UserID              int               `json:"user_id"`
	RiskID              int               `json:"risk_id"`
	Symbol              string            `json:"symbol"`
	Side                Side              `json:"side"`
	Size                int               `json:"size"`
	PositionValue       transport.Float64 `json:"position_value"`
	EntryPrice          transport.Float64 `json:"entry_price"`
	IsIsolated          bool              `json:"is_isolated"`
	AutoAddMargin       int               `json:"auto_add_margin"`
	Leverage            transport.Float64 `json:"leverage"`
	EffectiveLeverage   transport.Float64 `json:"effective_leverage"`
	PositionMargin      transport.Float64 `json:"position_margin"`
	LiqPrice            transport.Float64 `json:"liq_price"`
	BustPrice           transport.Float64 `json:"bust_price"`
	OccClosingFee       transport.Float64 `json:"occ_closing_fee"`
	OccFundingFee       transport.Float64 `json:"occ_funding_fee"`
	TakeProfit          transport.Float64 `json:"take_profit"`
	StopLoss            transport.Float64 `json:"stop_loss"`
	TrailingStop        transport.Float64 `json:"trailing_stop"`
	PositionStatus      string            `json:"position_status"`
	DeleverageIndicator int               `json:"deleverage_indicator"`
	OcCalcData          string            `json:"oc_calc_data"`
	OrderMargin         transport.Float64 `json:"order_margin"`
	WalletBalance       transport.Float64 `json:"wallet_balance"`
	RealisedPnl         transport.Float64 `json:"realised_pnl"`
	CumRealisedPnl      transport.Float64 `json:"cum_realised_pnl"`
	CrossSeq            int               `json:"cross_seq"`
	PositionSeq         int               `json:"position_seq"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type PositionData struct {
	PositionBase
	PositionIdx   PositionIdx `json:"position_idx"`
	Mode          int         `json:"mode"`
	UnrealisedPnl float64     `json:"unrealised_pnl"`
	TpSlMode      TpSlMode    `json:"tp_sl_mode"`
}

type PositionItem struct {
	Data    PositionData `json:"data"`
	IsValid bool         `json:"is_valid"`
}

func (this *Client) GetPosition(symbol *string) ([]PositionItem, bool) {
	return GetPosition{Symbol: symbol}.Do(this)
}

func (this *Client) GetOnePosition(symbol string) (i PositionItem, ok bool) {
	ret, ok := GetPosition{Symbol: &symbol}.Do(this)
	if ok {
		ok = len(ret) == 1
		if ok {
			i = ret[0]
		}
	}
	return
}

func (this *Client) GetAllPositions() ([]PositionItem, bool) {
	return GetPosition{}.Do(this)
}

// Change Margin (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-changemargin)
type ChangeMargin struct {
	Symbol string `param:"symbol"`
	Margin string `param:"margin"`
}

func (this ChangeMargin) Do(client *Client) (float64, bool) {
	return Post[float64](client, "position/change-position-margin", this)
}

func (this *Client) ChangeMargin(v ChangeMargin) (float64, bool) {
	return v.Do(this)
}

// Set Trading-Stop (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-tradingstop)
type SetTradingStop struct {
	Symbol            string        `param:"symbol"`
	TakeProfit        *int          `param:"take_profit"`
	StopLoss          *int          `param:"stop_loss"`
	TrailingStop      *int          `param:"trailing_stop"`
	TpTrigger         *TriggerPrice `param:"tp_trigger_by"`
	SlTrigger         *TriggerPrice `param:"sl_trigger_by"`
	NewTrailingActive *int          `param:"new_trailing_active"`
	SlSize            *int          `param:"sl_size"`
	TpSize            *int          `param:"tp_size"`
}

func (this SetTradingStop) Do(client *Client) (SetTradingStopResult, bool) {
	return Post[SetTradingStopResult](client, "position/trading-stop", this)
}

type SetTradingStopExt struct {
	TrailingActive string `json:"trailing_active"`
	SlTrigger      string `json:"sl_trigger_by"`
	TpTrigger      string `json:"tp_trigger_by"`
	V              int    `json:"v"`
	Mm             int    `json:"mm"`
}

type SetTradingStopResult struct {
	PositionBase
	CumCommission int               `json:"cum_commission"`
	ExtFields     SetTradingStopExt `json:"ext_fields"`
}

func (this *Client) SetTradingStop(v SetTradingStop) (SetTradingStopResult, bool) {
	return v.Do(this)
}

// Set Leverage (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-setleverage)
type SetLeverage struct {
	Symbol       string `param:"symbol"`
	Leverage     int    `param:"leverage"`
	LeverageOnly *bool  `param:"leverage_only"`
}

func (this SetLeverage) Do(client *Client) (int, bool) {
	return Post[int](client, "position/leverage/save", this)
}

func (this *Client) SetLeverage(v SetLeverage) (int, bool) {
	return v.Do(this)
}

// Full/Partial Position TP/SL Switch (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-switchmode)
//
// Switch mode between Full or Partial
type TpSlModeSwitch struct {
	Symbol   string    `param:"symbol"`
	TpSlMode *TpSlMode `param:"tp_sl_mode"`
}

func (this TpSlModeSwitch) Do(client *Client) (TpSlMode, bool) {
	type result struct {
		TpSlMode TpSlMode `json:"tp_sl_mode"`
	}
	r, ok := Post[result](client, "tpsl/switch-mode", this)
	return r.TpSlMode, ok
}

func (this *Client) TpSlModeSwitch(v TpSlModeSwitch) (TpSlMode, bool) {
	return v.Do(this)
}

// Cross/Isolated Margin Switch (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-marginswitch)
//
// Switch Cross/Isolated; must set leverage value when switching from Cross to Isolated
type MarginSwitch struct {
	Symbol       string `param:"symbol"`
	IsIsolated   bool   `param:"is_isolated"`
	BuyLeverage  int    `param:"buy_leverage"`
	SellLeverage int    `param:"sell_leverage"`
}

func (this MarginSwitch) Do(client *Client) bool {
	_, ok := Post[struct{}](client, "position/switch-isolated", this)
	return ok
}

func (this *Client) MarginSwitch(v MarginSwitch) bool {
	return v.Do(this)
}

// Get User Trade Records (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-usertraderecords)
//
// Get user's trading records. The results are ordered in ascending order (the first item is the oldest)
type GetTradeRecords struct {
	Symbol    string     `param:"symbol"`
	OrderID   *string    `param:"order_id"`
	StartTime *int       `param:"start_time"`
	Page      *int       `param:"page"`
	Limit     *int       `param:"limit"`
	Order     *SortOrder `param:"order"`
}

func (this GetTradeRecords) Do(client *Client) (TradeRecords, bool) {
	return Get[TradeRecords](client, "execution/list", this)
}

type TradeRecord struct {
	ClosedSize    int       `json:"closed_size"`
	CrossSeq      int       `json:"cross_seq"`
	ExecFee       string    `json:"exec_fee"`
	ExecID        string    `json:"exec_id"`
	ExecPrice     string    `json:"exec_price"`
	ExecQty       int       `json:"exec_qty"`
	ExecTime      int64     `json:"exec_time"`
	ExecType      ExecType  `json:"exec_type"`
	ExecValue     string    `json:"exec_value"`
	FeeRate       string    `json:"fee_rate"`
	LastLiquidity string    `json:"last_liquidity_ind"`
	LeavesQty     int       `json:"leaves_qty"`
	NthFill       int       `json:"nth_fill"`
	OrderID       string    `json:"order_id"`
	OrderLinkID   string    `json:"order_link_id"`
	OrderPrice    string    `json:"order_price"`
	OrderQty      int       `json:"order_qty"`
	OrderType     OrderType `json:"order_type"`
	Side          Side      `json:"side"`
	Symbol        string    `json:"symbol"`
	UserID        int       `json:"user_id"`
	TradeTime     uint64    `json:"trade_time_ms"`
}

type TradeRecords struct {
	OrderID      string        `json:"order_id"`
	TradeRecords []TradeRecord `json:"trade_list"`
}

func (this *Client) GetTradeRecords(v GetTradeRecords) (TradeRecords, bool) {
	return v.Do(this)
}

// Closed Profit and Loss (https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-closedprofitandloss)
type ClosedProfitLoss struct {
	Symbol    string    `param:"symbol"`
	StartTime *int      `param:"start_time"`
	EndTime   *int      `param:"end_time"`
	ExecType  *ExecType `param:"exec_type"`
	Page      *int      `param:"page"`
	Limit     *int      `param:"limit"`
}

func (this ClosedProfitLoss) Do(client *Client) (ClosedProfitLossResult, bool) {
	return Get[ClosedProfitLossResult](client, "trade/closed-pnl/list", this)
}

type ClosedData struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	OrderID       string    `json:"order_id"`
	Side          Side      `json:"side"`
	Qty           float64   `json:"qty"`
	OrderPrice    float64   `json:"order_price"`
	OrderType     OrderType `json:"order_type"`
	ExecType      ExecType  `json:"exec_type"`
	ClosedSize    float64   `json:"closed_size"`
	CumEntryValue float64   `json:"cum_entry_value"`
	AvgEntryPrice float64   `json:"avg_entry_price"`
	CumExitValue  float64   `json:"cum_exit_value"`
	AvgExitPrice  float64   `json:"avg_exit_price"`
	ClosedPnl     float64   `json:"closed_pnl"`
	FillCount     int       `json:"fill_count"`
	Leverage      int       `json:"leverage"`
	CreatedAt     uint64    `json:"created_at"`
}

type ClosedProfitLossResult struct {
	CurrentPage int          `json:"current_page"`
	Data        []ClosedData `json:"data"`
}

func (this *Client) ClosedProfitLoss(v ClosedProfitLoss) (ClosedProfitLossResult, bool) {
	return v.Do(this)
}
