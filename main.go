package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/dapo/crypto-trader/processor"

	//"github.com/pdepip/go-binance/binance"
	"log"
)

//"wss://stream.binance.com:9443/ws/bnbbtc@kline_1m"

func main() {
	/* Load configuration */
	if err := godotenv.Load(); err != nil {
		log.Panic("Could not load app configuration")
	}

	p := processor.GetIndicatorProcessor()

	p.Process()

}


