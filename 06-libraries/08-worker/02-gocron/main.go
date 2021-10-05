package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Seconds().Do(func(){
		fmt.Println("tick")
	})

	// strings parse to duration
	s.Every("5m").Do(func(){
		fmt.Println("tick")
	})

	s.Every(5).Days().Do(func(){
		fmt.Println("tick")
	})

	task := func() {
		fmt.Println("tick")
	}

	// cron expressions supported
	s.Cron("*/1 * * * *").Do(task) // every minute

	s.Every(2).Day().Tag("tag").At("10:00").Do(task)
	s.Every(1).Minute().Tag("tag").Do(task)
	s.RunByTag("tag")

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	s.StartAsync()

	// starts the scheduler and blocks current execution path
	s.StartBlocking()
}