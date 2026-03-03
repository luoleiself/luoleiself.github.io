package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type DRun struct {
	name string
}

func (d *DRun) Run() {
	fmt.Printf("%s:%s DRun implement Job interface...\n", time.Now().Local(), d.name)
}

func main() {
	c := cron.New()

	/*
		@yearly (or @annually)	Run once a year, midnight, Jan, 1st					0 0 1 1 *
		@monthly 	          		Run once a month, midnight, first of month	0 0 1 * *
		@weekly									Run once a week, midnight between Sat/Sun		0 0 * * 0
		@daily (or @midnight)		Run once a day, midnight										0 0 * * *
		@hourly									Run once a hour, beginning of hour					0 * * * *
		@every <duration>
	*/

	c.AddJob("*/2 * * * *", &DRun{name: "drun"})
	c.AddFunc("*/3 * * * *", func() {
		fmt.Printf("%s:每 3 分钟执行一次...func\n", time.Now().Local())
	})
	c.Start()

	c.AddFunc("@every 10m", func() {
		fmt.Printf("%s:@every 10m 每10分钟执行一次...func\n", time.Now().Local())
	})

	c.AddFunc("@hourly", func() {
		fmt.Printf("%s:@hourly 每小时执行一次...func\n", time.Now().Local())
	})

	c.AddFunc("*/5 9-12,13-18 * * *", func() {
		fmt.Printf("%s:每天 9-12am,13-18pm 期间的每隔 5 分钟执行一次...func\n", time.Now().Local())
	})

	select {}
}
