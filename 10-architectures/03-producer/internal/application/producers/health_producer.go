package producers

import (
	"encoding/json"

	"producer-sample/internal/application/services"
	producerInterfaces "producer-sample/internal/infrastructure/broker/interfaces"
	loggerInterfaces "producer-sample/internal/infrastructure/logger/interfaces"
)

type HealthProducer struct {
	service  *services.HealthService
	producer producerInterfaces.Producer
	logger   loggerInterfaces.Logger
}

func NewHealthProducer(service *services.HealthService, producer producerInterfaces.Producer, logger loggerInterfaces.Logger) *HealthProducer {
	return &HealthProducer{
		service:  service,
		producer: producer,
		logger:   logger,
	}
}

func (p *HealthProducer) Send() error {
	info := p.service.Get()

	data, err := json.Marshal(info)
	if err != nil {
		p.logger.Error("the message could not be marshaled", err)
	}

	return p.producer.Produce(data)
}
