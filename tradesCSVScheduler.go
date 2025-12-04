package business

import (
	"admin-app/orders/commons/constants"
	"context"
	"fmt"
	"omnenest-backend/src/utils/logger"

	"github.com/robfig/cron/v3"
)

func StartTradesScheduler(ctx context.Context, runSchedulerNow bool) {
	log := logger.GetLogger(ctx)
	log.Info(constants.TradesSchedulerStart)

	cr := cron.New(cron.WithSeconds())

	tradesJob := NewGenericJobs(ctx, constants.Trades, constants.TradesPath, ParseTradesCSV)
	if _, err := cr.AddJob("0 10 8 * * *", tradesJob); err != nil {
		log.Error(fmt.Sprintf(constants.FailedTradesJob, err))
		return
	}
	cr.Start()

	go tradesJob.Run()

	if runSchedulerNow {
		go tradesJob.Run()
	}

	go func() {
		<-ctx.Done()
		log.Info(constants.TradesSchedulerStop)
		cr.Stop()
	}()
}
