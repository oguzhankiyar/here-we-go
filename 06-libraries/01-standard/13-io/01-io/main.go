package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	Sample("ReadAll", ReadAll)
	Sample("ReadAtLeast", ReadAtLeast)
	Sample("ReadFull", ReadFull)
	Sample("WriteString", WriteString)
	Sample("Copy", Copy)
	Sample("CopyN", CopyN)
	Sample("CopyBuffer", CopyBuffer)
	Sample("Pipe", Pipe)
	Sample("MultiWriter", MultiWriter)
	Sample("MultiReader", MultiReader)
	Sample("SectionReader", SectionReader)
	Sample("LimitReader", LimitReader)
	Sample("TeeReader", TeeReader)
	Sample("Discard", Discard)
	Sample("EOF", EOF)
}

func ReadAll() {
	reader := strings.NewReader("Hello, there!")

	data, err := io.ReadAll(reader)
	fmt.Printf("%s %v\n", data, err) // Hello, there! <nil>
}

func ReadAtLeast() {
	r := strings.NewReader("Hello, there! I am Gopher. Who are you?\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", buf) // Hello, there! <nil>

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err) // error: short buffer
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err) // error: unexpected EOF
	}
}

func ReadFull() {
	r := strings.NewReader("Hello, there! I am Gopher. Who are you?\n")

	buf := make([]byte, 5)
	if _, err := io.ReadFull(r, buf); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", buf) // Hell

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err) // error: unexpected EOF
	}
}

func WriteString() {
	io.WriteString(os.Stdout, "Hello there! I am Gopher.") // Hello there! I am Gopher.

	fmt.Println()
}

func Copy() {
	reader := strings.NewReader("Hello there! I am Gopher.")

	written, err := io.Copy(os.Stdout, reader) // Hello there! I am Gopher.

	fmt.Println()
	fmt.Println(written, err) // 25 <nil>
}

func CopyN() {
	reader := strings.NewReader("Hello there! I am Gopher.")

	written, err := io.CopyN(os.Stdout, reader, 20) // Hello there! I am Go

	fmt.Println()
	fmt.Println(written, err) // 20 <nil>
}

func CopyBuffer() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		fmt.Println("error:", err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		fmt.Println("error:", err)
	}
}

func Pipe() {
	r, w := io.Pipe()

	go func() {
		time.Sleep(time.Second * 1)
		fmt.Fprint(w, "Hello there! I am Gopher.\n")
		w.Close()
	}()

	// Waits when other is closed
	_, err := io.Copy(os.Stdout, r) // Hello there! I am Gopher.
	if err != nil {
		fmt.Println("error:", err)
	}
}

func MultiWriter() {
	r := strings.NewReader("Hello there! I am Gopher.\n")

	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)

	_, err := io.Copy(w, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String()) // Hello there! I am Gopher.
	fmt.Print(buf2.String()) // Hello there! I am Gopher.
}

func MultiReader() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	_, err := io.Copy(os.Stdout, r) // first reader second reader third reader
	if err != nil {
		fmt.Println("error:", err)
	}
}

func SectionReader() {
	r := strings.NewReader("Hello there! I am Gopher.\n")

	sr := io.NewSectionReader(r, 0, 12)

	_, err := io.Copy(os.Stdout, sr) // Hello there!
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println()
}

func LimitReader() {
	r := strings.NewReader("Hello there! I am Gopher.\n")

	lr := io.LimitReader(r, 12)

	_, err := io.Copy(os.Stdout, lr) // Hello there!
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println()
}

func TeeReader() {
	var r io.Reader = strings.NewReader("Hello there! I am Gopher.\n")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	// If this line is deleted, nothing to send stdout
	io.ReadAll(r)
}

func Discard() {
	r := strings.NewReader("Hello!")
	w := io.MultiWriter(io.Discard, os.Stdout)

	io.Copy(w, r)

	fmt.Println()
}

func EOF() {
	reader := strings.NewReader("Hello!")

	i := 0

	for {
		b := make([]byte, 1)

		_, err := reader.ReadAt(b, int64(i))
		if err == io.EOF {
			fmt.Println()
			fmt.Println("<EOF>")
			break
		}

		fmt.Printf("%s", b)
		i++
	}

	/*
		Hello!
		<EOF>
	*/
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
