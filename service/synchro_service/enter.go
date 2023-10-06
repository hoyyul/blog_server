package synchro_service

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

func CronInit() {
	timezone, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	cron := gocron.NewScheduler(timezone)
	cron.Cron("0 0 0 * *").Do(SyncArticleData)
	cron.Cron("0 0 0 * *").Do(SyncCommentData)
	cron.StartBlocking()
}
