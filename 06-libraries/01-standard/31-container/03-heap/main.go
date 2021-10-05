package main

import (
	"container/heap"
	"fmt"
)

func main() {
	Sample("Heap", Heap)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	temp := *h
	n := len(temp)
	last := temp[n-1]
	*h = temp[0 : n-1]

	return last
}

func Heap() {
	h := &IntHeap{2, 1, 5}

	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("min:", (*h)[0])

	fmt.Printf("items: ")

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Println()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}