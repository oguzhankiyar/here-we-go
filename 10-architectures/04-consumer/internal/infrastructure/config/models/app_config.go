package models

import "consumer-sample/internal/infrastructure/config/errors"

type AppConfig struct {
	Name           string `mapstructure:"appName"`
	Environment    string `mapstructure:"APP_ENV" env:"APP_ENV"`
	Version        string `mapstructure:"APP_VERSION" env:"APP_VERSION"`
	HealthTreshold int64  `mapstructure:"healthTreshold"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Environment: "dev",
		Version:     "1.0.0",
	}
}

func (c *AppConfig) Validate() error {
	if len(c.Name) == 0 {
		return errors.NewConfigValidateError("invalid app name")
	}

	if len(c.Environment) == 0 {
		return errors.NewConfigValidateError("invalid environment")
	}

	if len(c.Version) == 0 {
		return errors.NewConfigValidateError("invalid app version")
	}

	if c.HealthTreshold == 0 {
		return errors.NewConfigValidateError("invalid health treshold")
	}

	return nil
}
