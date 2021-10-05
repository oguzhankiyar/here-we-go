package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func RunClient() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})

	RunImmediately(client)
	RunFuture(client)
	DefaultOptions(client)
	OverrideOptions(client)

	client.Close()
}

func RunImmediately(client *asynq.Client) {
	// Enqueue task to be processed immediately.

	task, err := NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func RunFuture(client *asynq.Client) {
	// Schedule task to be processed in the future.

	task, err := NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func DefaultOptions(client *asynq.Client) {
	// Set other options to tune task processing behavior.

	client.SetDefaultOptions(TypeImageResize, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))

	task, err := NewImageResizeTask("https://example.com/myassets/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func OverrideOptions(client *asynq.Client) {
	// Pass options to tune task processing behavior at enqueue time.

	task, err := NewImageResizeTask("https://example.com/myassets/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task, asynq.Queue("critical"), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}