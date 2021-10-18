package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/service/impl"
	"os"
	"strconv"

	//"github.com/pdepip/go-binance/binance"
	"log"
)

//"wss://stream.binance.com:9443/ws/bnbbtc@kline_1m"

func main() {
	/* Load configuration */
	if err := godotenv.Load(); err != nil {
		log.Panic("Could not load app configuration")
	}
	rsiPeriod, _ := strconv.ParseInt(os.Getenv("RSI_PERIOD"), 10, 0)

	//binance := impl.GetBinanceService()



	//url := "wss://stream.binancefuture.com/ws/btcusdt@markPrice"
	url := "wss://stream.binance.com:9443/ws/ethusdt@kline_1m"
	log.Printf("connecting to %s", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Panic("Unable to connect to binance exchange stream")
	}

	var closes []float64

	for {
		_, message, readErr := conn.ReadMessage()
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		var binRes = dto.BinanceStreamDto{}
		err := json.Unmarshal(message, &binRes)
		if err != nil {
			fmt.Printf("Error: %s \n", err)
		}
		if binRes.Kline.Closed {
			closePrice, _ := strconv.ParseFloat(binRes.Kline.ClosePrice, 64)
			closes = append(closes, closePrice)
			log.Println(">>> Closes at ", binRes.Kline.ClosePrice, closes)

			// Let's do RSI

			if len(closes) > int(rsiPeriod) {
				rsiIndicator := impl.GetRSIIndicatorService(closes)
				if err := rsiIndicator.Process(); err != nil {
					fmt.Println("Error occurred while processing RSI: ", err)
				}
				closes = closes[1:]
			}
		}
	}

}


