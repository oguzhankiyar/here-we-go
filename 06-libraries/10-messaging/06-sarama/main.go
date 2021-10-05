package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	ctx, done := context.WithCancel(context.Background())

	go Consume(ctx)
	go Produce(ctx)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	done()
}

func Produce(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.MinVersion
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	var (
		wg sync.WaitGroup
		enqueued int
		successes int
		producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{
			Topic: "my_topic",
			Value: sarama.StringEncoder("testing 123"),
		}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-ctx.Done():
			producer.AsyncClose()
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}

func Consume(ctx context.Context) {
	config := sarama.NewConfig()
	config.Version = sarama.MinVersion

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	consumed := 0

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-ctx.Done():
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
