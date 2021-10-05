package main

import (
	"fmt"
	"os"
)

func main() {
	Sample("Scan", Scan)
	Sample("Scanln", Scanln)
	Sample("Scanf", Scanf)
	Sample("Sscan", Sscan)
	Sample("Sscanln", Sscanln)
	Sample("Sscanf", Sscanf)
	Sample("Fscan", Fscan)
	Sample("Fscanln", Fscanln)
	Sample("Fscanf", Fscanf)
}

func Scan() {
	// Reads from standard input with space-separated

	var name string
	var age int

	fmt.Scan(&name)
	fmt.Scan(&age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Scanln() {
	// Reads from standard input with lines

	var name string
	var age int

	fmt.Scanln(&name)
	fmt.Scanln(&age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Scanf() {
	// Reads from standard input with format

	var name string
	var age int

	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("name: %s, age: %d\n", name, age)
}

func Sscan() {
	// Reads from string same as Scan func

	var i1, i2 int
	fmt.Sscan("100 200", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sscanln() {
	// Reads from string same as Scanln func

	var i1, i2 int
	fmt.Sscanln("100\n200", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sscanf() {
	// Reads from string same as Scanf func

	var i1, i2 int
	fmt.Sscanf("100 and 200", "%d and %d", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscan() {
	// Reads from reader same as Scan func

	var i1, i2 int
	fmt.Fscan(os.Stdin, &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscanln() {
	// Reads from reader same as Scanln func

	var i1, i2 int
	fmt.Fscanln(os.Stdin, &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Fscanf() {
	// Reads from reader same as Scanf func

	var i1, i2 int
	fmt.Fscanf(os.Stdin, "%d and %d", &i1, &i2)
	fmt.Printf("int1: %d, int2: %d\n", i1, i2)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}