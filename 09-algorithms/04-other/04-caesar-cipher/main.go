package main

import "fmt"

func main() {
	str := "Hi, I'm Gopher!"
	shift := 7

	fmt.Println("original:", str)

	encoded := Encode(str, shift)
	fmt.Println("encoded:", encoded)

	decoded := Decode(encoded, shift)
	fmt.Println("decoded:", decoded)
}

func Encode(str string, shift int) string {
	runes := []rune(str)

	for i, r := range runes {
		runes[i] = Cipher(r, rune(shift))
	}

	return string(runes)
}

func Decode(str string, shift int) string {
	runes := []rune(str)

	for i, r := range runes {
		runes[i] = Cipher(r, rune(26 - shift))
	}

	return string(runes)
}

func Cipher(char rune, shift rune) rune {
	if char >= 'a' && char <= 'z' {
		return (((char + shift) - 'a') % 26) + 'a'
	}

	if char >= 'A' && char <= 'Z' {
		return (((char + shift) - 'A') % 26) + 'A'
	}

	return char
}