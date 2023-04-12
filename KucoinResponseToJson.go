package main

import (
	"encoding/json"
	"github.com/Kucoin/kucoin-go-sdk"
	"log"
)

func KucoinResponseToJson(resp *kucoin.ApiResponse) ([]byte, error) {
	var data interface{}
	err := json.Unmarshal(resp.RawData, &data)
	if err != nil {
		log.Printf("Error decoding response: %S", err.Error())
		var jsonData []byte
		return jsonData, err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error encoding response: %S", err.Error())
		var jsonData []byte
		return jsonData, err
	}
	return jsonData, err
}
