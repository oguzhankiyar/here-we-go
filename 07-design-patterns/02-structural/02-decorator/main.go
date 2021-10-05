package main

import "fmt"

type Product interface {
	GetPrice() int
}

type Jean struct {

}

func (j Jean) GetPrice() int {
	return 50
}

type WithRedColor struct {
	product Product
}

func (j WithRedColor) GetPrice() int {
	return j.product.GetPrice() + 10
}

type WithLargeSize struct {
	product Product
}

func (j WithLargeSize) GetPrice() int {
	return j.product.GetPrice() + 15
}

func main() {
	jean := Jean{}

	redJean := WithRedColor{jean}
	redLargeJean := WithLargeSize{redJean}

	fmt.Printf("Red/Large Jean Price: $%v\n", redLargeJean.GetPrice())
}