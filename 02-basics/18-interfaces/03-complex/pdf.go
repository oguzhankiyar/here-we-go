package main

import "fmt"

type Pdf struct {
	Text string
}

func (p Pdf) Print() {
	fmt.Println(p.Text)
}