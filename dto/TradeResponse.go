package dto

import (
	"gitlab.com/dapo/crypto-trader/enum/status"
	"time"
)

type TradeResponse struct {
	Id                  uint          `json:"id"`
	Product             string        `json:"product"`
	Quantity            float64       `json:"quantity"`
	CreatedTime         *time.Time    `json:"createdTime"`
	UnitPrice           float64       `json:"unitPrice"`
	EntryTime           *time.Time    `json:"entryTime"`
	EntryPrice          float64       `json:"entryPrice"`
	TradingCost         float64       `json:"tradingCost"`
	CostPrice           float64       `json:"costPrice"`
	Status              status.Status `json:"status"`
	MinimumSellingPrice float64       `json:"minimumSellingPrice"`
	SellingPrice        float64       `json:"sellingPrice"`
	ExitTime            *time.Time    `json:"exitTime"`
	PercentageGain      float64       `json:"percentageGain"`
}
