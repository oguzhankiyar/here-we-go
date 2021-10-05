package main

import (
	"fmt"
	"strings"
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

	fmt.Printf("%q - %q => %d\n", "This is text", "i", strings.Count("This is text", "i"))
}

func Fields() {
	// Splits words with space

	fmt.Printf("%q => %q\n", "This is text", strings.Fields("This is text"))
	fmt.Printf("%q => %q\n", "  foo1;bar2,baz3...", strings.Fields("  foo1;bar2,baz3..."))
}

func FieldsFunc() {
	// Splits words with func

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("%q => %q\n", "This is text", strings.FieldsFunc("This is text", f))
	fmt.Printf("%q => %q\n", "  foo1;bar2,baz3...", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}

func Index() {
	// Finds index of substr in str
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "chicken", "ken", strings.Index("chicken", "ken"))
	fmt.Printf("%q - %q => %d\n", "chicken", "dmr", strings.Index("chicken", "dmr"))
}

func IndexAny() {
	// Finds index of chars of substr
	// If exist any, returns first char index
	// Else, returns -1

	fmt.Printf("%q - %q => %d\n", "chicken", "aeiouy", strings.IndexAny("chicken", "aeiouy"))
	fmt.Printf("%q - %q => %d\n", "crwth", "aeiouy", strings.IndexAny("crwth", "aeiouy"))
}

func IndexByte() {
	// Finds first index of char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "golang", 'g', strings.IndexByte("golang", 'g'))
	fmt.Printf("%q - %q => %d\n", "gophers", 'h', strings.IndexByte("gophers", 'h'))
	fmt.Printf("%q - %q => %d\n", "golang", 'x', strings.IndexByte("golang", 'x'))
}

func IndexFunc() {
	// Finds first index of matched with func
	// If not exist, returns -1

	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}

	fmt.Printf("%q => %d\n", "Hello, 世界", strings.IndexFunc("Hello, 世界", f))
	fmt.Printf("%q => %d\n", "Hello, world", strings.IndexFunc("Hello, world", f))
}

func IndexRune() {
	// Finds first index of rune char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "golang", rune(103), strings.IndexRune("golang", rune(103)))
	fmt.Printf("%q - %q => %d\n", "gophers", rune(104), strings.IndexRune("gophers", rune(104)))
	fmt.Printf("%q - %q => %d\n", "golang", rune(120), strings.IndexRune("golang", rune(120)))
}

func LastIndex() {
	// Finds last index of substr in str
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "go gopher", "go", strings.LastIndex("go gopher", "go"))
	fmt.Printf("%q - %q => %d\n", "go gopher", "rodent", strings.LastIndex("go gopher", "rodent"))
}

func LastIndexAny() {
	// Finds index of chars of substr
	// If exist any, returns last char index
	// Else, returns -1

	fmt.Printf("%q - %q => %d\n", "go gopher", "go", strings.LastIndexAny("go gopher", "go"))
	fmt.Printf("%q - %q => %d\n", "go gopher", "rodent", strings.LastIndexAny("go gopher", "rodent"))
	fmt.Printf("%q - %q => %d\n", "go gopher", "fail", strings.LastIndexAny("go gopher", "fail"))
}

func LastIndexByte() {
	// Finds first index of char
	// If not exist, returns -1

	fmt.Printf("%q - %q => %d\n", "Hello, world", 'l', strings.LastIndexByte("Hello, world", 'l'))
	fmt.Printf("%q - %q => %d\n", "Hello, world", 'o', strings.LastIndexByte("Hello, world", 'o'))
	fmt.Printf("%q - %q => %d\n", "Hello, world", 'x', strings.LastIndexByte("Hello, world", 'x'))
}

func LastIndexFunc() {
	// Finds last index of matched with func
	// If not exist, returns -1

	f := unicode.IsNumber

	fmt.Printf("%q => %d\n", "go 123", strings.LastIndexFunc("go 123", f))
	fmt.Printf("%q => %d\n", "123 go", strings.LastIndexFunc("123 go", f))
	fmt.Printf("%q => %d\n", "go", strings.LastIndexFunc("go", f))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}