package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	total := 10

	wg.Add(total)

	for i := 1; i <= total; i++ {
		mu.Lock()
		go func(i int) {
			fmt.Printf("%d ", i)
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
}