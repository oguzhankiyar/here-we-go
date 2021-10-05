package main

import (
	"fmt"
	"os"
)

func main() {
	Sample("Print", Print)
	Sample("Println", Println)
	Sample("Printf", Printf)
	Sample("SPrint", SPrint)
	Sample("SPrintln", SPrintln)
	Sample("SPrintf", SPrintf)
	Sample("FPrint", Fprint)
	Sample("FPrintln", Fprintln)
	Sample("FPrintf", Fprintf)
}

func Print() {
	// Prints in same line

	fmt.Print("Hello") // Hello
	fmt.Print(" ")
	fmt.Print("Hello", "World") // Hello World
	fmt.Print(" ")
	fmt.Print(100) // 100
	fmt.Print(" ")
	fmt.Print(7.2) // 7.2
	fmt.Print(" ")
	fmt.Print(true) // true
	fmt.Print(" ")
	fmt.Print([]string{ "Gopher" }) // [Gopher]
	fmt.Print(" ")
	fmt.Print(map[string]string{ "name": "Gopher" }) // map[name:Gopher]
	fmt.Print(" ")
	fmt.Print(struct { }{}) // {}
	fmt.Print(" ")
	fmt.Print(Item{ Text: "Hey!" }) // {Hey!}
	fmt.Print("\n")
}

func Println() {
	// Prints in new line

	fmt.Println("Hello") // Hello
	fmt.Println("Hello", "World") // Hello World
	fmt.Println(100) // 100
	fmt.Println(7.2) // 7.2
	fmt.Println(true) // true
	fmt.Println([]string{ "Gopher" }) // [Gopher]
	fmt.Println(map[string]string{ "name": "Gopher" }) // map[name:Gopher]
	fmt.Println(struct { }{}) // {}
	fmt.Println(Item{ Text: "Hey!" }) // {Hey!}
}

func Printf() {
	// Prints with format in same line

	// %v -> default format
	fmt.Printf("%v\n", "Hello") // Hello
	fmt.Printf("%v %v\n", "Hello", "World") // Hello World
	fmt.Printf("%v\n", 100) // 5
	fmt.Printf("%v\n", 7.2) // 7.2
	fmt.Printf("%v\n", true) // true
	fmt.Printf("%v\n", []string{ "Gopher" }) // [Gopher]
	fmt.Printf("%v\n", map[string]string{ "name": "Gopher" }) // map[name:Gopher]
	fmt.Printf("%v\n", struct { }{}) // {}
	fmt.Printf("%v\n", Item{ Text: "Hey!" }) // {Hey!}

	// %#v -> go syntax
	fmt.Printf("%#v\n", "Hello") // "Hello"
	fmt.Printf("%#v %#v\n", "Hello", "World") // "Hello" "World"
	fmt.Printf("%#v\n", 100) // 100
	fmt.Printf("%#v\n", 7.2) // 7.2
	fmt.Printf("%#v\n", true) // true
	fmt.Printf("%#v\n", []string{ "Gopher" }) // []string{"Gopher"}
	fmt.Printf("%#v\n", map[string]string{ "name": "Gopher" }) // map[string]string{"name":"Gopher"}
	fmt.Printf("%#v\n", struct { }{}) // struct {}{}
	fmt.Printf("%#v\n", Item{ Text: "Hey!" }) // main.go.Item{Text:"Hey!"}

	// %T -> go syntax type
	fmt.Printf("%T\n", "Hello") // string
	fmt.Printf("%T %T\n", "Hello", "World") // string string
	fmt.Printf("%T\n", 100) // int
	fmt.Printf("%T\n", 7.2) // float64
	fmt.Printf("%T\n", true) // bool
	fmt.Printf("%T\n", []string{ "Gopher" }) // []string
	fmt.Printf("%T\n", map[string]string{ "name": "Gopher" }) // map[string]string
	fmt.Printf("%T\n", struct { }{}) // struct {}
	fmt.Printf("%T\n", Item{ Text: "Hey!" }) // main.go.Item

	// %% -> percent sign
	fmt.Printf("50%%\n") // 50%

	// %t -> boolean
	fmt.Printf("%t\n", true) // true
	fmt.Printf("%t\n", false) // false

	// %b -> base 2 for int, scientific notation for float
	fmt.Printf("%b\n", 100) // 1100100
	fmt.Printf("%b\n", 7.2) // 8106479329266893p-50

	// %c -> char
	fmt.Printf("%c\n", 'H') // H

	// %d -> base 10
	fmt.Printf("%d\n", 100) // 100

	// %+d -> always print sign
	fmt.Printf("%+d\n", 100) // +100

	// %e, %E -> scientific notation
	fmt.Printf("%e\n", 7.2) // 7.200000e+00
	fmt.Printf("%E\n", 7.2) // 7.200000E+00

	// %f, %F -> decimal
	fmt.Printf("%f\n", 7.2) // 7.200000
	fmt.Printf("%F\n", 7.2) // 7.200000
	fmt.Printf("%.2f\n", 7.2545487) // 7.25
	fmt.Printf("%6.2f\n", 7.254) //   7.25

	// %g, %G -> %e, %E for large otherwise %f, %F
	fmt.Printf("%g\n", 7.2) // 7.2
	fmt.Printf("%g\n", 7.25887878) // 7.25887878
	fmt.Printf("%G\n", 7.2) // 7.2
	fmt.Printf("%G\n", 7.25887878) // 7.25887878

	// %o, %O -> base 8
	fmt.Printf("%o\n", 100) // 144
	fmt.Printf("%O\n", 100) // 0o144

	// %q -> quoted char
	fmt.Printf("%q\n", "Hello") // "Hello"
	fmt.Printf("%q\n", 'H') // 'H'

	// %x, %X -> base 16
	fmt.Printf("%x\n", 100) // 64
	fmt.Printf("%x\n", "Hello") // 48656c6c6f
	fmt.Printf("%X\n", "Hello") // 48656C6C6F

	// %U -> unicode format
	fmt.Printf("%U\n", 100) // U+0064

	// %p -> pointer
	a := 5
	fmt.Printf("%p\n", &a) // 0xc000128098

	// %s -> string
	fmt.Printf("%s", "Hello") // Hello
}

