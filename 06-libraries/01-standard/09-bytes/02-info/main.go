package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	Sample("Count", Count)
	Sample("Fields", Fields)
	Sample("FieldsFunc", FieldsFunc)
	Sample("Index", Index)
	Sample("IndexAny", IndexAny)
	Sample("IndexByte", IndexByte)
	Sample("IndexFunc", IndexFunc)
	Sample("IndexRune", IndexRune)
	Sample("LastIndex", LastIndex)
	Sample("LastIndexAny", LastIndexAny)
	Sample("LastIndexByte", LastIndexByte)
	Sample("LastIndexFunc", LastIndexFunc)
}

func Count() {
	// Searches and counts substr in str case insensitive

	fmt.Printf("%q - %q => %d\n", "This is text", "i", bytes.Count([]byte("This is text"), []byte("i")))
}

func Fields() {
	// Splits words with space

	fmt.Printf("%q => %q\n", "This is text", bytes.Fields([]byte("This is text")))
	fmt.Printf("%q => %q\n", "  foo1;bar2,baz3...", bytes.Fields([]byte("  foo1;bar2,baz3...")))
}

func FieldsFunc() {
	// Splits words with func

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("%q => %q\n", "This is text", bytes.FieldsFunc([]byte("This is text"), f))
	fmt.Printf("%q => %q\n", "  foo1;bar2,baz3...", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
}

func Index() {
	// Finds index of substr in str
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "chicken", "ken", bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Printf("%q - %q => %d\n", "chicken", "dmr", bytes.Index([]byte("chicken"), []byte("dmr")))
}

func IndexAny() {
	// Finds index of chars of substr
	// If exist any, returns first char index
	// Else, returns -1

	fmt.Printf("%q - %q => %d\n", "chicken", "aeiouy", bytes.IndexAny([]byte("chicken"), "aeiouy"))
	fmt.Printf("%q - %q => %d\n", "crwth", "aeiouy", bytes.IndexAny([]byte("crwth"), "aeiouy"))
}

func IndexByte() {
	// Finds first index of char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "golang", 'g', bytes.IndexByte([]byte("golang"), 'g'))
	fmt.Printf("%q - %q => %d\n", "gophers", 'h', bytes.IndexByte([]byte("gophers"), 'h'))
	fmt.Printf("%q - %q => %d\n", "golang", 'x', bytes.IndexByte([]byte("golang"), 'x'))
}

func IndexFunc() {
	// Finds first index of matched with func
	// If not exist, returns -1

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	fmt.Printf("%q => %d\n", "Hello, 世界", bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Printf("%q => %d\n", "Hello, world", bytes.IndexFunc([]byte("Hello, world"), f))
}

func IndexRune() {
	// Finds first index of rune char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "golang", rune(103), bytes.IndexRune([]byte("golang"), rune(103)))
	fmt.Printf("%q - %q => %d\n", "gophers", rune(104), bytes.IndexRune([]byte("gophers"), rune(104)))
	fmt.Printf("%q - %q => %d\n", "golang", rune(120), bytes.IndexRune([]byte("golang"), rune(120)))
}

func LastIndex() {
	// Finds last index of substr in str
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "go gopher", "go", bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Printf("%q - %q => %d\n", "go gopher", "rodent", bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}

func LastIndexAny() {
	// Finds index of chars of substr
	// If exist any, returns last char index
	// Else, returns -1

	fmt.Printf("%q - %q => %d\n", "go gopher", "go", bytes.LastIndexAny([]byte("go gopher"), "go"))
	fmt.Printf("%q - %q => %d\n", "go gopher", "rodent", bytes.LastIndexAny([]byte("go gopher"), "rodent"))
	fmt.Printf("%q - %q => %d\n", "go gopher", "fail", bytes.LastIndexAny([]byte("go gopher"), "fail"))
}

func LastIndexByte() {
	// Finds first index of char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "Hello, world", 'l', bytes.LastIndexByte([]byte("Hello, world"), 'l'))
	fmt.Printf("%q - %q => %d\n", "Hello, world", 'o', bytes.LastIndexByte([]byte("Hello, world"), 'o'))
	fmt.Printf("%q - %q => %d\n", "Hello, world", 'x', bytes.LastIndexByte([]byte("Hello, world"), 'x'))
}

func LastIndexFunc() {
	// Finds last index of matched with func
	// If not exist, returns -1

	f := unicode.IsNumber

	fmt.Printf("%q => %d\n", "go 123", bytes.LastIndexFunc([]byte("go 123"), f))
	fmt.Printf("%q => %d\n", "123 go", bytes.LastIndexFunc([]byte("123 go"), f))
	fmt.Printf("%q => %d\n", "go", bytes.LastIndexFunc([]byte("go"), f))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}