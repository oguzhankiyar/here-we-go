package configs

import "errors"

type ProductClientConfig struct {
	BaseUrl    string
	Username   string
	Password   string
	Timeout    int
	RetryCount int
}

func NewProductClientConfig() *ProductClientConfig {
	return &ProductClientConfig{
		Timeout:    10000,
		RetryCount: 3,
	}
}

func (config *ProductClientConfig) Validate() error {
	if len(config.BaseUrl) == 0 {
		return errors.New("invalid base url")
	}

	if len(config.Username) == 0 {
		return errors.New("invalid username")
	}

	if len(config.Password) == 0 {
		return errors.New("invalid password")
	}

	if config.Timeout < 0 {
		return errors.New("invalid timeout")
	}

	if config.RetryCount < 0 {
		return errors.New("invalid retry count")
	}

	return nil
}
