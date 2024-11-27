package binance_client

type OrderStatus string

// @remind : For now NEW and PARTIALLY_FILLED are not used, but they are implemented since in future they may be hepful
// The future scenario can be like:
// if PARTIALLY_FILLED then
//
//	cancel order
//
// endif
// Sell in profit what I have bought
// Same logic applies to when selling, of course if done with selling can happen that
// the profit are not exactly what we expect but something less
const (
	NEW              OrderStatus = "NEW"
	FILLED           OrderStatus = "FILLED"
	PARTIALLY_FILLED OrderStatus = "PARTIALLY_FILLED"
)
