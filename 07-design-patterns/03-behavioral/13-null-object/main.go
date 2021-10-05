package main

import "fmt"

type Product interface {
	GetCode() string
	GetName() string
}

type Shoe struct {
	Code string
	Name string
}

func (s Shoe) GetCode() string {
	return s.Code
}

func (s Shoe) GetName() string {
	return s.Name
}

type NullProduct struct {

}

func (s NullProduct) GetCode() string {
	return ""
}

func (n NullProduct) GetName() string {
	return ""
}

func Find(products []Product, name string) Product {
	for _, v := range products {
		if v.GetName() == name {
			return v
		}
	}

	return NullProduct{}
}

func main() {
	products := []Product{
		Shoe{"a-100", "adidas"},
		Shoe{"n-110", "nike"},
		Shoe{"p-123", "puma"},
	}

	fn := func (name string) {
		result := Find(products, name)
		switch result.(type) {
		case NullProduct:
			fmt.Printf("%s not found\n", name)
		default:
			fmt.Printf("[%s] %s\n", result.GetCode(), result.GetName())
		}
	}

	fn("puma")
	fn("converse")
}