package binance_client

// Develop binance client
import (
	"fmt"
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

func (bc *Binance) KlinesRequest(params map[string]string) (*Binance, error) {
	req, err := http.NewRequest(http.MethodGet, bc.Klines(params).String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := bc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return nil, nil
}
