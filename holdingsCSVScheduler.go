package business

import (
	"admin-app/orders/commons/constants"
	"context"
	"fmt"
	"omnenest-backend/src/utils/logger"

	"github.com/robfig/cron/v3"
)

func StartHoldingsScheduler(ctx context.Context, runSchedulerNow bool) {
	log := logger.GetLogger(ctx)
	log.Info(constants.StartScheduler)
	cr := cron.New(cron.WithSeconds())

	holdingsJob := NewGenericJobs(ctx, constants.Holdings, constants.HoldingsPath, ParseHoldingEDIS)
	_, err := cr.AddJob("0 0 8 * * *", holdingsJob)
	if err != nil {
		log.Error(fmt.Sprintf(constants.FailedAddCronJob, err))
		return
	}
	cr.Start()

	go holdingsJob.Run()

	if runSchedulerNow {
		log.Info(constants.RunHoldingsJobNow)
		go holdingsJob.Run()
	}

	go func() {
		<-ctx.Done()
		log.Info(constants.StopScheduler)
		cr.Stop()
	}()
}
