package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func main() {
	// Do job without params
	gocron.Every(1).Second().Do(Task)
	gocron.Every(2).Seconds().Do(Task)
	gocron.Every(1).Minute().Do(Task)
	gocron.Every(2).Minutes().Do(Task)
	gocron.Every(1).Hour().Do(Task)
	gocron.Every(2).Hours().Do(Task)
	gocron.Every(1).Day().Do(Task)
	gocron.Every(2).Days().Do(Task)
	gocron.Every(1).Week().Do(Task)
	gocron.Every(2).Weeks().Do(Task)

	// Do job with params
	gocron.Every(1).Second().Do(TaskWithParams, 1, "hello")

	// Do job on specific weekday
	gocron.Every(1).Monday().Do(Task)
	gocron.Every(1).Thursday().Do(Task)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	gocron.Every(1).Day().At("10:30").Do(Task)
	gocron.Every(1).Monday().At("18:30").Do(Task)
	gocron.Every(1).Tuesday().At("18:30:59").Do(Task)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(Task)

	// Begin job at a specific date/time
	t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	gocron.Every(1).Hour().From(&t).Do(Task)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	gocron.Remove(Task)

	// Clear all scheduled jobs
	gocron.Clear()

	// Start all the pending jobs
	<- gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(Task)
	<- s.Start()
}

func Task() {
	fmt.Println("I am running task.")
}

func TaskWithParams(a int, b string) {
	fmt.Println(a, b)
}