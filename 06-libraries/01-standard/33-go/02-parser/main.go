package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	fs := token.NewFileSet() // positions are relative to fset

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	f, err := parser.ParseFile(fs, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}