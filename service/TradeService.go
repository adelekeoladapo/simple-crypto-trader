package service

import "gitlab.com/dapo/crypto-trader/dto"

type TradeService interface {

	CreateTrade(request dto.NewTradeRequest) (dto.TradeResponse, error)

	//UpdateTrade(request dto.TradeRequest) (dto.TradeResponse, error)

	FindTradeById(id int) (dto.TradeResponse, error)

	ListTrades(request dto.ListRequest) (dto.ListResponse, error)

}

