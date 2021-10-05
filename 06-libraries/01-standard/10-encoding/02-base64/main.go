package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	Sample("Encoder", Encoder)
	Sample("Decoder", Decoder)
	Sample("Encode", Encode)
	Sample("Decode", Decode)
}

func Encoder() {
	enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	enc.Write([]byte("Hello @Gopher!")) // SGVsbG8gQEdvcGhlciE=
	enc.Close()

	fmt.Println()
}

func Decoder() {
	src := []byte("SGVsbG8gQEdvcGhlciE=")
	dec := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(src))

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(src)))

	n, err := dec.Read(dst)

	fmt.Println(string(dst), n, err) // Hello @Gopher!  14 <nil>
}

func Encode() {
	encoding := base64.StdEncoding

	data := encoding.EncodeToString([]byte("Hello, Gopher!"))

	fmt.Printf("%q\n", data) // "SGVsbG8sIEdvcGhlciE="
}

func Decode() {
	encoding := base64.StdEncoding

	data, err := encoding.DecodeString("SGVsbG8sIEdvcGhlciE=")

	fmt.Printf("%q %v\n", data, err) // "Hello, Gopher!" <nil>
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}