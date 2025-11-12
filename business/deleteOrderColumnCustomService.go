package business

import (
	"context"
	"errors"
	"fmt"

	"admin-app/search/commons/constants"
	"admin-app/search/models"
	"admin-app/search/repositories"

	genericConstants "omnenest-backend/src/constants"

	"gorm.io/gorm"

	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"
)

type OrderColumnDeleteCustomService struct {
	repository repositories.OrderColumnDeleteRepository
}

func NewOrderColumnDeleteCustomService(repository repositories.OrderColumnDeleteRepository) *OrderColumnDeleteCustomService {
	return &OrderColumnDeleteCustomService{
		repository: repository,
	}
}

func (service *OrderColumnDeleteCustomService) ServiceOrderColumnDelete(ctx context.Context, spanCtx context.Context, bffOrderColumnDeleteRequest models.BFFOrderColumnDeleteRequest) error {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "ServiceOrderColumnDelete")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	client := postgres.GetPostGresClient()
	db := client.GormDb

	var err error
	switch bffOrderColumnDeleteRequest.AdvanceOrderType {
	case constants.OrderBook:
		err = service.deleteForOrderBook(childSpanCtx, db, bffOrderColumnDeleteRequest)
	case constants.TradeBook:
		err = service.deleteForTradeBook(childSpanCtx, db, bffOrderColumnDeleteRequest)
	default:
		return fmt.Errorf(constants.InvalidAction)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(genericConstants.NoDataFoundError)
		}
		return errors.New(genericConstants.InternalServerError)
	}
	return nil
}

func (service *OrderColumnDeleteCustomService) deleteForOrderBook(ctx context.Context, db *gorm.DB, bffOrderColumnDeleteRequest models.BFFOrderColumnDeleteRequest) error {
	err := service.repository.DeleteForOrderBook(ctx, db, bffOrderColumnDeleteRequest)
	if err != nil {
		return err
	}
	return nil
}

func (service *OrderColumnDeleteCustomService) deleteForTradeBook(ctx context.Context, db *gorm.DB, bffOrderColumnDeleteRequest models.BFFOrderColumnDeleteRequest) error {
	err := service.repository.DeleteForTradeBook(ctx, db, bffOrderColumnDeleteRequest)
	if err != nil {
		return err
	}
	return nil
}
