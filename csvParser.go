package business

import (
	"admin-app/orders/commons/constants"
	"admin-app/orders/repositories"
	"encoding/csv"
	"fmt"
	genericModel "omnenest-backend/src/models"
	utils "omnenest-backend/src/utils/gchatAlerts"
	"omnenest-backend/src/utils/logger"
	"os"
	"strconv"
	"strings"
	"time"
)

func makeHoldingFail(row []string, accountID, isin, productCode, reason string) genericModel.HoldingsFailed {
	return genericModel.HoldingsFailed{
		AccountID:   accountID,
		ISIN:        isin,
		ProductCode: productCode,
		FailedRow:   strings.Join(row, ","),
		Reason:      reason,
	}
}

func handleCSVReadAndHeader(filePath string, expectedHeader []string, loggerTag string) ([][]string, int, error) {
	log := logger.GetLoggerWithoutContext()
	invalidCount := 0

	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, fmt.Errorf(constants.OpeningFileErr, loggerTag, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, 0, fmt.Errorf(constants.ReadingContentErr, loggerTag, err)
	}

	if len(records) == 0 {
		log.Warn(fmt.Sprintf(constants.EmptyCSVErr, loggerTag))
		return [][]string{}, 0, nil
	}

	for i, header := range expectedHeader {
		if i >= len(records[0]) || strings.TrimSpace(records[0][i]) != header {
			return nil, 0, fmt.Errorf(constants.ConInvalidHeader, ErrParse, loggerTag)
		}
	}

	var validRecords [][]string
	validRecords = append(validRecords, records[0])

	for lineNo, row := range records[1:] {
		if len(row) != len(expectedHeader) {
			invalidCount++
			continue
		}

		isEmpty := true
		for _, f := range row {
			if strings.TrimSpace(f) != "" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			invalidCount++
			log.Warn(fmt.Sprintf(constants.SkipLineAllEmpty, lineNo+2))
			continue
		}
		validRecords = append(validRecords, row)
	}
	return validRecords, invalidCount, nil
}

func ParseOrdersCSV(filePath string) (int, int, error) {
	expectedHeader := []string{constants.OrderId, constants.Symbol, constants.Quantity, constants.Price, constants.Status}
	loggerTag := constants.Orders
	log := logger.GetLoggerWithoutContext()
	records, headerInvalid, err := handleCSVReadAndHeader(filePath, expectedHeader, loggerTag)
	if err != nil {
		return 0, 0, err
	}

	if len(records) <= 1 {
		log.Warn(constants.EmptyCSV)
		return 0, 0, nil
	}
	failedRecords := headerInvalid
	var orders []genericModel.Order
	for _, row := range records[1:] {

		orderID, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			failedRecords++
			continue
		}

		quantity, err := strconv.Atoi(row[2])
		if err != nil {
			failedRecords++
			continue
		}

		price, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			failedRecords++
			continue
		}

		orders = append(orders, genericModel.Order{
			OrderID:  orderID,
			Symbol:   row[1],
			Quantity: quantity,
			Price:    price,
			Status:   row[4],
		})
	}

	log.Info(fmt.Sprintf(constants.OrderBulkSave, len(orders)))
	err = repositories.NewOrderRepository().TruncateAndLoad(orders)
	if err != nil {
		return 0, failedRecords + len(orders), fmt.Errorf("%w: %v", ErrDB, err)
	}
	return len(orders), failedRecords, nil
}

func ParseTradesCSV(filePath string) (int, int, error) {
	expectedHeader := []string{constants.TradeId, constants.TradePrice, constants.TradeQuantity, constants.TradeTime}
	loggerTag := constants.Trades
	log := logger.GetLoggerWithoutContext()
	records, headerInvalid, err := handleCSVReadAndHeader(filePath, expectedHeader, loggerTag)
	if err != nil {
		return 0, 0, err
	}

	if len(records) <= 1 {
		log.Warn(constants.EmptyCSV)
		return 0, 0, nil
	}

	failedRecords := headerInvalid
	var trades []genericModel.Trade

	for _, row := range records[1:] {

		tradeID, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			failedRecords++
			continue
		}

		tradePrice, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			failedRecords++
			continue
		}

		tradeQty, err := strconv.Atoi(row[2])
		if err != nil {
			failedRecords++
			continue
		}

		tradeTime, err := time.Parse(time.RFC3339, row[3])
		if err != nil {
			failedRecords++
			continue
		}

		trades = append(trades, genericModel.Trade{
			TradeID:    tradeID,
			TradePrice: tradePrice,
			TradeQty:   tradeQty,
			TradeTime:  tradeTime,
		})
	}

	log.Info(fmt.Sprintf(constants.TradeBulkSave, len(trades)))
	err = repositories.NewTradeRepository().TruncateAndLoad(trades)
	if err != nil {
		utils.SendGChatErrorAlert(log, loggerTag, constants.TruncateAndLoadFailed)
		return 0, failedRecords + len(trades), fmt.Errorf("%w: %v", ErrDB, err)
	}
	return len(trades), failedRecords, nil
}

