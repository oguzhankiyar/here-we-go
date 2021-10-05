package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	Sample("Join", Join)
	Sample("Map", Map)
	Sample("Repeat", Repeat)
	Sample("Replace", Replace)
	Sample("ReplaceAll", ReplaceAll)
	Sample("Split", Split)
	Sample("SplitN", SplitN)
	Sample("SplitAfter", SplitAfter)
	Sample("SplitAfterN", SplitAfterN)
	Sample("Title", Title)
	Sample("ToTitle", ToTitle)
	Sample("ToTitleSpecial", ToTitleSpecial)
	Sample("ToLower", ToLower)
	Sample("ToLowerSpecial", ToLowerSpecial)
	Sample("ToUpper", ToUpper)
	Sample("ToUpperSpecial", ToUpperSpecial)
	Sample("ToValidUTF8", ToValidUTF8)
	Sample("Trim", Trim)
	Sample("TrimFunc", TrimFunc)
	Sample("TrimSpace", TrimSpace)
	Sample("TrimLeft", TrimLeft)
	Sample("TrimLeftFunc", TrimLeftFunc)
	Sample("TrimRight", TrimRight)
	Sample("TrimRightFunc", TrimRightFunc)
	Sample("TrimPrefix", TrimPrefix)
	Sample("TrimSuffix", TrimSuffix)
}

func Join() {
	// Joins strings with given separator

	fmt.Printf("%q - %q => %v\n", []string{ "Gopher", "Go" }, "|", strings.Join([]string{ "Gopher", "Go" }, "|"))
}

func Map() {
	// Maps the string with given func

	f := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r - 'A' + 13) % 26
		case r >= 'a' && r <= 'z':
			return 'a' + (r - 'a' + 13) % 26
		}
		return r
	}

	fmt.Printf("%q => %v\n", "Gopher", strings.Map(f, "Gopher"))
}

func Repeat() {
	// Repeats the strings with given count

	fmt.Printf("%q - %d => %v\n", "go", 3, strings.Repeat("go", 3))
}

func Replace() {
	// Replaces old to new in strings with count

	fmt.Printf("%q - %q - %q - %d => %v\n", "Golang is a programming language Golang", "Golang", "Go", 1, strings.Replace("Golang is a programming language Golang", "Golang", "Go", 1))
	fmt.Printf("%q - %q - %q - %d => %v\n", "Golang is a programming language Golang", "Golang", "Go", -1, strings.Replace("Golang is a programming language Golang", "Golang", "Go", -1))
}

func ReplaceAll() {
	// Replaces all old to new in strings

	fmt.Printf("%q - %q - %q => %v\n", "Golang is a programming language Golang", "Golang", "Go", strings.ReplaceAll("Golang is a programming language Golang", "Golang", "Go"))
}

func Split() {
	// Split items

	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) // ["a" "b" "c"]
}

func SplitN() {
	// Split n items

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
}

func SplitAfter() {
	// Splits items and keeps separator

	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
}

func SplitAfterN() {
	// Splits n items and keeps separator

	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
}

func Title() {
	// Converts title case

	fmt.Println(strings.Title("her royal highness")) // Her Royal Highness
	fmt.Println(strings.Title("loud noises")) // Loud Noises
	fmt.Println(strings.Title("хлеб")) // Хлеб
}

func ToTitle() {
	// Converts title case

	fmt.Println(strings.ToTitle("her royal highness")) // HER ROYAL HIGHNESS
	fmt.Println(strings.ToTitle("loud noises")) // LOUD NOISES
	fmt.Println(strings.ToTitle("хлеб")) // ХЛЕБ
}

func ToTitleSpecial() {
	// Converts title with special case

	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	// DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
}

func ToLower() {
	// Converts lower

	fmt.Println(strings.ToLower("This is MAGIC")) // this is magic
}

func ToLowerSpecial() {
	// Converts lower with case

	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "This is MAGIC")) // this is magıc
}

func ToUpper() {
	// Converts upper

	fmt.Println(strings.ToUpper("This is MAGIC")) // THIS IS MAGIC
}

func ToUpperSpecial() {
	// Converts upper with case

	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "This is MAGIC")) // THİS İS MAGIC
}

func ToValidUTF8() {
	fmt.Println(strings.ToValidUTF8("Hello \xc5", "there")) // Hello there
}

func Trim() {
	// Trims given chars from left and right

	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡e")) // Hello, Gophers
}

func TrimFunc() {
	// Trims with given func from left and right

	fmt.Println(strings.TrimFunc("55 H3ll0 7", unicode.IsNumber)) // H3ll0
}

func TrimSpace() {
	// Trims spaces from left and right

	fmt.Println(strings.TrimSpace(" This is MAGIC  ")) // This is MAGIC
}

func TrimLeft() {
	// Trims given chars from left

	fmt.Println(strings.TrimLeft("55 H3ll0 7", "5 ")) // H3ll0 7
}

func TrimLeftFunc() {
	// Trims with given func from left

	fmt.Println(strings.TrimFunc("55 H3ll0 7", unicode.IsNumber)) // H3ll0 7
}

func TrimRight() {
	// Trims given chars from right

	fmt.Println(strings.TrimRight("55 H3ll0 7", "7 ")) // 55 H3ll0
}

func TrimRightFunc() {
	// Trims with given func from right

	fmt.Println(strings.TrimRightFunc("55 H3ll0 7", unicode.IsNumber)) // 55 H3ll0
}

func TrimPrefix() {
	// Trims given word from left

	fmt.Println(strings.TrimPrefix("H3ll0 7", "H3l")) // l0 7
}

func TrimSuffix() {
	// Trims given word from right

	fmt.Println(strings.TrimSuffix("H3ll0 7", "0 7")) // H3ll
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}