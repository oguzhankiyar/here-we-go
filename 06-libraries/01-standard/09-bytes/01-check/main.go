package main

import (
	"bytes"
	"fmt"
)

func main() {
	Sample("Compare", Compare)
	Sample("Contains", Contains)
	Sample("ContainsAny", ContainsAny)
	Sample("ContainsRune", ContainsRune)
	Sample("EqualFold", EqualFold)
	Sample("HasPrefix", HasPrefix)
	Sample("HasSuffix", HasSuffix)
}

func Compare() {
	// Compares two string
	// If ordered returns -1
	// If equals returns 0
	// If reverse ordered returns 1

	fmt.Printf("%q - %q => %d\n", "a", "b", bytes.Compare([]byte("a"), []byte("b")))
	fmt.Printf("%q - %q => %d\n", "a", "a", bytes.Compare([]byte("a"), []byte("a")))
	fmt.Printf("%q - %q => %d\n", "b", "a", bytes.Compare([]byte("b"), []byte("a")))
}

func Contains() {
	// Searches substr inside str, if exist case-sensitive, returns true

	str, substr := "This is text", "is"
	result := bytes.Contains([]byte(str), []byte(substr))
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func ContainsAny() {
	// Searches all substr chars, if exist any, returns true

	str, substr := "This is text", "at"
	result := bytes.ContainsAny([]byte(str), substr)
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func ContainsRune() {
	// Searches rune as string in str, if exist, returns true

	str, substr := "This is text", 97
	result := bytes.ContainsRune([]byte(str), rune(97))
	fmt.Printf("str: %q, substr: %q, result: %t\n", str, substr, result)
}

func EqualFold() {
	// Checks equality case insensitive

	fmt.Printf("%q - %q => %t\n", "Go", "go", bytes.EqualFold([]byte("Go"), []byte("go")))
}

func HasPrefix() {
	// Checks the str starts with substr case sensitive

	fmt.Printf("%q - %q => %t\n", "Gopher", "Go", bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Printf("%q - %q => %t\n", "Gopher", "er", bytes.HasPrefix([]byte("Gopher"), []byte("er")))
	fmt.Printf("%q - %q => %t\n", "Gopher", "", bytes.HasPrefix([]byte("Gopher"), []byte("")))
}

func HasSuffix() {
	// Checks the str ends with substr case sensitive

	fmt.Printf("%q - %q => %t\n", "Gopher", "Go", bytes.HasSuffix([]byte("Gopher"), []byte("Go")))
	fmt.Printf("%q - %q => %t\n", "Gopher", "er", bytes.HasSuffix([]byte("Gopher"), []byte("er")))
	fmt.Printf("%q - %q => %t\n", "Gopher", "", bytes.HasSuffix([]byte("Gopher"), []byte("")))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}