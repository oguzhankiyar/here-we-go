package main

import (
	"fmt"
	"strconv"
)

type CustomError struct {
	Code	int
	Message	string
}

func (e CustomError) Error() string {
	return "[" + strconv.Itoa(e.Code) + "] " + e.Message
}

func Fail(code int, message string) CustomError {
	return CustomError{
		Code: code,
		Message: message,
	}
}

func Pay(amount float64) error {
	balance := 1000.0

	if amount > balance {
		return Fail(315, "Insufficient Balance")
	}

	return nil
}

func main() {
	// Simple
	err := Pay(5000)
	if err != nil {
		fmt.Println(err)
	}

	// Assertion
	if c, ok := err.(CustomError); ok {
		fmt.Println(c.Code)
		fmt.Println(c.Message)
	}
}