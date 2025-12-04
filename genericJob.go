package business

import (
	"admin-app/orders/commons/constants"
	"context"
	"errors"
	"fmt"
	utils "omnenest-backend/src/utils/gchatAlerts"
	"omnenest-backend/src/utils/logger"
	"os"
	"time"
)

var (
	ErrParse = errors.New("parse_error")
	ErrDB    = errors.New("db_error")
)

// ParserFunc is the signature for CSV parsing functions
type ParseFunc func(filePath string) (int, int, error)

type GenericJob struct {
	ctx       context.Context
	name      string
	filePath  string
	parseFunc ParseFunc
}

func NewGenericJobs(ctx context.Context, name, filePath string, parseFunc ParseFunc) *GenericJob {
	return &GenericJob{
		ctx:       ctx,
		name:      name,
		filePath:  filePath,
		parseFunc: parseFunc,
	}
}

func (job *GenericJob) Run() {
	log := logger.GetLogger(job.ctx)
	log.Info(fmt.Sprintf(constants.TriggerTime, time.Now()))
	startTime := time.Now()

	// Checking if the file is present or not
	if _, err := os.Stat(job.filePath); os.IsNotExist(err) {
		log.Error(fmt.Sprintf(constants.FileNotFound, job.name, job.filePath))
		utils.SendGChatErrorAlert(log, fmt.Sprintf(constants.JobErr, job.name), constants.FileMissing)
		return
	}

	// Checking the parsing of file
	totalRecords, failedRecords, err := job.parseFunc(job.filePath)
	if err != nil {
		if errors.Is(err, ErrParse) {
			log.Error(fmt.Sprintf(constants.GeneralParseFailed, job.name, err))
			utils.SendGChatErrorAlert(log, fmt.Sprintf(constants.JobErr, job.name), constants.ParseFailedErr)
			return
		}
		if errors.Is(err, ErrDB) {
			log.Error(fmt.Sprintf(constants.GeneralDatabaseErr, job.name, err))
			utils.SendGChatErrorAlert(log, fmt.Sprintf(constants.JobErr, job.name), constants.DataBaseErr)
			return
		}
		log.Error(fmt.Sprintf(constants.GeneralUnknownErr, job.name, err))
		utils.SendGChatErrorAlert(log, fmt.Sprintf(constants.JobErr, job.name), constants.UnknownErr)
		return
	}

	// Giving success message
	duration := time.Since(startTime)
	log.Info(fmt.Sprintf(constants.GeneralParseSuccess, job.name))
	utils.SendGChatSuccessAlert(log, fmt.Sprintf(constants.JobErr, job.name), duration, totalRecords, failedRecords)
}

func RunManualJob(ctx context.Context, jobType string) {
	log := logger.GetLogger(ctx)

	switch jobType {
	case constants.ConOrders:
		NewGenericJobs(ctx, constants.Orders, constants.OrdersPath, ParseOrdersCSV).Run()

	case constants.ConTrades:
		NewGenericJobs(ctx, constants.Trades, constants.TradesPath, ParseTradesCSV).Run()

	case constants.Holdings:
		NewGenericJobs(ctx, constants.Holdings, constants.HoldingsPath, ParseHoldingEDIS).Run()

	default:
		log.Warn(fmt.Sprintf(constants.UnknownJobType, jobType))
	}
}
