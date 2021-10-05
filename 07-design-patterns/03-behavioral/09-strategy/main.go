package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type SortStrategy interface {
	Sort([]string) []string
}

type RandomSortStrategy struct {

}

func (s RandomSortStrategy) Sort(data []string) []string {
	var result []string
	indexes := make(map[int]int)

	for i := 0; i < len(data); i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Int() % len(data)

		if _, ok := indexes[index]; ok {
			i--
			continue
		}

		indexes[index] = i

		result = append(result, data[index])
	}

	return result
}

type AlphabeticalSortStrategy struct {

}

func (s AlphabeticalSortStrategy) Sort(data []string) []string {
	sort.Strings(data)
	return data
}

type Sorter struct {
	SortStrategy 	SortStrategy
	Data 			[]string
}

func (l *Sorter) Sort() {
	l.Data = l.SortStrategy.Sort(l.Data)
}

func main() {
	randomSortStrategy := RandomSortStrategy{}
	alphabeticalSortStrategy := AlphabeticalSortStrategy{}

	items := []string{"this", "is", "go", "programming", "language"}

	randomSorter := Sorter{randomSortStrategy, items}
	alphabeticalSorter := Sorter{alphabeticalSortStrategy, items}

	randomSorter.Sort()
	alphabeticalSorter.Sort()

	fmt.Println("Random:", randomSorter.Data)
	fmt.Println("Alphabetical:", alphabeticalSorter.Data)
}