package models

import "consumer-sample/internal/infrastructure/config/errors"

type LoggerConfig struct {
	LogName  string `mapstructure:"logName"`
	LogLevel string `mapstructure:"logLevel"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		LogLevel: "debug",
		DevMode:  true,
		Encoder:  "console",
	}
}

func (c *LoggerConfig) Validate() error {
	if len(c.LogName) == 0 {
		return errors.NewConfigValidateError("invalid log name")
	}

	if len(c.LogLevel) == 0 {
		return errors.NewConfigValidateError("invalid log level")
	}

	if c.Encoder != "json" && c.Encoder != "console" {
		return errors.NewConfigValidateError("invalid log encoder")
	}

	return nil
}
