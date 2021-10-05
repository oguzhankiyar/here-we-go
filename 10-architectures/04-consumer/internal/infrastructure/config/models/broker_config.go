package models

import "consumer-sample/internal/infrastructure/config/errors"

type BrokerConfig struct {
	Host      string `mapstructure:"BROKER_HOST" env:"BROKER_HOST"`
	Port      int    `mapstructure:"BROKER_PORT" env:"BROKER_PORT"`
	GroupId   string `mapstructure:"groupId"`
	TopicName string `mapstructure:"topicName"`
}

func NewBrokerConfig() *BrokerConfig {
	return &BrokerConfig{}
}

func (c *BrokerConfig) Validate() error {
	if len(c.Host) == 0 {
		return errors.NewConfigValidateError("invalid broker host")
	}

	if c.Port == 0 {
		return errors.NewConfigValidateError("invalid broker port")
	}

	if len(c.GroupId) == 0 {
		return errors.NewConfigValidateError("invalid broker group id")
	}

	if len(c.TopicName) == 0 {
		return errors.NewConfigValidateError("invalid broker topic name")
	}

	return nil
}
