package dto

type NewTradeRequest struct {
	Id 							int
	Quantity 					float64					`validate:"required,min=0"`
	Product 					string					`validate:"required"`
	CurrentPrice 				float64					`validate:"required,min=0"`
	EntryPrice 					float64					`validate:"required,min=0"`
	MinimumSellingPrice 		float64					`validate:"required,min=0"`
}
