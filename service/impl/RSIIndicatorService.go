package impl

import (
	"fmt"
	"gitlab.com/dapo/crypto-trader/service"
	"log"
)

type RSIIndicatorService struct {
	RSI_PERIOD int
	CLOSE_PRICES []float64
}

/*
	RSI = 100 - (100/(1+RS))
	RS = AVG GAIN / AVG LOSS
*/
func (rsiIndicatorService *RSIIndicatorService) GetIndicator() (rsi float64, e error) {
	closePrices := rsiIndicatorService.CLOSE_PRICES
	period := rsiIndicatorService.RSI_PERIOD
	var ups, downs []float64
	for i := 0; i < period; i++ {
		diff := closePrices[i + 1] - closePrices[i]
		if diff > 0 {
			ups = append(ups, diff)
		} else {
			downs = append(downs, diff * -1)
		}
	}
	p := float64(period)
	rs := (rsiIndicatorService.sumArray(ups) / p) / (rsiIndicatorService.sumArray(downs) / p)
	rsi = 100 - (100 / (1 + rs))
	fmt.Println("> > > RSI: ", rsi)
	return
}

func (rsiIndicatorService *RSIIndicatorService) GetBuyAmount() (amount float64, e error) {
	return
}

func (rsiIndicatorService *RSIIndicatorService) GetSellAmount() (amount float64, e error) {
	return
}


func (rsiIndicatorService *RSIIndicatorService) sumArray(arr []float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}

func GetRSIIndicatorService(closes []float64, period int) service.IndicatorService {
	log.Println("RSI Closes: ", closes)
	return &RSIIndicatorService{
		RSI_PERIOD:     int(period),
		CLOSE_PRICES:   closes,
	}
}