func ParseHoldingEDIS(filePath string) (int, int, error) {
	expectedHeader := []string{constants.AccountID, constants.Isin, constants.AveragePrice, constants.CuspaAvailableQuantity, constants.HoldingAvailableQuantity, constants.T1AvailableQuantity, constants.CollateralAvailableQuantity, constants.MtfAvailableQuantity, constants.ProductCode}
	loggerTag := constants.Holdings
	fileName := constants.WebHoldings
	log := logger.GetLoggerWithoutContext()

	batch := make([]genericModel.Holdings, 0, constants.Batchsize)

	record, headerInvalid, err := handleCSVReadAndHeader(filePath, expectedHeader, loggerTag)
	if err != nil {
		return 0, 0, err
	}

	if len(record) == 0 {
		log.Warn(constants.EmptyCSV)
		return 0, 0, nil
	}

	failedRecords := headerInvalid
	var holdings []genericModel.Holdings
	var failedRows []genericModel.HoldingsFailed

	repository := repositories.NewHoldingsEDISRepository()
	for _, row := range record[1:] {
		accountID := row[0]
		isin := row[1]
		productCodeRaw := row[8]

		avgPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidAvgPrice))
			continue
		}

		cuspaAvailableQuantity, err := strconv.Atoi(row[3])
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidCuspaAvailableQuantity))
			continue
		}

		holdingAvailableQuantity, err := strconv.Atoi(row[4])
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidHoldingAvailableQuantity))
			continue
		}

		t1AvailableQuantity, err := strconv.Atoi(row[5])
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidT1AvailableQuantity))
			continue
		}

		collateralAvailableQuantity, err := strconv.Atoi(row[6])
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidCollateralAvailableQuantity))
			continue
		}

		mtfAvailableQuantity, err := strconv.Atoi(row[7])
		if err != nil {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, constants.InvalidMtfAvailableQuantity))
			continue
		}

		productCode := strings.ToUpper(row[8])
		if productCode != constants.CNC && productCode != constants.MTF {
			failedRecords++
			failedRows = append(failedRows, makeHoldingFail(row, accountID, isin, productCodeRaw, "invalid product_code"))
			continue
		}

		batch = append(batch, genericModel.Holdings{
			AccountID:              row[0],
			ISIN:                   row[1],
			AveragePrice:           avgPrice,
			CuspaAvailableQty:      cuspaAvailableQuantity,
			HoldingAvailableQty:    holdingAvailableQuantity,
			T1AvailableQty:         t1AvailableQuantity,
			CollateralAvailableQty: collateralAvailableQuantity,
			MTFAvailableQty:        mtfAvailableQuantity,
			ProductCode:            productCode,
		})
		if len(batch) == constants.Batchsize {
			if err := repositories.NewHoldingsEDISRepository().UpsertHoldings(batch); err != nil {
				utils.SendGChatErrorAlert(log, fileName, constants.UpsertFailed)
				return 0, failedRecords + len(batch), fmt.Errorf("%w: %v", ErrDB, err)
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := repository.UpsertHoldings(batch); err != nil {
			utils.SendGChatErrorAlert(log, fileName, constants.UpsertFailed)
			return 0, failedRecords + len(batch), fmt.Errorf("%w: %v", ErrDB, err)
		}
	}

	if len(failedRows) > 0 {
		_ = repository.SaveFailedHoldingRowsInBatch(failedRows)
	}

	log.Info(fmt.Sprintf(constants.SavingRecordsUpsert, len(batch)))
	err = repositories.NewHoldingsEDISRepository().UpsertHoldings(holdings)
	if err != nil {
		dbfail := []genericModel.HoldingsFailed{}
		for _, v := range batch {
			dbfail = append(dbfail, genericModel.HoldingsFailed{
				AccountID:   v.AccountID,
				ISIN:        v.ISIN,
				ProductCode: v.ProductCode,
				FailedRow:   fmt.Sprintf("%+v", v),
				Reason:      constants.DataBaseErr,
			})
		}
		_ = repository.SaveFailedHoldingRowsInBatch(dbfail)

		utils.SendGChatErrorAlert(log, fileName, constants.UpsertFailed)
		return 0, failedRecords + len(holdings), fmt.Errorf("%w: %v", ErrDB, err)
	}
	return len(record) - 1 - failedRecords, failedRecords, nil

	// log.Info(fmt.Sprintf(constants.SavingRecordsUpsert, len(holdings)))

	// if err := repositories.NewHoldingsEDISRepository().TruncateAndInsert(holdings); err != nil {
	// 	return fmt.Errorf(constants.FailedToLoad, err)
	// }
	// return len(holdings), nil
}
