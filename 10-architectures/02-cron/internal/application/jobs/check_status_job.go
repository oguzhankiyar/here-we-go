package jobs

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"cron-sample/internal/application/jobs/models"
	"cron-sample/internal/infrastructure/job"
	"cron-sample/internal/infrastructure/logger/interfaces"
)

type CheckStatusJob struct {
	logger interfaces.Logger
}

func NewCheckStatusJob(logger interfaces.Logger) job.Job {
	return &CheckStatusJob{
		logger: logger,
	}
}

func (j *CheckStatusJob) Id() string {
	return "check_status"
}

func (j *CheckStatusJob) Run(args map[string]interface{}) error {
	var data models.CheckStatusArgs
	err := mapstructure.Decode(args, &data)
	if err != nil {
		return err
	}

	j.logger.Debug(fmt.Sprintf("checking %s", data.Url))

	resp, err := http.Get(data.Url)
	if err != nil {
		j.logger.Debug(fmt.Sprintf("failed %s", err.Error()))
	} else {
		j.logger.Debug(fmt.Sprintf("succeeded %v", resp.StatusCode))
	}

	j.logger.Debug(fmt.Sprintf("checked %s", data.Url))

	return nil
}
