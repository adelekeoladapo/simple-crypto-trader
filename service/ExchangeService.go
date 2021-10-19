package service

import "gitlab.com/dapo/crypto-trader/dto"

type ExchangeService interface {

	GetPositions() ([]dto.PositionDto, error)

	PlaceBuyOrder(symbol string, quantity float64) error

	PlaceSellOrder(symbol string, quantity float64) error
}
