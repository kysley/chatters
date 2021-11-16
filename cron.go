package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

var (
	cronHandler = cron.New(cron.WithLocation(time.UTC))
)

func StartCron() {
	cronHandler.AddFunc("@every 10m", func() { CacheWrite() })
	go cronHandler.Start()
}
