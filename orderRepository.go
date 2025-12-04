package repositories

import (
	"admin-app/orders/commons/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/postgres"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (repository *OrderRepository) TruncateAndLoad(orders []genericModel.Order) error {
	db := postgres.GetPostGresClient().GormDb

	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Exec(constants.OrderTruncateAndLoadQuery).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.CreateInBatches(orders, 100).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
