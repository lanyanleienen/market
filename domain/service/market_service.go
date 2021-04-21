package service

import (
	"github.com/lanyanleienen/market/domain/model"
	"github.com/lanyanleienen/market/domain/repository"
)

type IMarketService interface {
	CreateMarket(*model.Market)(int64,error)
	FindMarketByID(int64)(*model.Market,error)
}

func NewMarketService(marketRepository repository.IMarketRepository) IMarketService{
	return &MarketService{MarketRepository: marketRepository}
}

type MarketService struct {
	MarketRepository repository.IMarketRepository
}

func (m *MarketService) CreateMarket(market *model.Market)(int64,error){
	return m.MarketRepository.CreateMarket(market)
}

func (m *MarketService) FindMarketByID(marketID int64)(*model.Market,error){
	return m.MarketRepository.FindMarketByID(marketID)
}