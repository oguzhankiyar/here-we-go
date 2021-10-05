package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	zw.Name = "golang.txt"
	zw.Comment = "the go programming language"
	zw.ModTime = time.Date(2021, time.May, 28, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("Go is a programming language for Gophers"))
	if err != nil {
		fmt.Println("error:", err)
	}

	if err := zw.Close(); err != nil {
		fmt.Println("error:", err)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		fmt.Println("error:", err)
	}

	if err := zr.Close(); err != nil {
		fmt.Println("error:", err)
	}
}