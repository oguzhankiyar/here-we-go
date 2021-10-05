package main

import "fmt"

type Printer struct {

}

func (p Printer) Print(data string) {
	fmt.Println("printed:", data)
}

type Scanner struct {

}

func (s Scanner) Scan() string {
	data := "some_text"
	fmt.Println("scanned:", data)
	return data
}

type Computer struct {
	printer Printer
	scanner Scanner
}

func (c Computer) Copy() {
	data := c.scanner.Scan()
	c.printer.Print(data)
}

func main() {
	printer := Printer{}
	scanner := Scanner{}
	computer := Computer{printer, scanner}

	computer.Copy()
}