package main

import (
	"github.com/MartinBechtle/bybit-api/rest"
	"log"
)

func main() {
	baseURL := "https://api.bybit.com/"
	b := rest.New(nil,
		baseURL, "", "",
		true)

	//orders, err := b.GetOrdersV2("", "", "BTCUSD")
	//if err != nil {
	//	log.Printf("%v", err)
	//	return
	//}

	tickers, err := b.GetTicker("BTCUSD")
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("result: %#v", tickers)
}
