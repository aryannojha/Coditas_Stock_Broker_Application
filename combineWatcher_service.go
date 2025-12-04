package business

import (
	"admin-app/orders/commons/constants"
	"context"
	"omnenest-backend/src/utils/logger"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func StartFileWatcher(ctx context.Context) {
	log := logger.GetLogger(ctx)
	log.Info(constants.InitializingWatcher)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(constants.FailedCreatingWatcher + err.Error())
		return
	}
	defer watcher.Close()

	// Watching the directory
	if err := watcher.Add(constants.WatchDir); err != nil {
		log.Error(constants.FailedWatchingDirectory + constants.WatchDir)
		return
	}
	log.Info(constants.WatchingDirectory + constants.WatchDir)

	filesToWatch := []string{constants.OrdersAddress, constants.TradesAddress}

	// Adding existing CSV files if present
	for _, file := range filesToWatch {
		if _, err := os.Stat(file); err == nil {
			if err := watcher.Add(file); err == nil {
				log.Info(constants.WatchingExistingFiles + file)
			}
			if strings.HasSuffix(file, constants.OrdersAddress) {
				log.Info(constants.ExistingOrdersRead + file)
				_, _, err := ParseOrdersCSV(file)
				if err != nil {
					log.Error(constants.ExistingOrdersReadFailed + err.Error())
				}
			} else if strings.HasSuffix(file, constants.TradesAddress) {
				log.Info(constants.ExistingTradesRead + file)
				_, _, err := ParseTradesCSV(file)
				if err != nil {
					log.Error(constants.ExistingTradesReadFailed + err.Error())
				}
			}
		} else {
			log.Warn(constants.FileNotFoundWait + file)
		}
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			name := event.Name
			isOrder := strings.HasSuffix(name, constants.OrdersCSV)
			isTrade := strings.HasSuffix(name, constants.TradesCSV)

			// Handle new file creation
			if event.Op&fsnotify.Create != 0 {
				log.Info(constants.FileCreated + name)
				if isOrder || isTrade {
					time.Sleep(300 * time.Millisecond)
					if err := watcher.Add(name); err != nil {
						log.Warn(constants.WatchingStartFail + name)
					}
				}
			}

			// Handling write or replace
			if event.Op&(fsnotify.Write|fsnotify.Rename) != 0 {
				if _, err := os.Stat(name); err == nil {
					if isOrder {
						log.Info(constants.UpdateDetectionOrders)
						_, _, err := ParseOrdersCSV(name)
						if err != nil {
							log.Error(err.Error())
						}
					} else if isTrade {
						log.Info(constants.UpdateDetectionTrades)
						_, _, err := ParseTradesCSV(name)
						if err != nil {
							log.Error(err.Error())
						}
					} else {
						log.Debug(constants.ChangeInUnrelatedFile + name)
					}
				} else {
					log.Warn(constants.FileTemporarilyUnavailable + name)
					time.Sleep(500 * time.Millisecond)
					if _, err := os.Stat(name); err == nil {
						log.Info(constants.FileReappear + name)
						watcher.Add(name)
					}
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Error(constants.WatcherErr + err.Error())

		case <-ctx.Done():
			log.Info(constants.StoppingCSV)
			return
		}
	}
}
