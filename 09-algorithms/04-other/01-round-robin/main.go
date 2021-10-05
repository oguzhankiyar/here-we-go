package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	roundRobin := New([]string{
		"192.168.100.10",
		"192.168.100.11",
		"192.168.100.12",
		"192.168.100.13",
	})

	for i := 0; i < 5; i++ {
		next := roundRobin.Next()
		fmt.Println(next)
	}
}

type RoundRobin struct {
	Servers	[]string
	Count	uint32
}

func New(servers []string) *RoundRobin {
	return &RoundRobin{
		Servers: servers,
	}
}

func (r *RoundRobin) Next() string {
	count := atomic.AddUint32(&r.Count, 1)
	return r.Servers[(int(count) - 1) % len(r.Servers)]
}