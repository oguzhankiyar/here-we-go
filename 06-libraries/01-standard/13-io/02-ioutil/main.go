package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	Sample("ReadAll", ReadAll)
	Sample("ReadDir", ReadDir)
	Sample("ReadFile", ReadFile)
	Sample("TempDir", TempDir)
	Sample("TempFile", TempFile)
	Sample("WriteFile", WriteFile)
}

func ReadAll() {
	r := strings.NewReader("Hello there! I am Gopher.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", b) // Hello there! I am Gopher.
}

func ReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, file := range files {
		fmt.Println(file.Name()) // main.go.go
	}
}

func ReadFile() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		fmt.Println("error:", err) // error: open testdata/hello: The system cannot find the path specified.
		return
	}

	fmt.Printf("File contents: %s\n", content)
}

func TempDir() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")

	if err != nil {
		fmt.Println("error:", err)
	}

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		fmt.Println("error:", err)
	}

	content, err = ioutil.ReadFile(dir + "/tmpfile")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", content) // temporary file's content
}

func TempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		fmt.Println("error:", err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		fmt.Println("error:", err)
	}

	if err := tmpfile.Close(); err != nil {
		fmt.Println("error:", err)
	}

	content, err = ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", content) // temporary file's content
}

func WriteFile() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("hello.txt", message, 0644)
	if err != nil {
		fmt.Println("error:", err)
	}

	defer os.Remove("hello.txt")

	content, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", content) // Hello, Gophers!
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
