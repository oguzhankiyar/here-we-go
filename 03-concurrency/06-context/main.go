package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Sample("WithTimeout", WithTimeout)
	Sample("WithCancel", WithCancel)
}

func WithTimeout() {
	// This context will be canceled after 100ms timeout
	ctx, cancelFn := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	defer cancelFn()

	select {
	case <-time.After(1 * time.Minute):
		break
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func WithCancel() {
	ctx, cancelFn := context.WithCancel(context.Background())

	go func() {
		time.Sleep(100 * time.Millisecond)

		// This function will cancel the context after 100ms sleep
		cancelFn()
	}()

	select {
	case <-time.After(1 * time.Minute):
		break
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}