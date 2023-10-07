package synchro_service

import (
	"time"

	"github.com/robfig/cron/v3"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Tokyo")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	Cron.AddFunc("0 0 0 * * *", SyncArticleData)
	Cron.AddFunc("0 0 0 * * *", SyncCommentData)
	Cron.Start()
}
