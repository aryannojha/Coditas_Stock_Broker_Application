package business

import (
	"admin-app/search/models"
	"admin-app/search/repositories"
	"context"
	// "errors"
	"fmt"
	genericConstants "omnenest-backend/src/constants"
	"omnenest-backend/src/utils/mapStruct"

	"omnenest-backend/src/utils/postgres"
	"omnenest-backend/src/utils/tracer"

	// "gorm.io/gorm"
)

type OrderColumnCustomGetService struct {
	repository repositories.OrderColumnGetRepository
}

func NewOrderColumnGetCustomService(repository repositories.OrderColumnGetRepository) *OrderColumnCustomGetService {
	return &OrderColumnCustomGetService{
		repository: repository,
	}
}

func (service *OrderColumnCustomGetService) ServiceOrderColumnGet(ctx context.Context, spanCtx context.Context, bffOrderColumnCustomGetRequest models.BFFOrderColumnCustomGetRequest) ([]models.BFFOrderColumnCustomGetResponse, error) {
	childSpanCtx, span := tracer.AddToSpan(spanCtx, "ServiceOrderColumnGet")
	defer func() {
		if span != nil {
			span.End()
		}
	}()

	client := postgres.GetPostGresClient()
	db := client.GormDb

	columnData, err := service.repository.GetColumnIdForUser(childSpanCtx, db, bffOrderColumnCustomGetRequest)
	if err != nil {
		return nil, fmt.Errorf(genericConstants.DatabaseQueryError, err)
	}
	
	responseData := make([]models.BFFOrderColumnCustomGetResponse, len(columnData))
	mapStruct.MapStructArray(childSpanCtx, columnData, responseData)
	return responseData, nil
}