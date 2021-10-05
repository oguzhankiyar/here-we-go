package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte("cos(x) + 1i*sin(x) // Euler")

	var s scanner.Scanner
	fs := token.NewFileSet()
	f := fs.AddFile("", fs.Base(), len(src))

	s.Init(f, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fs.Position(pos), tok, lit)
	}
}