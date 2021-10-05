package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/streadway/amqp"
)

type Message []byte

const ExchangeName = "pub-sub"

func main() {
	ctx := context.Background()

	url := "amqp://guest:guest@localhost:5672/"

	sessionsCh := redial(ctx, url)

	go func() {
		readCh := read()
		publish(sessionsCh, readCh)
	}()

	go func() {
		writeCh := write()
		subscribe(sessionsCh, writeCh)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

type session struct {
	*amqp.Connection
	*amqp.Channel
}

func (s session) Close() error {
	if s.Connection == nil {
		return nil
	}
	return s.Connection.Close()
}

func redial(ctx context.Context, url string) chan chan session {
	sessions := make(chan chan session)

	go func() {
		sess := make(chan session)
		defer close(sessions)

		for {
			select {
			case sessions <- sess:
			case <-ctx.Done():
				log.Println("shutting down session factory")
				return
			}

			conn, err := amqp.Dial(url)
			if err != nil {
				log.Fatalf("cannot (re)dial: %v: %q", err, url)
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			if err := ch.ExchangeDeclare(ExchangeName, "fanout", false, true, false, false, nil); err != nil {
				log.Fatalf("cannot declare fanout exchange: %v", err)
			}

			select {
			case sess <- session{conn, ch}:
			case <-ctx.Done():
				log.Println("shutting down new session")
				return
			}
		}
	}()

	return sessions
}

func publish(sessions chan chan session, messages <-chan Message) {
	for session := range sessions {
		var (
			running bool
			reading = messages
			pending = make(chan Message, 1)
			confirm = make(chan amqp.Confirmation, 1)
		)

		pub := <-session

		// publisher confirms for this channel/connection
		if err := pub.Confirm(false); err != nil {
			log.Printf("publisher confirms not supported")
			close(confirm) // confirms not supported, simulate by always nacking
		} else {
			pub.NotifyPublish(confirm)
		}

		log.Printf("publishing...")

	Publish:
		for {
			var body Message
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break Publish
				}
				if !confirmed.Ack {
					log.Printf("nack message %d, body: %q", confirmed.DeliveryTag, string(body))
				}
				reading = messages

			case body = <-pending:
				routingKey := "ignored for fanout exchanges, application dependent for other exchanges"
				err := pub.Publish(ExchangeName, routingKey, false, false, amqp.Publishing{
					Body: body,
				})

				// Retry failed delivery on the next session
				if err != nil {
					pending <- body
					pub.Close()
					break Publish
				}

			case body, running = <-reading:
				// all messages consumed
				if !running {
					return
				}

				// work on pending delivery until ack
				pending <- body
				reading = nil
			}
		}
	}
}

func subscribe(sessions chan chan session, messages chan<- Message) {
	queue := fmt.Sprintf("%v", time.Now().Unix())

	for session := range sessions {
		sub := <-session

		if _, err := sub.QueueDeclare(queue, false, true, true, false, nil); err != nil {
			log.Printf("cannot consume from exclusive queue: %q, %v", queue, err)
			return
		}

		routingKey := "application specific routing key for fancy toplogies"
		if err := sub.QueueBind(queue, routingKey, ExchangeName, false, nil); err != nil {
			log.Printf("cannot consume without a binding to exchange: %q, %v", ExchangeName, err)
			return
		}

		deliveries, err := sub.Consume(queue, "", false, true, false, false, nil)
		if err != nil {
			log.Printf("cannot consume from: %q, %v", queue, err)
			return
		}

		log.Printf("subscribed...")

		for msg := range deliveries {
			messages <- Message(msg.Body)
			sub.Ack(msg.DeliveryTag, false)
		}
	}
}

func read() <-chan Message {
	lines := make(chan Message)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			lines <- Message(scan.Bytes())
		}
	}()
	return lines
}

func write() chan<- Message {
	lines := make(chan Message)
	go func() {
		for line := range lines {
			fmt.Printf("received: %s\n", line)
		}
	}()
	return lines
}