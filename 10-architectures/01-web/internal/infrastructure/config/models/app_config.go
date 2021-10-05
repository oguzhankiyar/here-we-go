package models

import "web-sample/internal/infrastructure/config/errors"

type AppConfig struct {
	Name          string `mapstructure:"appName"`
	Environment   string `mapstructure:"APP_ENV" env:"APP_ENV"`
	Version       string `mapstructure:"APP_VERSION" env:"APP_VERSION"`
	Host          string `mapstructure:"APP_HOST" env:"APP_HOST"`
	Port          int    `mapstructure:"APP_PORT" env:"APP_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET" env:"JWT_SECRET"`
	AdminUsername string `mapstructure:"ADMIN_USERNAME" env:"ADMIN_USERNAME"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD" env:"ADMIN_PASSWORD"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Environment: "dev",
		Version:     "1.0.0",
		Host:        "localhost",
		Port:        8080,
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

	if len(c.Host) == 0 {
		return errors.NewConfigValidateError("invalid app host")
	}

	if c.Port == 0 {
		return errors.NewConfigValidateError("invalid app port")
	}

	if len(c.JwtSecret) == 0 {
		return errors.NewConfigValidateError("invalid jwt secret")
	}

	if len(c.AdminUsername) == 0 {
		return errors.NewConfigValidateError("invalid admin username")
	}

	if len(c.AdminPassword) == 0 {
		return errors.NewConfigValidateError("invalid admin password")
	}

	return nil
}
