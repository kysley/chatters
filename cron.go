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
	cronHandler.AddFunc("@midnight", WriteCache)
}
