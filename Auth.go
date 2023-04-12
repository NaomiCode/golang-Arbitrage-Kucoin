package main

import "github.com/Kucoin/kucoin-go-sdk"

// gets credentials and returns token

func Auth(key string, secret string, passphrase string) *kucoin.ApiService {
	s := kucoin.NewApiService(
		kucoin.ApiKeyOption(key),
		kucoin.ApiSecretOption(secret),
		kucoin.ApiPassPhraseOption(passphrase),
	)
	return s
}
