package main

import (
	"fmt"
)

func main() {
	tree := &BinaryTree{}
	tree.Insert(50)
	tree.Insert(-10)
	tree.Insert(-20)
	tree.Insert(-15)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(-2)
	tree.Insert(0)
	tree.Insert(25)
	tree.Display()
}

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Left  	*Node
	Right 	*Node
	Data  	int
}

func (tree *BinaryTree) Insert(data int) *BinaryTree {
	if tree.Root == nil {
		tree.Root = &Node{Data: data, Left: nil, Right: nil}
	} else {
		tree.Root.Insert(data)
	}

	return tree
}

func (tree *BinaryTree) Display() {
	tree.Root.Display(15)
}

func (node *Node) Insert(data int) {
	if node == nil {
		return
	} else if data <= node.Data {
		if node.Left == nil {
			node.Left = &Node{Data: data, Left: nil, Right: nil}
		} else {
			node.Left.Insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Data: data, Left: nil, Right: nil}
		} else {
			node.Right.Insert(data)
		}
	}
}

func (node *Node) Display(ns int) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Printf(" ")
	}

	fmt.Println(node.Data)

	node.Left.Display(ns - 3)
	node.Right.Display(ns + 3)
}