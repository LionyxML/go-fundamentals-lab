package main

import (
	"cryptomaster/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "ETH", "SOL"}

	var wg sync.WaitGroup

	// For go < 1.22
	//
	// for _, currency := range currencies {
	// 	wg.Add(1)
	// 	go func(currencyCode string) {
	// 		getCurrencyData(currencyCode)
	// 		wg.Done()
	// 	}(currency)
	// }
	// or we will have only the last for(ed) currencyCode

	// go >= 1.22
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}

	wg.Wait()
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}
