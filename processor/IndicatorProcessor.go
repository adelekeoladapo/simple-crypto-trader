package processor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/service"
	"gitlab.com/dapo/crypto-trader/service/impl"
	"log"
	"os"
	"strconv"
)

type IndicatorProcessor struct {
	RSI_PERIOD int
	RSI_OVERBOUGHT float64
	RSI_OVERSOLD float64
	TRADE_SYMBOL string
	TRADE_QUANTITY float64
	IN_POSITION bool
	ExchangeService service.ExchangeService
}

func (indicatorProcessor *IndicatorProcessor) Process() {
	url := "wss://stream.binance.com:9443/ws/ethusdt@kline_1m"
	log.Printf("connecting to %s", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Panic("Unable to connect to binance exchange stream")
	}

	var closes []float64
	inPosition := indicatorProcessor.IN_POSITION
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

			if len(closes) > indicatorProcessor.RSI_PERIOD {
				rsiIndicatorService := impl.GetRSIIndicatorService(closes, indicatorProcessor.RSI_PERIOD)
				rsiIndicatorValue, err := rsiIndicatorService.GetIndicator()
				if err != nil {
					log.Println("Error occurred while trying to get RSI indicator value: ", err)
				} else {
					if rsiIndicatorValue > indicatorProcessor.RSI_OVERBOUGHT {
						if inPosition {
							fmt.Println("> > > Sell! Sell! Sell!")
							if sellErr := indicatorProcessor.ExchangeService.PlaceSellOrder(indicatorProcessor.TRADE_SYMBOL, indicatorProcessor.TRADE_QUANTITY); sellErr != nil {
								log.Println("> > > ERROR: Could not place sell order.", sellErr)
							} else {
								log.Println("> > > Sell order was successfully placed")
								indicatorProcessor.showPosition()
								inPosition = false
							}
						} else {
							fmt.Println("> > > It is overbought but I don't own any. Nothing to sell")
						}
					} else if rsiIndicatorValue < indicatorProcessor.RSI_OVERSOLD {
						if inPosition {
							fmt.Println("> > > It is oversold, but I already own it. No need to buy")
						} else {
							fmt.Println("> > > Buy! Buy! Buy!")
							if buyErr := indicatorProcessor.ExchangeService.PlaceBuyOrder(indicatorProcessor.TRADE_SYMBOL, indicatorProcessor.TRADE_QUANTITY); buyErr != nil {
								log.Println("> > > ERROR: Could not place buy order.", buyErr)
							} else {
								log.Println("> > > Buy order was successfully placed")
								indicatorProcessor.showPosition()
								inPosition = true
							}
						}
					}
				}
				closes = closes[1:]
			}
		}
	}
}


func (indicatorProcessor *IndicatorProcessor) showPosition() {
	positions, err := indicatorProcessor.ExchangeService.GetPositions()
	if err != nil {
		log.Println("> > > Error occurred while getting position. ", err)
	} else {
		log.Println("Position: ", positions)
	}
}




func GetIndicatorProcessor() IndicatorProcessor {
	period, _ := strconv.ParseInt(os.Getenv("RSI_PERIOD"), 10, 0)
	oversold, _ := strconv.ParseFloat(os.Getenv("RSI_OVERSOLD"), 64)
	overbought, _ := strconv.ParseFloat(os.Getenv("RSI_OVERBOUGHT"), 64)
	symbol := os.Getenv("TRADE_SYMBOL")
	quantity, _ := strconv.ParseFloat(os.Getenv("TRADE_QUANTITY"), 64)
	inPosition, _ := strconv.ParseBool(os.Getenv("IN_POSITION"))
	log.Println("Create new instance of RSI Indicator")
	return IndicatorProcessor{
		RSI_PERIOD:     	int(period),
		RSI_OVERBOUGHT: 	overbought,
		RSI_OVERSOLD: 		oversold,
		TRADE_SYMBOL: 		symbol,
		TRADE_QUANTITY: 	quantity,
		IN_POSITION: 		inPosition,
		ExchangeService: 	impl.GetBinanceService(),
	}
}