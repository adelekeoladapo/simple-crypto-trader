package repository

import (
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/model"
)

type TradeRepository interface {

	CreateTrade(trade model.Trade) (model.Trade, error)

	UpdateTrade(trade model.Trade) (model.Trade, error)

	FindTradeById(id int) (model.Trade, error)

	ListTrades(request dto.ListRequest) ([]model.Trade, error)

}
