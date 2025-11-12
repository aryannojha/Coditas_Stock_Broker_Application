package business

import (
	"context"
	"errors"
	"admin-app/search/commons/constants"
	"admin-app/search/models"
	"admin-app/search/repositories"
	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
)

type OrderColumnAddCustomService struct {
	repository repositories.OrderColumnAddRepository
}

func NewOrderColumnAddCustomService(repository repositories.OrderColumnAddRepository) *OrderColumnAddCustomService {
	return &OrderColumnAddCustomService{
		repository: repository,
	}
}

func (service *OrderColumnAddCustomService) ServiceOrderColumnAdd(ctx context.Context, spanCtx context.Context, bffOrderColumnAddRequest models.BFFOrderColumnAddRequest) error {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "ServiceOrderColumnAdd")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	client := postgres.GetPostGresClient()
	db := client.GormDb

	switch bffOrderColumnAddRequest.AdvanceOrderType {
	case constants.OrderBook:
		if err := service.repository.AddColumnForOrderBook(childSpanCtx, db, bffOrderColumnAddRequest); err != nil {
			return err
		}
	case constants.TradeBook:
		if err := service.repository.AddColumnForTradeBook(childSpanCtx, db, bffOrderColumnAddRequest); err != nil {
			return err
		}
	default:
		return errors.New(constants.InvalidAction)
	}
	return nil
}
