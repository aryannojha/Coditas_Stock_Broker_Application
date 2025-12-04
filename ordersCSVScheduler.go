package business

import (
	"admin-app/orders/commons/constants"
	"context"
	"fmt"
	"omnenest-backend/src/utils/logger"

	"github.com/robfig/cron/v3"
)

func StartOrdersScheduler(ctx context.Context, runSchedulerNow bool) {
	log := logger.GetLogger(ctx)
	log.Info(constants.OrdersSchedulerStart)

	cr := cron.New(cron.WithSeconds())

	ordersJob := NewGenericJobs(ctx, constants.Orders, constants.OrdersPath, ParseOrdersCSV)
	if _, err := cr.AddJob("0 5 8 * * *", ordersJob); err != nil {
		log.Error(fmt.Sprintf(constants.FailedOrdersJob, err))
		return
	}
	cr.Start()

	go ordersJob.Run()

	if runSchedulerNow {
		go ordersJob.Run()
	}

	go func() {
		<-ctx.Done()
		log.Info(constants.OrdersSchedulerStop)
		cr.Stop()
	}()
}
