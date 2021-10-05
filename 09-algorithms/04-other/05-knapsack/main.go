package main

import (
	"fmt"
	"math"
)

func main() {
	items := []Item{
		{"item-1", 2, 1},
		{"item-2", 2, 2},
		{"item-3", 1, 10},
		{"item-4", 1, 10},
		{"item-5", 2, 5},
		{"item-6", 6, 1},
	}

	totalWeight, totalValue, result := Solve(items, 10)

	for _, v := range result {
		fmt.Printf("name: %s | weight: %v | value: %v\n", v.Name, v.Value, v.Weight)
	}

	fmt.Printf("total %v items with weight %v value %v\n", len(result), totalWeight, totalValue)
}

type Item struct {
	Name   string
	Value  float64
	Weight float64
}

func Solve(items []Item, capacity float64) (float64, float64, []Item) {
	totalValue := 0.0
	totalWeight := 0.0
	result := make([]Item, 0)

	sets := Check(items)

	for _, set := range sets {
		weight, value := Calculate(set)
		if weight <= capacity && value > totalValue {
			totalWeight = weight
			totalValue = value
			result = set
		}
	}

	return totalWeight, totalValue, result
}

func Check(items []Item) [][]Item {
	sets := make([][]Item, 0)

	count := int(math.Pow(2., float64(len(items))))

	for i := 0; i < count; i++ {
		set := make([]Item, 0)
		for j := 0; j < len(items); j++ {
			if (i >> uint(j)) & 1 == 1 {
				set = append(set, items[j])
			}
		}
		sets = append(sets, set)
	}

	return sets
}

func Calculate(set []Item) (weight float64, value float64) {
	for _, i := range set {
		weight += i.Weight
		value += i.Value
	}

	return
}