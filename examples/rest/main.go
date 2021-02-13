package main

import (
	"github.com/MartinBechtle/bybit-api/rest"
	"log"
	"github.com/google/uuid"
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

	orderId := uuid.New().String()
	order, err := b.CreateOrderV2("Buy", "Limit", 46000, 10., "PostOnly", 49000, 45000, false, false, orderId, "BTCUSD")
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("result: %#v", order)
}
