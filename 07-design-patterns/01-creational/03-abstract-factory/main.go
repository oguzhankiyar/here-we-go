package main

import (
	"errors"
	"fmt"
)

type DrinkType string

const (
	TEA		DrinkType = "tea"
	COFFEE	DrinkType = "coffee"
)

type Drink interface {
	GetName() string
	GetDegree() int
}

type DrinkMaker interface {
	Make(int) Drink
}

func GetFactory(t DrinkType) (DrinkMaker, error) {
	if t == TEA {
		return TeaMaker{}, nil
	} else if t == COFFEE {
		return CoffeeMaker{}, nil
	}

	return nil, errors.New("undefined type")
}

type Tea struct {
	name	string
	degree 	int
}

func (t Tea) GetName() string {
	return t.name
}

func (t Tea) GetDegree() int {
	return t.degree
}

type TeaMaker struct {

}

func (m TeaMaker) Make(degree int) Drink {
	return Tea{string(TEA), degree}
}

type Coffee struct {
	name	string
	degree 	int
}

func (t Coffee) GetName() string {
	return t.name
}

func (t Coffee) GetDegree() int {
	return t.degree
}

type CoffeeMaker struct {

}

func (m CoffeeMaker) Make(degree int) Drink {
	return Tea{string(COFFEE), degree}
}

func main() {
	maker, err := GetFactory(COFFEE)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	drink := maker.Make(75)

	fmt.Println("name:", drink.GetName())
	fmt.Println("degree:", drink.GetDegree())
}