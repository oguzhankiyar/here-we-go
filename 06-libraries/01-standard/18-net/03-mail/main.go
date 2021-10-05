package main

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
)

func main() {
	Sample("ParseAddress", ParseAddress)
	Sample("ReadMessage", ReadMessage)
}

func ParseAddress() {
	fn := func(str string) {
		address, err := mail.ParseAddress(str)
		if err != nil {
			fmt.Printf("%s -> Error: %s\n", str, err)
			return
		}

		fmt.Printf("%s -> Name: %s, Address: %s\n", str, address.Name, address.Address)
	}

	fn("Alice <alice@example.com>")
	fn("bob@example.com")
	fn("test.com")
	fn("@test.com")
	fn("admin@test")
}

func ReadMessage() {
	msg := `Date: Mon, 1 May 2020 15:32:23 -0400
From: Alice <alice@example.com>
To: Bob <bob@example.com>
Subject: Hello, Gopher!

Hello, Gopher! This mail is for test purpose.
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:     ", header.Get("Date"))
	fmt.Println("From:     ", header.Get("From"))
	fmt.Println("To:       ", header.Get("To"))
	fmt.Println("Subject:  ", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Body:      %s", body)
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}