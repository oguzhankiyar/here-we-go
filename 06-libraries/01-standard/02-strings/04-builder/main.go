package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.Write([]byte{ ',', ' ' })
	builder.WriteRune(65)
	builder.WriteByte('!')

	fmt.Printf("Len: %d\n", builder.Len()) // 9
	fmt.Printf("Cap: %d\n", builder.Cap()) // 16
	fmt.Printf("Value: %q\n", builder.String()) // "Hello, A!"

	builder.Reset()

	fmt.Printf("Len: %d\n", builder.Len()) // 0
	fmt.Printf("Cap: %d\n", builder.Cap()) // 0
	fmt.Printf("Value: %q\n", builder.String()) // ""

	builder.WriteString("Hello, Gopher!")
	builder.WriteByte(' ')
	builder.WriteString("How are you?")

	fmt.Printf("Len: %d\n", builder.Len()) // 26
	fmt.Printf("Cap: %d\n", builder.Cap()) // 32
	fmt.Printf("Value: %q\n", builder.String()) // "Hello, Gopher! How are you?"

	builder.Grow(100)

	fmt.Printf("Len: %d\n", builder.Len()) // 26
	fmt.Printf("Cap: %d\n", builder.Cap()) // 164
	fmt.Printf("Value: %q\n", builder.String()) // "Hello, Gopher! How are you?"
}