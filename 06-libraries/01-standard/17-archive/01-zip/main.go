package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	Sample("Writer", Writer)
	Sample("Reader", Reader)
}

func Writer() {
	buffer := new(bytes.Buffer)
	writer := zip.NewWriter(buffer)

	files := []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := writer.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := writer.Close()
	if err != nil {
		fmt.Println("error:", err)
	}

	err = ioutil.WriteFile("test.zip", buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func Reader() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Println("-", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println("error:", err)
		}

		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			fmt.Println("error:", err)
		}

		rc.Close()
		fmt.Println()
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
