package main

import (
	"github.com/MartinBechtle/bybit-api/rest"
	"log"
)

func main() {
	baseURL := "https://api.bybit.com/"
	b := rest.New(nil,
		baseURL, "apiKey", "secret",
		true)

	orders, err := b.GetStopOrdersV2("", "", "BTCUSD")
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("orders: %#v", orders)
}
