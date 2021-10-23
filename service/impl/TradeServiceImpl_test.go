package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/model"
	"testing"
	"time"
)

func TestTradeServiceImpl_CreateTrade(t *testing.T) {
	// given
	now := time.Now()
	request := dto.NewTradeRequest{
		Id:                  0,
		Quantity:            0,
		Product:             "",
		CurrentPrice:        0,
		EntryPrice:          0,
		MinimumSellingPrice: 0,
	}
	repo := TradeRepositoryMock{}
	repo.On("CreateTrade").Return(model.Trade{
		Model:               gorm.Model{},
		Product:             "",
		Quantity:            0,
		CreatedTime:         &now,
		UnitPrice:           0,
		EntryTime:           nil,
		EntryPrice:          0,
		TradingCost:         0,
		CostPrice:           0,
		Status:              "",
		MinimumSellingPrice: 0,
		SellingPrice:        0,
		ExitTime:            nil,
		PercentageGain:      0,
	}, nil)

	// when
	service := TradeServiceImpl{&repo}
	_, err := service.CreateTrade(request)

	// than
	if err != nil {
		t.Errorf("Error occurred while creating trade. %v", err)
	}

}




type TradeRepositoryMock struct {
	mock.Mock
}

func (r *TradeRepositoryMock) CreateTrade(trade model.Trade) (t model.Trade, e error) {
	now := time.Now()
	args := r.Called()
	t = model.Trade{
		Model:               gorm.Model{},
		Product:             "",
		Quantity:            0,
		CreatedTime:         &now,
		UnitPrice:           0,
		EntryTime:           nil,
		EntryPrice:          0,
		TradingCost:         0,
		CostPrice:           0,
		Status:              "",
		MinimumSellingPrice: 0,
		SellingPrice:        0,
		ExitTime:            nil,
		PercentageGain:      0,
	}
	e = args.Error(1)
	return
}

func (r *TradeRepositoryMock) UpdateTrade(trade model.Trade) (t model.Trade, e error) {
	return
}

func (r *TradeRepositoryMock) FindTradeById(id int) (t model.Trade, e error) {
	return
}

func (r *TradeRepositoryMock) ListTrades(request dto.ListRequest) (t []model.Trade, e error) {
	return
}