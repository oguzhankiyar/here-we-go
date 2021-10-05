package main

import (
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	Sample("New", New)
	Sample("Sum", Sum)
}

func New() {
	str := "123456"

	hash := sha256.New()
	io.WriteString(hash, str)
	result := hash.Sum(nil)

	fmt.Printf("%s -> %x\n", str, result) // 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
}

func Sum() {
	str := "123456"
	result := sha256.Sum256([]byte(str))

	fmt.Printf("%s -> %x\n", str, result) // 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
