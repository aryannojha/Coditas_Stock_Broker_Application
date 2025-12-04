package repositories

import (
	"admin-app/orders/commons/constants"
	genericModels "omnenest-backend/src/models"
	"omnenest-backend/src/utils/postgres"

	"gorm.io/gorm/clause"
)

type HoldingsEDISRepository struct{}

func NewHoldingsEDISRepository() *HoldingsEDISRepository {
	return &HoldingsEDISRepository{}
}

func (repository *HoldingsEDISRepository) SaveFailedHoldingRowsInBatch(rows []genericModels.HoldingsFailed) error {
	if len(rows) == 0 {
		return nil
	}
	db := postgres.GetPostGresClient().GormDb

	return db.Create(&rows).Error
}

// Upsert the data
func (repository *HoldingsEDISRepository) UpsertHoldings(holdings []genericModels.Holdings) error {
	db := postgres.GetPostGresClient().GormDb

	err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: constants.AccountID},
			{Name: constants.Isin},
			{Name: constants.ProductCode}},
		DoUpdates: clause.AssignmentColumns([]string{
			constants.AveragePrice,
			constants.CuspaAvailableQuantity,
			constants.HoldingAvailableQuantity,
			constants.T1AvailableQuantity,
			constants.CollateralAvailableQuantity,
			constants.MtfAvailableQuantity,
		}),
	}).CreateInBatches(holdings, 400).Error

	return err
}

// Truncate and Load the Data
// func (repo *HoldingsEDISRepository) TruncateAndInsert(holdings []genericModel.Holdings) error {
//     db := postgres.GetPostGresClient().GormDb
//     tx := db.Begin()
//     if tx.Error != nil {
//         return tx.Error
//     }

//     if err := tx.Exec(constants.HoldingsEDISTruncateQuery).Error; err != nil {
//         tx.Rollback()
//         return err
//     }

//     if len(holdings) > 0 {
//         if err := tx.CreateInBatches(&holdings, 100).Error; err != nil {
//             tx.Rollback()
//             return err
//         }
//     }

//     return tx.Commit().Error
// }
