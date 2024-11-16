package binance_client

// Develop binance client
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	bub "github.com/monkeymatt0/binance_url_builder"
)

// @todo : Refactor
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

func (bc *Binance) OrderRequest(params map[string]string, apiKey string, secret string, method string) (uint64, error) {
	switch method {
	case http.MethodPost:
		req, err := http.NewRequest(http.MethodPost, bc.Order(params, secret).String(), nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("X-MBX-APIKEY", apiKey)
		resp, err := bc.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		orderResponse := OrderPlacedResponse{}
		if err := json.Unmarshal(body, &orderResponse); err != nil {
			return 0, err
		}
		return orderResponse.OrderId, nil
	case http.MethodDelete:
		req, err := http.NewRequest(http.MethodDelete, bc.Order(params, secret).String(), nil)
		if err != nil {
			return 0, err
		}
		req.Header.Set("X-MBX-APIKEY", apiKey)
		resp, err := bc.Do(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		orderDeletedResponse := OrderDeletedResponse{}
		if err := json.Unmarshal(body, &orderDeletedResponse); err != nil {
			return 0, err
		}
		return orderDeletedResponse.OrderId, nil

	}

	return 0, nil
}

// Need to create an object to represent the returned values
func (bc *Binance) AccountRequest(params map[string]string, apiKey string, secret string) (*AccountInfo, error) {
	req, err := http.NewRequest(http.MethodGet, bc.Account(secret).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-MBX-APIKEY", apiKey)
	resp, err := bc.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	accountInfo := &AccountInfo{}
	if err := json.Unmarshal(body, accountInfo); err != nil {
		return nil, err
	}

	return accountInfo, nil
}
