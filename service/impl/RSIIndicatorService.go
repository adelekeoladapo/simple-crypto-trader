package impl

import (
	"errors"
	"fmt"
	"gitlab.com/dapo/crypto-trader/service"
	"os"
	"strconv"
)

type RSIIndicatorService struct {
	RSI_PERIOD int
	RSI_OVERSOLD float64
	RSI_OVERBOUGHT float64
	TRADE_SYMBOL string
	TRADE_QUANTITY float64
	CLOSE_PRICES []float64
	IN_POSITION bool
}


func (rsiIndicatorService *RSIIndicatorService) Process() (e error) {
	/* ensure there's enough closing prices */
	if len(rsiIndicatorService.CLOSE_PRICES) < rsiIndicatorService.RSI_PERIOD + 1 {
		e = errors.New("insufficient closing prices")
		return
	}
	rsi, e := rsiIndicatorService.getRSI(rsiIndicatorService.CLOSE_PRICES, rsiIndicatorService.RSI_PERIOD)
	if e != nil {
		fmt.Printf("Error: %s", e.Error())
		return
	}

	if rsi > rsiIndicatorService.RSI_OVERBOUGHT {
		if rsiIndicatorService.IN_POSITION {
			fmt.Println("> > > Sell! Sell! Sell!")
		} else {
			fmt.Println("> > > It is overbought but I don't own any. Nothing to sell")
		}
	} else if rsi < rsiIndicatorService.RSI_OVERSOLD {
		if rsiIndicatorService.IN_POSITION {
			fmt.Println("> > > It is oversold, but I already own it. No need to buy")
		} else {
			fmt.Println("> > > Buy! Buy! Buy!")
		}
	}

	return
}

func (rsiIndicatorService *RSIIndicatorService) GetBuyAmount() (amount float64, e error) {
	return
}

func (rsiIndicatorService *RSIIndicatorService) GetSellAmount() (amount float64, e error) {
	return
}

/*
	RSI = 100 - (100/(1+RS))
	RS = AVG GAIN / AVG LOSS
 */
func (rsiIndicatorService *RSIIndicatorService) getRSI(closePrices []float64, period int) (rsi float64, e error) {
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

func (rsiIndicatorService *RSIIndicatorService) sumArray(arr []float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}

func GetRSIIndicatorService(closes []float64) service.IndicatorService {
	period, _ := strconv.ParseInt(os.Getenv("RSI_PERIOD"), 10, 0)
	oversold, _ := strconv.ParseFloat(os.Getenv("RSI_OVERSOLD"), 64)
	overbought, _ := strconv.ParseFloat(os.Getenv("RSI_OVERBOUGHT"), 64)
	symbol := os.Getenv("TRADE_SYMBOL")
	quantity, _ := strconv.ParseFloat(os.Getenv("TRADE_QUANTITY"), 64)

	return &RSIIndicatorService{
		RSI_PERIOD:     int(period),
		RSI_OVERSOLD:   oversold,
		RSI_OVERBOUGHT: overbought,
		TRADE_SYMBOL:   symbol,
		TRADE_QUANTITY: quantity,
		CLOSE_PRICES:   closes,
	}
}
