package main

import (
	"fmt"
	"unicode"
)

func main() {
	Sample("Is", Is)
	Sample("Case", Case)
}

func Is() {
	fmt.Println("IsLower('A') ->", unicode.IsLower('A'))
	fmt.Println("IsUpper('A') ->", unicode.IsUpper('A'))
	fmt.Println("IsNumber('8') ->", unicode.IsNumber('8'))
	fmt.Println("IsSpace(' ') ->", unicode.IsSpace(' '))
	fmt.Println("IsLetter('5') ->", unicode.IsLetter('5'))
	fmt.Println("IsTitle('a') ->", unicode.IsTitle('a'))
	fmt.Println("Is(unicode.Space, ' ') ->", unicode.Is(unicode.Space, ' '))
}

func Case() {
	fmt.Println("ToLower('A') ->", string(unicode.ToLower('A')))
	fmt.Println("ToUpper('a') ->", string(unicode.ToUpper('a')))
	fmt.Println("ToTitle('a') ->", string(unicode.ToTitle('a')))
	fmt.Println("To(unicode.UpperCase, 'i') ->", string(unicode.To(unicode.UpperCase, 'i')))
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}