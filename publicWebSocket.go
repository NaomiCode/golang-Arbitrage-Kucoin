package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
)

type TickerFromWebsocket struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Sn      string `json:"sn"`
	Topic   string `json:"topic"`
	Subject string `json:"subject"`
	Data    struct {
		BestAsk     string `json:"bestAsk"`
		BestAskSize string `json:"bestAskSize"`
		BestBid     string `json:"bestBid"`
		BestBidSize string `json:"bestBidSize"`
		Price       string `json:"price"`
		Sequence    string `json:"sequence"`
		Size        string `json:"size"`
		Time        int64  `json:"time"`
	} `json:"data"`
}

func PublicWebsocketForAllTicker() {
	rsp, err := S.WebSocketPublicToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}

	c := S.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	channel := kucoin.NewSubscribeMessage("/market/ticker:all", false)
	//uch := kucoin.NewUnsubscribeMessage("/market/ticker:ETH-BTC", false)

	if err := c.Subscribe(channel); err != nil {
		// Handle error
		return
	}
	//var i = 0
	for {
		select {
		case err := <-ec:
			log.Printf("Error: %s", err.Error())
			continue
		case msg := <-mc:
			//fmt.Println("Received: %s", kucoin.ToJsonString(msg))
			var ticker TickerFromWebsocket
			err := json.Unmarshal([]byte(kucoin.ToJsonString(msg)), &ticker)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			//handle response
		}
	}
}

func PublicWebsocketForSybmolTicker(symbol string) {
	rsp, err := S.WebSocketPublicToken()
	if err != nil {
		// Handle error
		return
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err := rsp.ReadData(tk); err != nil {
		// Handle error
		return
	}

	c := S.NewWebSocketClient(tk)

	mc, ec, err := c.Connect()
	if err != nil {
		// Handle error
		return
	}

	channel := kucoin.NewSubscribeMessage(fmt.Sprintf("/market/ticker:%v", symbol), false)
	//uch := kucoin.NewUnsubscribeMessage("/market/ticker:ETH-BTC", false)

	if err := c.Subscribe(channel); err != nil {
		// Handle error
		return
	}
	for {
		select {
		case err := <-ec:
			//c.Stop() // Stop subscribing the WebSocket feed
			log.Printf("Error: %s", err.Error())
			// Handle error
			continue
		case msg := <-mc:
			log.Printf("Received: %s", kucoin.ToJsonString(msg))
		}
	}
}
