package main

import (
	"fmt"
)

type Observer interface {
	OnReceived(data interface{})
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) Subscribe(o Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) Unsubscribe(o Observer) {
	var index int
	for i, v := range p.ObserverList {
		if v == o {
			index = i
			break
		}
	}

	p.ObserverList = append(p.ObserverList[:index], p.ObserverList[index + 1:]...)
}

func (p *Publisher) Publish(data interface{}) {
	fmt.Println("Publishing...")
	for _, o := range p.ObserverList {
		o.OnReceived(data)
	}
	fmt.Println("Published!")
}

type MessageObserver struct {

}

func (o MessageObserver) OnReceived(data interface{}) {
	fmt.Println("Received:", data)
}

func main() {
	publisher := Publisher{}

	publisher.Subscribe(MessageObserver{})

	publisher.Publish("Hello!")
}