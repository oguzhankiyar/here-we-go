package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	go func() {
		Consume(ctx, client)
	}()

	go func() {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			Produce(ctx, client, scan.Bytes())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func Produce(ctx context.Context, client *redis.Client, msg []byte) {
	log.Println("sending")

	err := client.Publish(ctx, "test-channel", msg).Err()

	if err != nil {
		log.Println(err)
	}

	log.Println("sent")
}

func Consume(ctx context.Context, client *redis.Client) {
	ps := client.Subscribe(ctx, "test-channel")

	ch := ps.Channel()

	for msg := range ch {
		log.Printf("received: %s\n", msg.Payload)
	}
}
