package fmxt

import "github.com/adshao/go-binance/v2"

type Binance struct {
	client *binance.Client
	test   bool
}

// func NewBinance(key, secret string, test bool, pubsub PubSub, DB db.DB) Binance {
// 	log.Trace().Str("type", "binance").Bool("test", test).Msg("Binance.Init")

// 	binance.UseTestnet = test
// 	client := binance.NewClient(key, secret)

// 	return Binance{client, test, pubsub, DB}
// }
