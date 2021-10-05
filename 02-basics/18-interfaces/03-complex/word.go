package main

import "fmt"

type Word struct {
	Text string
}

func (w Word) Print() {
	fmt.Println(w.Text)
}