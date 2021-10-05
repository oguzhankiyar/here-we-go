package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("0 30 * * * *", func() {
		fmt.Println(time.Now(), "Every hour on the half hour")
	})

	c.AddFunc("@hourly", func() {
		fmt.Println(time.Now(), "Every hour")
	})

	c.AddFunc("@every 1h30m", func() {
		fmt.Println(time.Now(), "Every hour thirty")
	})

	c.AddFunc("@every 10s", func() {
		fmt.Println(time.Now(), "Every 10 seconds")
	})

	c.Start()

	// Func may also be added to a running Cron
	c.AddFunc("@daily", func() {
		fmt.Println(time.Now(), "Every day")
	})

	// Inspect the cron job entries' next and previous run times.
	for i, v := range c.Entries() {
		fmt.Println(i+1, "- prev:", v.Prev, "next:", v.Next)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// Stop the scheduler (does not stop any jobs already running).
	c.Stop()
}
