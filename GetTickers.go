package main

import (
	"log"
)

func GetTickers(symbol string) []byte {
	resp, err := S.Tickers()
	if err != nil {
		log.Printf("Error: %S", err.Error())
	}

	jsonData, err := KucoinResponseToJson(resp)
	if err != nil {
		log.Printf("Error encoding response: %S", err.Error())
	}
	return jsonData

}
