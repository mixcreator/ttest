package main

import (
	"context"
	"fmt"

	"github.com/aiviaio/go-binance/v2"
)

var (
	apiKey    = "your api key"
	secretKey = "your secret key"
)

type TData struct {
	symbol string
	price  string
}

func get_price(s string, res chan TData) (string, error) {
	//fmt.Println("get_price")
	//fmt.Println("str:", s)
	client := binance.NewClient(apiKey, secretKey)
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, p := range prices {
		if p.Symbol == s {
			//return p.Price, nil
			var d TData
			d.price = p.Price
			d.symbol = p.Symbol
			res <- d
			return "", nil
		}
	}
	return "", nil
}

func main() {

	message := make(chan TData)

	fmt.Println("Hello Go!")
	cl := binance.NewClient(apiKey, secretKey)
	//fmt.Println(cl)
	resss, error := cl.NewExchangeInfoService().Do(context.Background())
	fmt.Println("----------------------------")
	if error != nil {
		fmt.Println(error)
		return
	} else {
		fmt.Println("Ok")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(resss.Symbols[i].Symbol)
		go get_price(resss.Symbols[i].Symbol, message)
	}

	for {
		select {
		case data := <-message:
			//fmt.Println("Res:")
			fmt.Println(data.price, data.symbol)
		}
	}
}
