package main

import (
	"sync"
)

const key = "key"
const secret = "secret"
const passphrase = "passphrase"

var S = Auth(key, secret, passphrase)
var wg sync.WaitGroup

func main() {
	//symbols := []string{"BTC-USDT"} // , "LTC-USDT", "NEO-USDT"
	//
	////for {
	////	var token string
	////	_, _ = fmt.Scanln(&token)
	////	if token == "done" {
	////		break
	////	}
	////	symbols = append(symbols, token)
	////}
	//for _, _ = range symbols {
	//	wg.Add(1)
	PublicWebsocketForAllTicker()
	//}
	//wg.Wait()
	//resp := createSellMarketOrder("BTC-USDT", 10)
	//fmt.Println(resp.Message)
	//
}
