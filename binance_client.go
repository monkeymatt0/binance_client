package binance_client

// Develop binance client
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
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
	defer resp.Body.Close()
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

func (bc *Binance) OrderRequest(params map[string]string, apiKey string, secret string, method string) ([]uint64, error) {
	switch method {
	case http.MethodPost:
		req, err := http.NewRequest(http.MethodPost, bc.Order(params, secret).String(), nil)
		if err != nil {
			return []uint64{0, 0}, err
		}
		req.Header.Set("X-MBX-APIKEY", apiKey)
		resp, err := bc.Do(req)
		if err != nil {
			return []uint64{0, 0}, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []uint64{0, 0}, err
		}
		side := ""
		for key, value := range params {
			if key == "SIDE" {
				side = value
				break
			}
		}
		if side == "BUY" {
			orderResponse := OrderPlacedResponse{}
			if err := json.Unmarshal(body, &orderResponse); err != nil {
				return []uint64{0, 0}, err
			}
			return []uint64{orderResponse.OrderId, 0}, nil
		} else {
			orderResponse := OCOOrderPlaced{}
			if err := json.Unmarshal(body, &orderResponse); err != nil {
				return []uint64{0, 0}, err
			}
			return []uint64{
				uint64(orderResponse.PartialOrders[0].OrderId), // Loss order
				uint64(orderResponse.PartialOrders[1].OrderId), // Profit order
			}, nil
		}
	case http.MethodDelete:
		req, err := http.NewRequest(http.MethodDelete, bc.Order(params, secret).String(), nil)
		if err != nil {
			return []uint64{0, 0}, err
		}
		req.Header.Set("X-MBX-APIKEY", apiKey)
		resp, err := bc.Do(req)
		if err != nil {
			return []uint64{0, 0}, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return []uint64{0, 0}, err
		}
		orderDeletedResponse := OrderDeletedResponse{}
		if err := json.Unmarshal(body, &orderDeletedResponse); err != nil {
			return []uint64{0, 0}, err
		}
		return []uint64{orderDeletedResponse.OrderId, 0}, nil
	}

	return []uint64{0, 0}, nil
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
	defer resp.Body.Close()
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

func (bc *Binance) ListenKeyRequest(apiKey string, orderId uint64) (*Key, error) {
	req, err := http.NewRequest(http.MethodPost, bc.ListenKey(nil).String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-MBX-APIKEY", apiKey)

	resp, err := bc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	listenKey := &Key{}
	if err := json.Unmarshal(body, listenKey); err != nil {
		return nil, err
	}
	return listenKey, nil
}

func (bc *Binance) UserDataStreamSocket(listenKey string, apiKey *string, validTime *time.Duration, profitOrderId, lossOrderId uint64) (bool, error) {
	conn, _, err := websocket.DefaultDialer.Dial(bc.UserDataStream(listenKey), nil)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	start := time.Now()
	if validTime != nil {

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return false, err
			}
			orderMessage := &OrderMessage{}
			if err := json.Unmarshal(message, orderMessage); err != nil {
				return false, err
			}

			if time.Since(start) <= *validTime && orderMessage.CurrentOrderStatus == string(FILLED) { // The has been executed within validTime
				return true, nil
			} else if time.Since(start) > *validTime { // The order exeeded the maximun valid time for it's execution means I should delete the order
				return false, nil
			}

		}
	} else {
		start = start.Add(30 * time.Minute) // When I will call this function for selling order in worst case 25 minutes are passed, in this way the listenKey is directly refreshed
		// Implement keep alive connection
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return false, err
			}
			orderMessage := &OrderMessage{}
			if err := json.Unmarshal(message, orderMessage); err != nil {
				return false, err
			}

			if time.Since(start) >= 30*time.Minute { // Keep alive connection ping for the listen key
				bc.keepAliveListenKey(listenKey, *apiKey)
				return true, nil
			}

			if orderMessage.CurrentOrderStatus == string(FILLED) && orderMessage.OrderID == profitOrderId {
				return true, nil
			} else if orderMessage.CurrentOrderStatus == string(FILLED) && orderMessage.OrderID == lossOrderId {
				return false, nil
			}

		}
	}
}

func (bc *Binance) keepAliveListenKey(listenKey, apiKey string) error {
	req, err := http.NewRequest(http.MethodPut, bc.ListenKey(&listenKey).String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-MBX-APIKEY", apiKey)

	resp, err := bc.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
