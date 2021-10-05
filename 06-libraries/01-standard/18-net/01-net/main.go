package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	Sample("IPv4", IPv4)
	Sample("LookupIP", LookupIP)
	Sample("ParseIP", ParseIP)
	Sample("Listener", Listener)
	Sample("Dialer", Dialer)
}

func IPv4() {
	fmt.Println(net.IPv4(8, 8, 8, 8))
}

func LookupIP() {
	ips, err := net.LookupIP("golang.org")
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func ParseIP() {
	fmt.Println(net.ParseIP("192.0.2.1"))    // 192.0.2.1
	fmt.Println(net.ParseIP("2001:db8::68")) // 2001:db8::68
	fmt.Println(net.ParseIP("192.0.2"))      // nil
}

func Listener() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

func Dialer() {
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", ":2000")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		fmt.Println("error:", err)
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
