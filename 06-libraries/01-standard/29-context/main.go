package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Sample("Background", Background)
	Sample("WithCancel", WithCancel)
	Sample("WithDeadline", WithDeadline)
	Sample("WithTimeout", WithTimeout)
	Sample("WithValue", WithValue)
}

func Background() {
	ctx := context.Background()

	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Err())
}

func WithCancel() {
	fn := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case ch <- n:
					n++
				}
			}
		}()
		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range fn(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func WithDeadline() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("done:", ctx.Err())
	}
}

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("done:", ctx.Err())
	}
}

func WithValue() {
	fn := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}

		fmt.Println("not found:", k)
	}

	ctx := context.WithValue(context.Background(), "language", "Go")

	fn(ctx, "language")
	fn(ctx, "color")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}