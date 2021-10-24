package impl

import (
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/enum/status"
	"gitlab.com/dapo/crypto-trader/model"
	"gitlab.com/dapo/crypto-trader/repository"
	"gitlab.com/dapo/crypto-trader/repository/impl"
	"gitlab.com/dapo/crypto-trader/service"
	"log"
	"time"
)

type TradeServiceImpl struct{
	repository repository.TradeRepository
}

func (TradeServiceImpl *TradeServiceImpl) CreateTrade(request dto.NewTradeRequest) (response dto.TradeResponse, e error) {
	now := time.Now()
	trade := model.Trade{}
	trade.CreatedTime = &now
	trade.UnitPrice = request.CurrentPrice
	trade.Quantity = request.Quantity
	trade.Product = request.Product
	trade.EntryPrice = request.EntryPrice
	trade.MinimumSellingPrice = request.MinimumSellingPrice
	trade.TradingCost = 0  // TODO: find out trading cost
	trade.CostPrice = (trade.EntryPrice * trade.Quantity) + trade.TradingCost
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

func (TradeServiceImpl *TradeServiceImpl) StartTrade(id int) (response dto.TradeResponse, e error) {
	trade, e := TradeServiceImpl.repository.FindTradeById(id)
	if e != nil {
		return
	}
	trade.Status = status.RUNNING
	if trade.EntryTime == nil {
		now := time.Now()
		trade.EntryTime = &now
	}
	trade, e = TradeServiceImpl.repository.UpdateTrade(trade)
	if e != nil {
		return
	}
	response, _ = TradeServiceImpl.generateResponseFromTrade(trade)
	return
}

func (TradeServiceImpl *TradeServiceImpl) generateResponseFromTrade(trade model.Trade) (response dto.TradeResponse, e error) {
	response.Id = trade.ID
	response.Product = trade.Product
	response.Quantity = trade.Quantity
	response.CreatedTime = trade.CreatedTime
	response.UnitPrice = trade.UnitPrice
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


func GetTradeServiceImpl() service.TradeService {
	return &TradeServiceImpl{
		repository: impl.GetTradeRepositoryImpl(),
	}
}