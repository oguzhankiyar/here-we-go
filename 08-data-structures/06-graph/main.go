package main

import (
	"fmt"
)

func main() {
	var graph Graph

	a := Node{"A"}
	b := Node{"B"}
	c := Node{"C"}
	d := Node{"D"}
	e := Node{"E"}
	f := Node{"F"}

	graph.AddNode(&a)
	graph.AddNode(&b)
	graph.AddNode(&c)
	graph.AddNode(&d)
	graph.AddNode(&e)
	graph.AddNode(&f)

	graph.AddEdge(&a, &b)
	graph.AddEdge(&a, &c)
	graph.AddEdge(&b, &e)
	graph.AddEdge(&c, &e)
	graph.AddEdge(&e, &f)
	graph.AddEdge(&d, &a)

	graph.Display()
}

type Node struct {
	Value interface{}
}

type Graph struct {
	Nodes 	[]*Node
	Edges 	map[Node][]*Node
}

func (graph *Graph) AddNode(n *Node) {
	graph.Nodes = append(graph.Nodes, n)
}

func (graph *Graph) AddEdge(n1, n2 *Node) {
	if graph.Edges == nil {
		graph.Edges = make(map[Node][]*Node)
	}

	graph.Edges[*n1] = append(graph.Edges[*n1], n2)
	graph.Edges[*n2] = append(graph.Edges[*n2], n1)
}

func (graph *Graph) Display() {
	for i := 0; i < len(graph.Nodes); i++ {
		line := fmt.Sprintf("%s -> ", graph.Nodes[i].Value)
		near := graph.Edges[*graph.Nodes[i]]

		for j := 0; j < len(near); j++ {
			line += fmt.Sprintf("%s ", near[j].Value)
		}

		fmt.Println(line)
	}
}