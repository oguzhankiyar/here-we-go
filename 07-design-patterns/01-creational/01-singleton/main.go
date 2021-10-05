package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Connection struct {
	id string
}

var (
	conn 	Connection
	once 	sync.Once
)

func Connect() Connection {
	once.Do(func() {
		id := strconv.Itoa(int(time.Now().Unix() % 55776))
		conn = Connection{id}
	})

	return conn
}

func main() {
	fn := func() {
		conn := Connect()

		fmt.Println("conn_id:", conn.id)
	}

	fn()
	fn()
	fn()
}