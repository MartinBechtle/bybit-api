package main

import (
	"log"
	"os"

	"github.com/MartinBechtle/bybit-api/rest"
)

func main() {
	baseURL := "https://api.bybit.com/"
	b := rest.New(nil,
		baseURL, os.Getenv("API_KEY"), os.Getenv("API_SECRET"),
		true)

	orders, err := b.GetStopOrdersV2("", "", "BTCUSD")
	if err != nil {
		log.Printf("%v", err)
		return
	}

	//orderId := uuid.New().String()
	//order, err := b.CreateOrderV2("Buy", "Market", 0, 10., "PostOnly", 51000, 47000, false, false, orderId, "BTCUSD")
	//if err != nil {
	//	log.Printf("%v", err)
	//	return
	//}

	log.Printf("result: %#v", orders)
}
