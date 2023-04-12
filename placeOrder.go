package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AccountResponse struct {
	Id        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

func createBuyMarketOrder(symbol string, size uint) kucoin.ApiResponse {
	order := &kucoin.CreateOrderModel{}
	order.ClientOid = RandomString(60)
	order.Side = "sell"
	order.Symbol = symbol
	order.Type = "market"
	order.Remark = "Arbitrage Buy"
	order.Size = strconv.FormatUint(uint64(size), 10)

	resp, err := S.CreateOrder(order)
	if err != nil {
		log.Printf("Error: %S", err.Error())
	}

	var response kucoin.ApiResponse
	err = json.Unmarshal([]byte(kucoin.ToJsonString(resp)), &response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return response
}

func createSellMarketOrder(symbol string, size uint) kucoin.ApiResponse {
	order := &kucoin.CreateOrderModel{}
	order.ClientOid = RandomString(60)
	order.Side = "sell"
	order.Symbol = symbol
	order.Type = "market"
	order.Remark = "Arbitrage Buy"
	order.Size = strconv.FormatUint(uint64(size), 10)

	resp, err := S.CreateOrder(order)
	if err != nil {
		log.Printf("Error: %S", err.Error())
	}

	var response kucoin.ApiResponse
	err = json.Unmarshal([]byte(kucoin.ToJsonString(resp)), &response)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return response
}

func RandomString(length int) string {
	var letters = []rune("abcdef0123456789")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Arbitrage(firstPair string, secondPair string, thirdPair string, size uint, mainCur string) {
	first, second := TradeToken(mainCur, firstPair, secondPair)
	// size problem here!
	DirectionTrader(DirectionMaker(mainCur, firstPair), firstPair, size)
	DirectionTrader(DirectionMaker(first, secondPair), secondPair, size)
	DirectionTrader(DirectionMaker(thirdPair, second), thirdPair, size)
}

func SizeMaker() {
	// get amount in wallet
	// do the math for size
}

func GetBalance(symbol string) string {
	for {
		resp, err := S.Accounts(symbol, "trade")
		if err != nil {
			continue
		}

		var response AccountResponse
		error := json.Unmarshal([]byte(kucoin.ToJsonString(resp)), &response)
		if err != nil {
			fmt.Println("Error:", error)
			continue
		}
		return response.Available
	}
}

func DirectionTrader(action string, pair string, size uint) bool {
	switch action {
	case "buy":
		createBuyMarketOrder(pair, size)
	case "sell":
		createSellMarketOrder(pair, size)
	}
	return true
}

func DirectionMaker(mainCurrency string, pair string) string {
	substrings := strings.Split(pair, " ")
	if substrings[0] == mainCurrency {
		return "buy"
	} else {
		return "sell"
	}
}

func TradeToken(mainCurrency string, firstPair string, secondPair string) (string, string) {
	var firstToken string
	var secondToken string

	substrings := strings.Split(firstPair, " ")
	if substrings[0] == mainCurrency {
		firstToken = substrings[1]
	} else {
		firstToken = substrings[2]
	}
	substrings2 := strings.Split(secondPair, " ")
	if substrings[0] == firstToken {
		secondToken = substrings2[1]
	} else {
		secondToken = substrings2[2]
	}
	return firstToken, secondToken
}
