package main

import "fmt"

func main() {
	stack := New()
	stack.Push(1)
	stack.Push(3)
	stack.Push(4)
	stack.Pop()
	stack.Push(5)
	stack.Display()
}

type Stack struct {
	Data []interface{}
}

func New() *Stack {
	return &Stack{
		Data: []interface{}{},
	}
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.Data) == 0
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	return stack.Data[len(stack.Data) - 1]
}

func (stack *Stack) Push(item interface{}) *Stack {
	stack.Data = append(stack.Data, item)

	return stack
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	lastIndex := len(stack.Data) - 1

	item := stack.Data[lastIndex]
	stack.Data = stack.Data[:lastIndex]

	return item
}

func (stack *Stack) Display() {
	for _, v := range stack.Data {
		fmt.Println(v)
	}
}