package producer

import (
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	brokerInterfaces "producer-sample/internal/infrastructure/broker/interfaces"
	"producer-sample/internal/infrastructure/config/models"
	loggerInterfaces "producer-sample/internal/infrastructure/logger/interfaces"
)

type KafkaProducer struct {
	config   models.BrokerConfig
	producer *kafka.Producer
	logger   loggerInterfaces.Logger
}

func NewKafkaProducer(config models.BrokerConfig, logger loggerInterfaces.Logger) brokerInterfaces.Producer {
	return &KafkaProducer{
		config: config,
		logger: logger,
	}
}

func (p *KafkaProducer) Open() error {
	server := fmt.Sprintf("%s:%v", p.config.Host, p.config.Port)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})

	if err != nil {
		return err
	}

	go func() {
		for e := range producer.Events() {
			switch event := e.(type) {
			case *kafka.Message:
				if event.TopicPartition.Error != nil {
					p.logger.Debug("delivery failed to " + *event.TopicPartition.Topic)
				} else {
					p.logger.Debug("delivered message to " + *event.TopicPartition.Topic)
				}
			}
		}
	}()

	p.producer = producer

	return nil
}

func (p *KafkaProducer) Close() error {
	if p.producer == nil {
		return errors.New("the producer was not opened")
	}

	p.producer.Close()

	return nil
}

func (p *KafkaProducer) Produce(data []byte) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.config.TopicName, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}
