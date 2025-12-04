package repositories

import (
	"admin-app/orders/commons/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/postgres"
)

type TradeRepository struct{}

func NewTradeRepository() *TradeRepository {
	return &TradeRepository{}
}

func (repository *TradeRepository) TruncateAndLoad(trades []genericModel.Trade) error {
	db := postgres.GetPostGresClient().GormDb
	tx := db.Begin()

	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Exec(constants.TradeTruncateAndLoadQuery).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.CreateInBatches(trades, 100).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
