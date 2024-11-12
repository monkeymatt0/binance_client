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
	rc.CloseTime = time.Unix(int64(closeTime), 0)
	rc.QuoteAssetVolume, _ = strconv.ParseFloat(kline[7].(string), 64)
	rc.NumberOfTrades = kline[8].(float64)
	rc.TakerBuyBaseAsset, _ = strconv.ParseFloat(kline[9].(string), 64)
	rc.TakerBuyQuoteAsset, _ = strconv.ParseFloat(kline[10].(string), 64)
	return rc
}
