package consumer

import (
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	brokerInterfaces "consumer-sample/internal/infrastructure/broker/interfaces"
	brokerModels "consumer-sample/internal/infrastructure/broker/models"
	configModels "consumer-sample/internal/infrastructure/config/models"
	loggerInterfaces "consumer-sample/internal/infrastructure/logger/interfaces"
)

type KafkaConsumer struct {
	config        configModels.BrokerConfig
	consumer      *kafka.Consumer
	logger        loggerInterfaces.Logger
	subscriptions []brokerModels.Subscription
}

func NewKafkaConsumer(config configModels.BrokerConfig, logger loggerInterfaces.Logger) brokerInterfaces.Consumer {
	return &KafkaConsumer{
		config:        config,
		logger:        logger,
		subscriptions: make([]brokerModels.Subscription, 0),
	}
}

func (c *KafkaConsumer) Open() error {
	server := fmt.Sprintf("%s:%v", c.config.Host, c.config.Port)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          c.config.GroupId,
	})

	if err != nil {
		return err
	}

	consumer.SubscribeTopics([]string{c.config.TopicName}, nil)

	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)

			if err != nil {
				c.logger.Error("consumer error", err)
				continue
			}

			for _, subscription := range c.subscriptions {
				go subscription.Fn(msg.Value)
			}
		}
	}()

	c.consumer = consumer

	return nil
}

func (c *KafkaConsumer) Close() error {
	if c.consumer == nil {
		return errors.New("the consumer was not opened")
	}

	c.consumer.Close()

	return nil
}

func (c *KafkaConsumer) Subscribe(subscription brokerModels.Subscription) {
	c.subscriptions = append(c.subscriptions, subscription)
}

func (c *KafkaConsumer) Unsubscribe(subscription brokerModels.Subscription) {
	for i, v := range c.subscriptions {
		if v.Id == subscription.Id {
			c.subscriptions = append(c.subscriptions[:i], c.subscriptions[i+1:]...)
			break
		}
	}
}
