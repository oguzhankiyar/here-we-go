package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	Sample("New", New)
	Sample("Sum", Sum)
}

func New() {
	str := "123456"

	hash := md5.New()
	io.WriteString(hash, str)
	result := hash.Sum(nil)

	fmt.Printf("%s -> %x\n", str, result) // e10adc3949ba59abbe56e057f20f883e
}

func Sum() {
	str := "123456"
	result := md5.Sum([]byte(str))

	fmt.Printf("%s -> %x\n", str, result) // e10adc3949ba59abbe56e057f20f883e
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
