package models

type Config struct {
	App    AppConfig    `mapstructure:",squash"`
	Logger LoggerConfig `mapstructure:",squash"`
	Jobs   JobsConfig   `mapstructure:"jobs"`
}

func NewConfig() *Config {
	return &Config{
		App:    *NewAppConfig(),
		Logger: *NewLoggerConfig(),
		Jobs:   *NewJobsConfig(),
	}
}

func (c *Config) Validate() error {
	if err := c.App.Validate(); err != nil {
		return err
	}

	if err := c.Logger.Validate(); err != nil {
		return err
	}

	if err := c.Jobs.Validate(); err != nil {
		return err
	}

	return nil
}
