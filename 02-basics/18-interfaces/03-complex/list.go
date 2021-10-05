package main

type List []Printable

func (list List) Print() {
	for _, v := range list {
		v.Print()
	}
}