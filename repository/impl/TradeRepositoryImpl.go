package impl

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/dapo/crypto-trader/db"
	"gitlab.com/dapo/crypto-trader/dto"
	"gitlab.com/dapo/crypto-trader/model"
	"gitlab.com/dapo/crypto-trader/repository"
	"sync"
)

type TradeRepositoryImpl struct {
	Database *gorm.DB
}

func (TradeRepositoryImpl *TradeRepositoryImpl) CreateTrade(trade model.Trade) (t model.Trade, e error) {
	if e = TradeRepositoryImpl.Database.Create(&trade).Error; e != nil {
		return
	}
	t = trade
	return
}

func (TradeRepositoryImpl *TradeRepositoryImpl) UpdateTrade(trade model.Trade) (t model.Trade, e error) {
	if e = TradeRepositoryImpl.Database.Save(&trade).Error; e != nil {
		return
	}
	t = trade
	return
}

func (TradeRepositoryImpl *TradeRepositoryImpl) FindTradeById(id int) (t model.Trade, e error) {
	if e = TradeRepositoryImpl.Database.First(&t, id).Error; e != nil {
		return
	}
	return
}

func (TradeRepositoryImpl *TradeRepositoryImpl) ListTrades(request dto.ListRequest) (t []model.Trade, e error) {
	switch request.SortField {
	case "product":
		request.SortField = "product"
	default:
		request.SortField = "id"
	}

	expressionList := TradeRepositoryImpl.Database.Model(&model.Trade{})
	if request.Filter != "" {
		expressionList = expressionList.Where("product LIKE ? ", "%" + request.Filter + "%").Or("status LIKE ? ", "%" + request.Filter + "%")
	}
	expressionList = expressionList.Order(request.SortField + " " + request.SortOrder)
	expressionList = expressionList.Limit(request.Limit).Offset(request.Offset).Find(&t)
	if err := expressionList.Error; err != nil {
		e = err
		return
	}
	return
}


var singleton repository.TradeRepository
var lock sync.Mutex

func GetTradeRepositoryImpl() repository.TradeRepository {
	lock.Lock()
	defer lock.Unlock()
	if singleton == nil {
		singleton = &TradeRepositoryImpl{
			Database: db.Db,
		}
	}
	return singleton
}
