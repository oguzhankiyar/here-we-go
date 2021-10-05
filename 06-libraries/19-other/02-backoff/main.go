package main

import (
	"errors"
	"log"

	"github.com/cenkalti/backoff/v4"
)

func main() {
	var count int
	operation := func() error {
		count++
		log.Printf("#%v running\n", count)

		if count == 3 {
			log.Printf("#%v succeeded\n", count)
			return nil
		}

		return errors.New("failed")
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("success")
}