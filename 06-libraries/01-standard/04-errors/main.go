package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	Sample("New", New)
	Sample("Is", Is)
	Sample("As", As)
	Sample("Unwrap", Unwrap)
}

func New() {
	// Creates new error with given message

	err := errors.New("this is error")
	fmt.Println(err)
}

func Is() {
	// Checks two errors

	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}

func As() {
	// Finds matched error in chain

	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func Unwrap() {
	// Unwraps error

	if _, err := os.Open("non-existing"); err != nil {
		fmt.Println(err) // open non-existing: The system cannot find the file specified.
		if err = errors.Unwrap(err); err != nil {
			fmt.Println(err) // The system cannot find the file specified.
		}
	}
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}