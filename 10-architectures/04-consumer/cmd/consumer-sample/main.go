package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"consumer-sample/internal/application/consumers"
	"consumer-sample/internal/application/services"
	"consumer-sample/internal/infrastructure/broker/consumer"
	"consumer-sample/internal/infrastructure/config/parser"
	"consumer-sample/internal/infrastructure/logger"
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

	kafkaConsumer := consumer.NewKafkaConsumer(config.Broker, logger)
	healthService := services.NewHealthService(config.App, logger)
	healthConsumer := consumers.NewHealthConsumer(healthService, kafkaConsumer, logger)

	err = kafkaConsumer.Open()
	if err != nil {
		logger.Error("the consumer could not be opened", err)
	}

	healthConsumer.Start()

	go func() {
		for {
			healthService.Check()
			time.Sleep(time.Duration(config.App.HealthTreshold/2) * time.Second)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	err = kafkaConsumer.Close()
	if err != nil {
		logger.Error("the consumer could not be opened", err)
	}
}
