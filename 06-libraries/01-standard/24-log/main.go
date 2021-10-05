package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	Sample("Default", Default)
	Sample("New", New)
	Sample("Print", Print)
	Sample("Prefix", Prefix)
	Sample("Output", Output)
	Sample("Flags", Flags)
	Sample("Fatal", Fatal)
	Sample("Panic", Panic)
}

func Default() {
	logger := log.Default()

	_, err := logger.Writer().Write([]byte("Hello, there!"))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println()
}

func New() {
	logger := log.New(os.Stdout, "[GO-LOG] ", log.Ldate | log.Lmicroseconds | log.LUTC)

	logger.Println("Hello, there!")
}

func Print() {
	log.Print("Hello, print!")

	log.Printf("%s", "Hello, printf!")

	log.Println("Hello, println!")
}

func Prefix() {
	fmt.Printf("Current prefix: %q\n", log.Prefix())
	log.Println("this is test log")

	log.SetPrefix("[GO-LOG] ")

	fmt.Printf("Current prefix: %q\n", log.Prefix())
	log.Println("this is test log")

	log.SetPrefix("")
}

func Output() {
	log.SetOutput(os.Stdout)

	err := log.Output(1, "Say, hi!")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func Flags() {
	check := func() {
		fmt.Println(log.Flags() & log.Ldate > 0)
		fmt.Println(log.Flags() & log.Ltime > 0)
		fmt.Println(log.Flags() & log.Lmicroseconds > 0)
		fmt.Println(log.Flags() & log.Lshortfile > 0)
		fmt.Println(log.Flags() & log.LUTC > 0)

		fmt.Println()
	}

	check()

	log.SetFlags(log.Flags() | log.LUTC)

	check()
}

func Fatal() {
	go func() {
		log.Fatal("Hello, fatal!")
	}()

	go func() {
		log.Fatalf("%s", "Hello, fatalf!")
	}()

	go func() {
		log.Fatalln("Hello, fatalln!")
	}()
}

func Panic() {
	go func() {
		log.Panic("Hello, panic!")
	}()

	go func() {
		log.Panicf("%s", "Hello, panicf!")
	}()

	go func() {
		log.Panicln("Hello, panicln!")
	}()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
