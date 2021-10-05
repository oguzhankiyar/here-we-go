package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

func main() {
	const src = `
// This is the package comment.
package ppp

import "fmt"

// This comment is associated with the Greet function.
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
`
	const test = `
package p_test

// This comment is associated with the ExampleGreet_world example.
func ExampleGreet_world() {
	Greet("world")
}
`

	fs := token.NewFileSet()
	files := []*ast.File{
		mustParse(fs, "src.go", src),
		mustParse(fs, "src_test.go", test),
	}

	p, err := doc.NewFromFiles(fs, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	fmt.Printf("package %s - %s", p.Name, p.Doc)
	fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc)
	fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc)
}

func mustParse(fs *token.FileSet, filename, src string) *ast.File {
	f, err := parser.ParseFile(fs, filename, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	return f
}