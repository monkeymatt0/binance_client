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

type OrderMessage struct {
	// @todo : Develop the data structure
	/*
		{
		  "e": "executionReport",        // Event type
		  "E": 1499405658658,            // Event time
		  "s": "ETHBTC",                 // Symbol
		  "c": "mUvoqJxFIILMdfAW5iGSOW", // Client order ID
		  "S": "BUY",                    // Side
		  "o": "LIMIT",                  // Order type
		  "f": "GTC",                    // Time in force
		  "q": "1.00000000",             // Order quantity
		  "p": "0.10264410",             // Order price
		  "P": "0.00000000",             // Stop price
		  "F": "0.00000000",             // Iceberg quantity
		  "g": -1,                       // OrderListId
		  "C": "",                       // Original client order ID; This is the ID of the order being canceled
		  "x": "NEW",                    // Current execution type
		  "X": "NEW",                    // Current order status
		  "r": "NONE",                   // Order reject reason; will be an error code.
		  "i": 4293153,                  // Order ID
		  "l": "0.00000000",             // Last executed quantity
		  "z": "0.00000000",             // Cumulative filled quantity
		  "L": "0.00000000",             // Last executed price
		  "n": "0",                      // Commission amount
		  "N": null,                     // Commission asset
		  "T": 1499405658657,            // Transaction time
		  "t": -1,                       // Trade ID
		  "I": 8641984,                  // Ignore
		  "w": true,                     // Is the order on the book?
		  "m": false,                    // Is this trade the maker side?
		  "M": false,                    // Ignore
		  "O": 1499405658657,            // Order creation time
		  "Z": "0.00000000",             // Cumulative quote asset transacted quantity
		  "Y": "0.00000000",             // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
		  "Q": "0.00000000",             // Quote Order Quantity
		  "W": 1499405658657,            // Working Time; This is only visible if the order has been placed on the book.
		  "V": "NONE"                    // selfTradePreventionMode
		}
	*/
}
