package main

import (
	"fmt"
	"strconv"
	"time"
)

type Connection struct {
	id		string
	url 	string
}

func NewConnection(host string, port int) Connection {
	id := strconv.Itoa(int(time.Now().Unix() % 55776))
	url := fmt.Sprintf("%s:%v", host, port)

	return Connection{id, url}
}

func main() {
	conn := NewConnection("127.0.0.1", 2805)

	fmt.Println("conn_id:", conn.id)
	fmt.Println("conn_url:", conn.url)
}