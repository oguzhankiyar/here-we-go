package consumers

import (
	"encoding/json"
	"sync"

	"consumer-sample/internal/application/services"
	healthModels "consumer-sample/internal/core/models"
	consumerInterfaces "consumer-sample/internal/infrastructure/broker/interfaces"
	brokerModels "consumer-sample/internal/infrastructure/broker/models"
	loggerInterfaces "consumer-sample/internal/infrastructure/logger/interfaces"
)

type HealthConsumer struct {
	service      *services.HealthService
	consumer     consumerInterfaces.Consumer
	logger       loggerInterfaces.Logger
	subscription brokerModels.Subscription
	mu           sync.Mutex
}

func NewHealthConsumer(service *services.HealthService, consumer consumerInterfaces.Consumer, logger loggerInterfaces.Logger) *HealthConsumer {
	healthConsumer := HealthConsumer{
		service:  service,
		consumer: consumer,
		logger:   logger,
	}

	healthConsumer.subscription = brokerModels.Subscription{
		Id: "health-consumer",
		Fn: healthConsumer.Consume,
	}

	return &healthConsumer
}

func (c *HealthConsumer) Start() {
	c.consumer.Subscribe(c.subscription)
}

func (c *HealthConsumer) Stop() {
	c.consumer.Unsubscribe(c.subscription)
}

func (c *HealthConsumer) Consume(data []byte) {
	var model healthModels.HealthModel

	err := json.Unmarshal(data, &model)
	if err != nil {
		c.logger.Error("the message could not be unmarshaled", err)
	}

	c.mu.Lock()
	c.service.Set(&model)
	c.mu.Unlock()
}
