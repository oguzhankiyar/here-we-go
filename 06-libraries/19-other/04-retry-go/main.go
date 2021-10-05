package main

import (
	"errors"
	"log"

	"github.com/avast/retry-go"
)

func main() {
	var count int

	err := retry.Do(
		func() error {
			count++
			log.Printf("#%v running\n", count)

			if count == 3 {
				log.Printf("#%v succeeded\n", count)
				return nil
			}

			return errors.New("failed")
		},
		retry.OnRetry(func(n uint, err error) {
			log.Printf("#%d %s\n", n + 1, err)
		}),
		retry.Attempts(3),
	)
	if err != nil {
		log.Fatal(err)
	}
}