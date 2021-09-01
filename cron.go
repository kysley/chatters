package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

var (
	location, _ = time.LoadLocation("America/New_York")
	cronHandler = cron.New(cron.WithLocation(location))
)

func StartCron() {
	cronHandler.AddFunc("0 0 0 * * *", func() { WriteCache() })
	go cronHandler.Start()
}
