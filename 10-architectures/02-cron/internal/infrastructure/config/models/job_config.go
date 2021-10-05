package models

import "cron-sample/internal/infrastructure/config/errors"

type JobConfig struct {
	Id    string                 `mapstructure:"id"`
	Cron  string                 `mapstructure:"cron"`
	Retry uint                   `mapstructure:"retry"`
	Args  map[string]interface{} `mapstructure:"args"`
}

func (c *JobConfig) Validate() error {
	if len(c.Id) == 0 {
		return errors.NewConfigValidateError("invalid job id")
	}

	if len(c.Cron) == 0 {
		return errors.NewConfigValidateError("invalid job id")
	}

	return nil
}

type JobsConfig []JobConfig

func NewJobsConfig() *JobsConfig {
	config := JobsConfig(make([]JobConfig, 0))
	return &config
}

func (c *JobsConfig) Validate() error {
	for _, v := range *c {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
