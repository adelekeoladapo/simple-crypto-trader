package model

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/dapo/crypto-trader/enum/status"
	"time"
)

type Trade struct {
	gorm.Model
	Product             string
	Quantity            float64
	CreatedTime         time.Time
	Price               float64
	EntryTime           time.Time
	EntryPrice          float64
	TradingCost         float64
	CostPrice           float64
	Status              status.Status
	MinimumSellingPrice float64
	SellingPrice        float64
	ExitTime            time.Time
	PercentageGain      float64
}
