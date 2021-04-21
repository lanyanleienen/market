package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/lanyanleienen/market/domain/model"
)

type IMarketRepository interface {
	InitTable() error
	CreateMarket(*model.Market)(int64,error)
	FindMarketByID(int64)(*model.Market,error)
}

func NewMarketRepository(db *gorm.DB) IMarketRepository {
	return &MarketRepository{mysqlDb: db}
}

type MarketRepository struct {
	mysqlDb *gorm.DB
}

func (m *MarketRepository) InitTable() error{
	return m.mysqlDb.CreateTable(&model.Market{}).Error
}

func (m *MarketRepository) CreateMarket(market *model.Market)(int64,error){
	return market.ID,m.mysqlDb.Create(&model.Market{}).Error
}

func (m *MarketRepository) FindMarketByID(marketID int64)(market *model.Market,err error){
	market = &model.Market{}
	return market,m.mysqlDb.First(market, market.ID).Error
}
