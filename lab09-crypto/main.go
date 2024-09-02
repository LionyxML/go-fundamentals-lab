package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

type Rate struct {
	Currency string
	Price    float64
}

type CEXResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"`
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

func GetRate(currency string) (*Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required; %d received", len(currency))
	}

	upCurrency := strings.ToUpper(currency)

	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}

	rate := Rate{Currency: currency, Price: response.Bid}

	return &rate, nil

}

func getCurrencyData(currency string) {
	rate, err := GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}

func main() {
	currencies := []string{"BTC", "ETH", "SOL"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}

	wg.Wait()
}
