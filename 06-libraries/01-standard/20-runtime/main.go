package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

func main() {
	Sample("Trace", Trace)
}

func Trace() {
	f, err := os.Create("06-libraries/01-standard/20-runtime/trace.txt")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	func() {
		fmt.Println("this function will be traced")
	}()

	content, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Content: %s\n", content)

	_ = os.Remove(f.Name())
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}