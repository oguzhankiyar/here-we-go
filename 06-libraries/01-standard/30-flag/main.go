package main

import (
	"flag"
	"fmt"
)

func main() {
	// Run from terminal
	// go run main.go -firstname=oguzhan -lastname=kiyar -active=false verbose

	Sample("Simple", Simple)
	Sample("PrintDefaults", PrintDefaults)
	Sample("Args", Args)
}

func Simple() {
	firstname := flag.String("firstname", "", "the firstname")
	age := flag.Int("age", 0, "the age")
	active := flag.Bool("active", true, "the activation")

	var lastname string
	flag.StringVar(&lastname, "lastname", "", "the lastname")

	flag.Parse()

	fmt.Println("name:", *firstname)
	fmt.Println("lastname:", lastname)
	fmt.Println("age:", *age)
	fmt.Println("active:", *active)
}

func PrintDefaults() {
	flag.PrintDefaults()
}

func Args() {
	fmt.Println(flag.Args())
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}