package binance_client

import (
	"strconv"
	"time"
)

type RawCandlestick struct {
	OpenTime           time.Time
	OpenPrice          float64
	HighPrice          float64
	LowPrice           float64
	ClosePrice         float64
	Volume             float64
	CloseTime          time.Time
	QuoteAssetVolume   float64
	NumberOfTrades     float64
	TakerBuyBaseAsset  float64
	TakerBuyQuoteAsset float64
}

// This new will receive as array interface, a candlestick
func (rc *RawCandlestick) New(kline []interface{}) *RawCandlestick {
	openTime := kline[0].(float64)
	closeTime := kline[6].(float64)
	rc.OpenTime = time.Unix(int64(openTime/1000), 0)
	rc.OpenPrice, _ = strconv.ParseFloat(kline[1].(string), 64)
	rc.HighPrice, _ = strconv.ParseFloat(kline[2].(string), 64)
	rc.LowPrice, _ = strconv.ParseFloat(kline[3].(string), 64)
	rc.ClosePrice, _ = strconv.ParseFloat(kline[4].(string), 64)
	rc.Volume, _ = strconv.ParseFloat(kline[5].(string), 64)
	rc.CloseTime = time.Unix(int64(closeTime/1000), 0)
	rc.QuoteAssetVolume, _ = strconv.ParseFloat(kline[7].(string), 64)
	rc.NumberOfTrades = kline[8].(float64)
	rc.TakerBuyBaseAsset, _ = strconv.ParseFloat(kline[9].(string), 64)
	rc.TakerBuyQuoteAsset, _ = strconv.ParseFloat(kline[10].(string), 64)
	return rc
}

type OrderPlacedResponse struct {
	Symbol          string    `json:"symbol"`
	OrderId         uint64    `json:"orderId"`
	OrderListId     int64     `json:"orderListId"`
	ClientOrderId   string    `json:"clientOrderId"`
	TransactionTime time.Time `json:"transactionTime"`
}

type OrderDeletedResponse struct {
	Symbol                  string    `json:"symbol"`
	OrigClientOrderId       string    `json:"origClientOrderId"`
	OrderId                 uint64    `json:"orderId"`
	OrderListId             int64     `json:"orderListId"`
	ClientOrderId           string    `json:"clientOrderId"`
	TransactionTime         time.Time `json:"transactTime"`
	Price                   string    `json:"price"`
	OrigQty                 string    `json:"origQty"`
	ExecutedQty             string    `json:"executedQty"`
	CummulativeQuoteQty     string    `json:"cummulativeQuoteQty"`
	Status                  string    `json:"status"`
	TimeInForce             string    `json:"timeInForce"`
	Type                    string    `json:"type"`
	Side                    string    `json:"side"`
	SelfTradePreventionMode string    `json:"selfTradePreventionMode"`
}

type CommissionRate struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInfo struct {
	MakerCommission            uint8          `json:"makerCommission"`
	TakerCommission            uint8          `json:"takerCommission"`
	BuyerCommission            uint8          `json:"buyerCommission"`
	SellerCommission           uint8          `json:"sellerCommission"`
	CommissionRates            CommissionRate `json:"commissionRates"`
	CanTrade                   bool           `json:"canTrade"`
	CanWithdraw                bool           `json:"canWithdraw"`
	CanDeposit                 bool           `json:"canDeposit"`
	Brokered                   bool           `json:"brokered"`
	RequireSelfTradePrevention bool           `json:"requireSelfTradePrevention"`
	PreventSor                 bool           `json:"preventSor"`
	UpdateTime                 uint64         `json:"updateTime"`
	AccountType                string         `json:"accountType"`
	Balances                   []Balance      `json:"balances"`
	Permissions                []string       `json:"permissions"`
	Uid                        uint64         `json:"uid"`
}

type Key struct {
	ListenKey string `json:"listenKey"`
}

type PartialOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type OrderReport struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int    `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	TransactionTime         int    `json:"transactionTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice"`
	WorkingTime             int    `json:"workingTime"`
	IcebergTime             string `json:"icebergTime"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}
type OCOOrderPlaced struct {
	OrderListId     int            `json:"orderListId"`
	ContingencyType string         `json:"contingencyType"`
	ListStatuType   string         `json:"listStatusType"`
	ListOrderStatus string         `json:"listOrderStatus"`
	TransactionTime int            `json:"transactionTime"`
	Symbol          string         `json:"symbol"`
	PartialOrders   []PartialOrder `json:"orders"`
	OrderReports    []OrderReport  `json:"orderRports"`
}

type OrderMessage struct {
	EventType                              string  `json:"e"` // The event type we are interested to is -------- executionReport ----------
	EventTime                              uint64  `json:"E"`
	Symbol                                 string  `json:"s"`
	ClientOrderId                          string  `json:"c"`
	Side                                   string  `json:"S"`
	OrderType                              string  `json:"o"`
	TimeInForce                            string  `json:"f"`
	OrderQuantity                          string  `json:"q"`
	OrderPrice                             string  `json:"p"`
	StopPrice                              string  `json:"P"`
	IcebergQuantity                        string  `json:"F"`
	OrderListId                            int     `json:"g"`
	OriginalOrderClientId                  string  `json:"C"`
	CurrentExecutionType                   string  `json:"x"`
	CurrentOrderStatus                     string  `json:"X"` // This is the one which needed to check the order status
	OrderRejectReason                      string  `json:"r"`
	OrderID                                uint64  `json:"i"` // This is important to understand which order is fired after an OCO order placed, of course only on buy supposing you will have multiple order
	LastExecutedQuantity                   string  `json:"l"`
	CumulativeFilledQuantity               string  `json:"z"`
	LastExecutedPrice                      string  `json:"L"`
	CommissionAmount                       string  `json:"n"`
	CommissionAsset                        *string `json:"N"`
	TransactionTime                        uint64  `json:"T"`
	TradeId                                int     `json:"t"`
	Ignore                                 uint64  `json:"I"`
	IsTheOrderOnTheBook                    bool    `json:"w"`
	IsTheTradeMakerSide                    bool    `json:"m"`
	Ignore2                                bool    `json:"M"`
	OrderCreationTime                      uint64  `json:"O"`
	CumulativeQuoteAssetTransactedQuantity string  `json:"Z"`
	LastQuoteAssetTransactedQuantity       string  `json:"Y"`
	QuoteOrderQuantity                     string  `json:"Q"`
	WorkingTime                            uint64  `json:"W"`
	SelfTradePrevetionMode                 string  `json:"V"`
}
