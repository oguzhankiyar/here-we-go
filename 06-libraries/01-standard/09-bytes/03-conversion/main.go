package main

import (
	"bytes"
	"fmt"
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

	fmt.Printf("%q - %q => %s\n", []string{ "Gopher", "Go" }, "|", bytes.Join([][]byte{ []byte("Gopher"), []byte("Go") }, []byte("|")))
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

	fmt.Printf("%q => %s\n", "Gopher", bytes.Map(f, []byte("Gopher")))
}

func Repeat() {
	// Repeats the strings with given count

	fmt.Printf("%q - %d => %s\n", "go", 3, bytes.Repeat([]byte("go"), 3))
}

func Replace() {
	// Replaces old to new in strings with count

	fmt.Printf("%q - %q - %q - %d => %s\n", "Golang is a programming language Golang", "Golang", "Go", 1, bytes.Replace([]byte("Golang is a programming language Golang"), []byte("Golang"), []byte("Go"), 1))
	fmt.Printf("%q - %q - %q - %d => %s\n", "Golang is a programming language Golang", "Golang", "Go", -1, bytes.Replace([]byte("Golang is a programming language Golang"), []byte("Golang"), []byte("Go"), -1))
}

func ReplaceAll() {
	// Replaces all old to new in strings

	fmt.Printf("%q - %q - %q => %s\n", "Golang is a programming language Golang", "Golang", "Go", bytes.ReplaceAll([]byte("Golang is a programming language Golang"), []byte("Golang"), []byte("Go")))
}

func Split() {
	// Split items

	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(","))) // ["a" "b" "c"]
}

func SplitN() {
	// Split n items

	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2)) // ["a" "b,c"]
}

func SplitAfter() {
	// Splits items and keeps separator

	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(","))) // ["a," "b," "c"]
}

func SplitAfterN() {
	// Splits n items and keeps separator

	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2)) // ["a," "b,c"]
}

func Title() {
	// Converts title case

	fmt.Printf("%q\n", bytes.Title([]byte("her royal highness"))) // Her Royal Highness
	fmt.Printf("%q\n", bytes.Title([]byte("loud noises"))) // Loud Noises
	fmt.Printf("%q\n", bytes.Title([]byte("хлеб"))) // Хлеб
}

func ToTitle() {
	// Converts title case

	fmt.Printf("%s\n", bytes.ToTitle([]byte("her royal highness"))) // HER ROYAL HIGHNESS
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises"))) // LOUD NOISES
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб"))) // ХЛЕБ
}

func ToTitleSpecial() {
	// Converts title with special case

	fmt.Printf("%s\n", bytes.ToTitleSpecial(unicode.TurkishCase, []byte("dünyanın ilk borsa yapısı Aizonai kabul edilir")))
	// DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
}

func ToLower() {
	// Converts lower

	fmt.Printf("%s\n", bytes.ToLower([]byte("This is MAGIC"))) // this is magic
}

func ToLowerSpecial() {
	// Converts lower with case

	fmt.Printf("%s\n", bytes.ToLowerSpecial(unicode.TurkishCase, []byte("This is MAGIC"))) // this is magıc
}

func ToUpper() {
	// Converts upper

	fmt.Printf("%s\n", bytes.ToUpper([]byte("This is MAGIC"))) // THIS IS MAGIC
}

func ToUpperSpecial() {
	// Converts upper with case

	fmt.Printf("%s\n", bytes.ToUpperSpecial(unicode.TurkishCase, []byte("This is MAGIC"))) // THİS İS MAGIC
}

func ToValidUTF8() {
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("Hello \xc5"), []byte("there"))) // Hello there
}

func Trim() {
	// Trims given chars from left and right

	fmt.Printf("%s\n", bytes.Trim([]byte("¡¡¡Hello, Gophers!!!"), "!¡e")) // Hello, Gophers
}

func TrimFunc() {
	// Trims with given func from left and right

	fmt.Printf("%s\n", bytes.TrimFunc([]byte("55 H3ll0 7"), unicode.IsNumber)) // H3ll0
}

func TrimSpace() {
	// Trims spaces from left and right

	fmt.Printf("%s\n", bytes.TrimSpace([]byte(" This is MAGIC  "))) // This is MAGIC
}

func TrimLeft() {
	// Trims given chars from left

	fmt.Printf("%s\n", bytes.TrimLeft([]byte("55 H3ll0 7"), "5 ")) // H3ll0 7
}

func TrimLeftFunc() {
	// Trims with given func from left

	fmt.Printf("%s\n", bytes.TrimFunc([]byte("55 H3ll0 7"), unicode.IsNumber)) // H3ll0 7
}

func TrimRight() {
	// Trims given chars from right

	fmt.Printf("%s\n", bytes.TrimRight([]byte("55 H3ll0 7"), "7 ")) // 55 H3ll0
}

func TrimRightFunc() {
	// Trims with given func from right

	fmt.Printf("%s\n", bytes.TrimRightFunc([]byte("55 H3ll0 7"), unicode.IsNumber)) // 55 H3ll0
}

func TrimPrefix() {
	// Trims given word from left

	fmt.Printf("%s\n", bytes.TrimPrefix([]byte("H3ll0 7"), []byte("H3l"))) // l0 7
}

func TrimSuffix() {
	// Trims given word from right

	fmt.Printf("%s\n", bytes.TrimSuffix([]byte("H3ll0 7"), []byte("0 7"))) // H3ll
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}