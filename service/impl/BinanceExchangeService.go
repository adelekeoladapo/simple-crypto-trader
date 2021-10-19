package impl

import (
	"fmt"
	"github.com/pdepip/go-binance/binance"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/service"
	"os"
)

type BinanceExchangeService struct {
	Client *binance.Binance
}

func (binanceExchangeService *BinanceExchangeService) GetPositions() (positions []dto.PositionDto, e error) {
	res, e := binanceExchangeService.Client.GetPositions()
	if e != nil {
		return
	}
	for _, p := range res {
		position := dto.PositionDto{Asset: p.Asset, Free: p.Free, Locked: p.Locked}
		positions = append(positions, position)
	}
	return
}

func (binanceExchangeService *BinanceExchangeService) PlaceBuyOrder(symbol string, quantity float64) (e error) {
	order := binance.MarketOrder{Symbol: symbol, Side: "BUY", Type: "MARKET", Quantity: quantity}
	res, e := binanceExchangeService.Client.PlaceMarketOrder(order)
	if e != nil {
		return e
	}
	fmt.Println(res)
	return
}

func (binanceExchangeService *BinanceExchangeService) PlaceSellOrder(symbol string, quantity float64) (e error) {
	order := binance.MarketOrder{Symbol: symbol, Side: "SELL", Type: "MARKET", Quantity: quantity}
	res, e := binanceExchangeService.Client.PlaceMarketOrder(order)
	if e != nil {
		return e
	}
	fmt.Println(res)
	return
}

func GetBinanceService() service.ExchangeService {
	return &BinanceExchangeService{
		Client: binance.New(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY")),
	}
}
