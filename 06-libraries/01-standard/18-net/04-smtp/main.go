package main

import (
	"fmt"
	"net/smtp"
)

const (
	address = "mail.example.com:25"
	host = "mail.example.com"
	port = 25
	username = "admin@example.com"
	password = "PASSWORD"
	sender = "sender@example.com"
	recipient = "receiver@example.com"
)

func main() {
	Sample("Smtp", Smtp)
	Sample("SendMail", SendMail)
}

func Smtp() {
	// Connect to the remote SMTP server.
	c, err := smtp.Dial(address)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Set the sender and recipient first
	if err := c.Mail(sender); err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := c.Rcpt(recipient); err != nil {
		fmt.Println("error:", err)
		return
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = wc.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func SendMail() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail(address, auth, sender, to, msg)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}