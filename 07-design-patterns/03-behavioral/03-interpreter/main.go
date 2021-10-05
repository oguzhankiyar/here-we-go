package main

import "fmt"

type Expression interface {
	Interpret(str string) bool
}

type ContainExpression struct {
	Data []string
}

func (e ContainExpression) Interpret(data string) bool {
	for _, d := range e.Data {
		if d == data {
			return true
		}
	}
	return false
}

type AndExpression struct {
	Left 	Expression
	Right 	Expression
}

func (a AndExpression) Interpret(data string) bool {
	return a.Left.Interpret(data) && a.Right.Interpret(data)
}

type OrExpression struct {
	Left 	Expression
	Right 	Expression
}

func (a OrExpression) Interpret(data string) bool {
	return a.Left.Interpret(data) || a.Right.Interpret(data)
}

func main() {
	leftContainExp := ContainExpression{[]string{ "one", "two", "three" }}
	rightContainExp := ContainExpression{[]string{ "three", "four", "five" }}

	andExp := AndExpression{leftContainExp, rightContainExp}
	orExp := OrExpression{leftContainExp, rightContainExp}

	fmt.Println("left:", leftContainExp.Data)
	fmt.Println("right:", rightContainExp.Data)

	fmt.Println()

	fmt.Println("andExp one:", andExp.Interpret("one"))
	fmt.Println("andExp three:", andExp.Interpret("three"))

	fmt.Println("orExp one:", orExp.Interpret("one"))
	fmt.Println("orExp six:", orExp.Interpret("six"))
}