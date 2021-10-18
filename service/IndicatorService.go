package service

type IndicatorService interface {

	Process() error

	GetBuyAmount() (float64, error)

	GetSellAmount() (float64, error)

}
