package binance_client

// Develop binance client
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	bub "github.com/monkeymatt0/binance_url_builder"
)

type Binance struct {
	bub.BinanceURLBuilder
	http.Client
}

func (bc *Binance) New(test bool) {
	bc.BinanceURLBuilder.New(test)

}

func (bc *Binance) KlinesRequest(params map[string]string) ([]RawCandlestick, error) {
	resp, err := http.Get(bc.Klines(params).String())
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var klines [][]interface{}
	rawCandlesticks := []RawCandlestick{}
	if err := json.Unmarshal(body, &klines); err != nil {
		fmt.Println(err)
	}
	temp := &RawCandlestick{}
	for _, kline := range klines {
		rawCandlesticks = append(rawCandlesticks, *temp.New(kline))
	}
	return rawCandlesticks, nil
}
