package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/avast/retry-go"
	"github.com/robfig/cron"

	"cron-sample/internal/application/jobs"
	"cron-sample/internal/infrastructure/config/parser"
	"cron-sample/internal/infrastructure/job"
	"cron-sample/internal/infrastructure/logger"
)

func main() {
	configParser := parser.NewConfigParser("./configs", "dev", "json")
	config, err := configParser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(config.Logger)
	err = logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	checkStatusJob := jobs.NewCheckStatusJob(logger)
	jobMap := map[string]job.Job{
		checkStatusJob.Id(): checkStatusJob,
	}

	c := cron.New()

	for _, job := range config.Jobs {
		if _, ok := jobMap[job.Id]; !ok {
			continue
		}

		err := c.AddFunc(job.Cron, func() {
			logger.Debug(fmt.Sprintf("started %s", job.Id))
			err := retry.Do(
				func() error {
					return jobMap[job.Id].Run(job.Args)
				},
				retry.OnRetry(func(n uint, err error) {
					logger.Debug(fmt.Sprintf("#%d %s\n", n+1, err))
				}),
				retry.Attempts(job.Retry),
			)
			if err != nil {
				logger.Fatal("failed", err)
			}
			logger.Debug(fmt.Sprintf("finished %s", job.Id))
		})
		if err != nil {
			logger.Fatal("failed", err)
		}
	}

	c.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	c.Stop()
}
