package processor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/service"
	"gitlab.com/dapo/crypto-trader/service/impl"
	"log"
	"strconv"
)

type MarketSimulator struct {
	Name 				string
	StreamUrl 			string
	Product 			string
	RSI_PERIOD 			int
	RSI_OVERBOUGHT 		float64
	RSI_OVERSOLD 		float64
	TRADE_SYMBOL 		string
	TRADE_QUANTITY 		float64
	ExchangeService 	service.ExchangeService
}

func (simulator *MarketSimulator) Start() {
	log.Printf("Connecting to %s market on %s", simulator.Name, simulator.StreamUrl)
	conn, _, err := websocket.DefaultDialer.Dial(simulator.StreamUrl, nil)
	if err != nil {
		log.Panic("Unable to connect to ", simulator.StreamUrl)
	}
	var closes []float64
	for {
		_, message, readErr := conn.ReadMessage()
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		var binanceResponse = dto.BinanceStreamDto{}
		if err := json.Unmarshal(message, &binanceResponse); err != nil {
			fmt.Println("Error: ", err)
		}
		if binanceResponse.Kline.Closed {
			closePrice, _ := strconv.ParseFloat(binanceResponse.Kline.ClosePrice, 64)
			closes = append(closes, closePrice)
			if len(closes) > simulator.RSI_PERIOD {
				rsiService := impl.GetRSIIndicatorService(closes, simulator.RSI_PERIOD)
				indicator, err := rsiService.GetIndicator()
				if err != nil {
					log.Panic("Could not get RSI value. ", err)
					return
				}
				if indicator > simulator.RSI_OVERBOUGHT {
					// TODO: Implement sell order
				} else if indicator < simulator.RSI_OVERSOLD {
					// TODO: Implement buy order
				}
			}
			closes = closes[1:]
		}

	}

}
