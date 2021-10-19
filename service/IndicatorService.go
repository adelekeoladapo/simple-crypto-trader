package service

type IndicatorService interface {

	GetIndicator() (float64, error)

	GetBuyAmount() (float64, error)

	GetSellAmount() (float64, error)

}
