package main

import (
	"log"
)

func GeMarktPrice(pair string) []byte {
	for {
		resp, err := S.CurrentMarkPrice(pair)
		if err != nil {
			log.Printf("Error: %S", err.Error())
			continue
		}

		jsonData, err := KucoinResponseToJson(resp)
		if err != nil {
			log.Printf("Error encoding response: %S", err.Error())
			continue
		}

		return jsonData
	}
}
