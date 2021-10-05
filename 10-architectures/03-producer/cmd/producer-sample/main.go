package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"producer-sample/internal/application/producers"
	"producer-sample/internal/application/services"
	"producer-sample/internal/infrastructure/broker/producer"
	"producer-sample/internal/infrastructure/config/parser"
	"producer-sample/internal/infrastructure/logger"
)

func main() {
	configParser := parser.NewConfigParser("./configs", "dev", "json")
	config, err := configParser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(config.Logger)
	err = logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	kafkaProducer := producer.NewKafkaProducer(config.Broker, logger)
	healthService := services.NewHealthService(config.App)
	healthProducer := producers.NewHealthProducer(healthService, kafkaProducer, logger)

	err = kafkaProducer.Open()
	if err != nil {
		logger.Error("the producer could not be opened", err)
	}

	go func() {
		for {
			err := healthProducer.Send()
			if err != nil {
				logger.Error("the message could not be sent", err)
			}
			time.Sleep(time.Duration(config.App.HealthInterval) * time.Second)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	err = kafkaProducer.Close()
	if err != nil {
		logger.Error("the producer could not be opened", err)
	}
}
