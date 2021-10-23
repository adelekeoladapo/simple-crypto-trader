package impl

import (
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/enum/status"
	"gitlab.com/dapo/crypto-trader/model"
	"gitlab.com/dapo/crypto-trader/repository"
	"log"
	"time"
)

type TradeServiceImpl struct{
	repository repository.TradeRepository
}

func (TradeServiceImpl *TradeServiceImpl) CreateTrade(request dto.NewTradeRequest) (response dto.TradeResponse, e error) {
	trade := model.Trade{}
	trade.CreatedTime = time.Now()
	trade.Price = request.CurrentPrice
	trade.Quantity = request.Quantity // TODO: validate quantity
	trade.Product = request.Product
	trade.EntryPrice = request.EntryPrice
	trade.MinimumSellingPrice = request.MinimumSellingPrice
	trade.TradingCost = 0  // TODO: find out trading cost
	trade.CostPrice = trade.EntryPrice + trade.TradingCost
	trade.Status = status.PENDING

	trade, e = TradeServiceImpl.repository.CreateTrade(trade)
	if e != nil {
		log.Println("An error occurred. ", e)
		return
	}
	response, _ = TradeServiceImpl.generateResponseFromTrade(trade)
	return
}

//func (TradeServiceImpl *TradeServiceImpl) UpdateTrade(request dto.TradeRequest) (dto.TradeResponse, error) {
//	panic("implement me")
//}

func (TradeServiceImpl *TradeServiceImpl) FindTradeById(id int) (dto.TradeResponse, error) {
	panic("implement me")
}

func (TradeServiceImpl *TradeServiceImpl) ListTrades(request dto.ListRequest) (dto.ListResponse, error) {
	panic("implement me")
}

func (TradeServiceImpl *TradeServiceImpl) generateResponseFromTrade(trade model.Trade) (response dto.TradeResponse, e error) {
	response.Id = trade.ID
	response.Product = trade.Product
	response.Quantity = trade.Quantity
	response.CreatedTime = trade.CreatedTime
	response.Price = trade.Price
	response.EntryTime = trade.ExitTime
	response.EntryPrice = trade.EntryPrice
	response.TradingCost = trade.TradingCost
	response.CostPrice = trade.CostPrice
	response.Status = trade.Status
	response.MinimumSellingPrice = trade.MinimumSellingPrice
	response.SellingPrice = trade.SellingPrice
	response.ExitTime = trade.ExitTime
	response.PercentageGain = trade.PercentageGain
	return
}