func SPrint() {
	// Returns string same as Print func

	result := ""

	result = fmt.Sprint("Hello")
	fmt.Printf("%#v\n", result) // "Hello"

	result = fmt.Sprint("Hello", "World")
	fmt.Printf("%#v\n", result) // "HelloWorld"

	result = fmt.Sprint(100)
	fmt.Printf("%#v\n", result) // "100"

	result = fmt.Sprint(7.2)
	fmt.Printf("%#v\n", result) // "7.2"

	result = fmt.Sprint(true)
	fmt.Printf("%#v\n", result) // "true"
}

func SPrintln() {
	// Returns string same as Println func

	result := ""

	result = fmt.Sprintln("Hello")
	fmt.Printf("%#v\n", result) // "Hello\n"

	result = fmt.Sprintln("Hello", "World")
	fmt.Printf("%#v\n", result) // "Hello World\n"

	result = fmt.Sprintln(100)
	fmt.Printf("%#v\n", result) // "100\n"

	result = fmt.Sprintln(7.2)
	fmt.Printf("%#v\n", result) // "7.2\n"

	result = fmt.Sprintln(true)
	fmt.Printf("%#v\n", result) // "true\n"
}

func SPrintf() {
	// Returns string same as Printf func

	result := ""

	result = fmt.Sprintf("%s\n", "Hello")
	fmt.Printf("%#v\n", result) // "Hello\n"

	result = fmt.Sprintf("%s %s\n", "Hello", "World")
	fmt.Printf("%#v\n", result) // "Hello World\n"

	result = fmt.Sprintf("%d\n", 100)
	fmt.Printf("%#v\n", result) // "100\n"

	result = fmt.Sprintf("%f\n", 7.2)
	fmt.Printf("%#v\n", result) // "7.2\n"

	result = fmt.Sprintf("%t\n", true)
	fmt.Printf("%#v\n", result) // "true\n"
}

func Fprint() {
	// Writes string same as Print func

	fmt.Fprint(os.Stdout, "Hello") // Hello
	fmt.Println()
	fmt.Fprint(os.Stdout, "Hello", "World") // HelloWorld
	fmt.Println()
	fmt.Fprint(os.Stdout, 100) // 100
	fmt.Println()
	fmt.Fprint(os.Stdout, 7.2) // 7.2
	fmt.Println()
	fmt.Fprint(os.Stdout, true) // true
	fmt.Println()
}

func Fprintln() {
	// Writes string same as Println func

	fmt.Fprintln(os.Stdout, "Hello") // Hello
	fmt.Fprintln(os.Stdout, "Hello", "World") // Hello World
	fmt.Fprintln(os.Stdout, 100) // 100
	fmt.Fprintln(os.Stdout, 7.2) // 7.2
	fmt.Fprintln(os.Stdout, true) // true
}

func Fprintf() {
	// Writes string same as Printf func

	fmt.Fprintf(os.Stdout, "%s\n", "Hello") // Hello
	fmt.Fprintf(os.Stdout, "%s %s\n", "Hello", "World") // Hello World
	fmt.Fprintf(os.Stdout, "%d\n", 100) // 100
	fmt.Fprintf(os.Stdout, "%f\n", 7.2) // 7.2
	fmt.Fprintf(os.Stdout, "%t\n", true) // true
}

type Item struct {
	Text string
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}