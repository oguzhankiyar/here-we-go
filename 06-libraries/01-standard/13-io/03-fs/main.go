package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	Sample("ReadDir", ReadDir)
	Sample("ReadFile", ReadFile)
	Sample("Glob", Glob)
	Sample("Stat", Stat)
	Sample("Sub", Sub)
	Sample("ValidPath", ValidPath)
	Sample("WalkDir", WalkDir)
}

func ReadDir() {
	dirFs := os.DirFS("test")
	dirs, err := fs.ReadDir(dirFs, ".")

	if err != nil {
		fmt.Println("error:", err)
	}

	for _, entry := range dirs {
		fmt.Println(entry.Name())
	}
	/*
		test.txt
		two
	*/
}

func ReadFile() {
	dirFs := os.DirFS("test")
	content, err := fs.ReadFile(dirFs, "test.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", content) // Hello, there!
}

func Glob() {
	dirFs := os.DirFS("test")
	matches, err := fs.Glob(dirFs, "*.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(matches) // [test.txt]
}

func Stat() {
	dirFs := os.DirFS("test")
	file, err := fs.Stat(dirFs, "test.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(file.Name(), file.Size()) // test.txt 13
}

func Sub() {
	dirFs := os.DirFS(".")
	sub, err := fs.Sub(dirFs, "test")

	if err != nil {
		fmt.Println("error:", err)
	}

	file, err := sub.Open("test.txt")

	if err != nil {
		fmt.Println("error:", err)
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(info.Name(), info.Size()) // test.txt 13
}

func ValidPath() {
	fn := func(p string) {
		fmt.Printf("%s -> %t\n", p, fs.ValidPath(p))
	}

	fn("test")          // true
	fn("test/test.txt") // true
	fn("example")       // true
	fn("../")           // false
	fn(".")             // true
}

func WalkDir() {
	dirFs := os.DirFS("test")
	fs.WalkDir(dirFs, "two/two.json", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("visited file or dir: %q\n", path) // visited file or dir: "two/two.json"
		return nil
	})
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
