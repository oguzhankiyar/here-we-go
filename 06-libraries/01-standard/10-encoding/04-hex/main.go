package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	Sample("Encode", Encode)
	Sample("Decode", Decode)
	Sample("EncodeToString", EncodeToString)
	Sample("DecodeString", DecodeString)
}

func Encode() {
	src := []byte("Hello Gopher!")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)
}

func Decode() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])
}

func EncodeToString() {
	src := "Hello there!"
	encoded := hex.EncodeToString([]byte(src))
	fmt.Printf("%q -> %q\n", src, encoded)
}

func DecodeString() {
	src := "48656c6c6f20746865726521"
	decoded, err := hex.DecodeString(src)
	fmt.Printf("%q -> %q %v\n", src, decoded, err)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}