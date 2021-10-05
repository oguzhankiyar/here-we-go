package main

import (
	"fmt"
)

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"C", "D"},
		"C": {"E"},
		"D": {"E"},
		"E": {},
	}

	path := FindPath(graph, "A", "E", make([]string, 0))

	fmt.Println(path)
}

func FindPath(graph map[string][]string, from string, to string, visits []string) []string {
	if _, exist := graph[from]; !exist {
		return visits
	}

	visits = append(visits, from)

	if from == to {
		return visits
	}

	shortest := make([]string, 0)

	for _, node := range graph[from] {
		if !IsVisited(visits, node) {
			path := FindPath(graph, node, to, visits)

			if len(path) == 0 {
				continue
			}

			if len(shortest) == 0 || len(path) < len(shortest) {
				shortest = path
			}
		}
	}

	return shortest
}

func IsVisited(visits []string, node string) bool {
	for _, v := range visits {
		if node == v {
			return true
		}
	}

	return false
}