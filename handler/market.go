package handler

import (
	"context"
	"github.com/lanyanleienen/market/common"
	"github.com/lanyanleienen/market/domain/model"
	"github.com/lanyanleienen/market/domain/service"
	. "github.com/lanyanleienen/market/proto/market"
)

type Market struct {
	MarketService service.IMarketService
}

func (m *Market) CreateMarket(ctx context.Context, request *MarketInfo, response *MarketID) error{
	market := &model.Market{}
	err := common.SwapTo(request, market)
	if err != nil{
		return err
	}
	marketID,err := m.MarketService.CreateMarket(market)
	if err != nil{
		return err
	}
	response.MarketId = marketID
	return nil
}

func (m *Market) FindMarketByID(ctx context.Context, request *MarketID, response *MarketInfo) error{
	market,err := m.MarketService.FindMarketByID(request.MarketId)
	if err != nil{
		return err
	}
	if err = common.SwapTo(market, response); err != nil{
		return nil
	}
	return nil
}