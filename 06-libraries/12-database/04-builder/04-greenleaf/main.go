package main

import (
	"fmt"

	"github.com/slavabobik/greenleaf"
)

func main() {
	filter := greenleaf.
		Filter().
		EqString("name", "Jhon").
		InString("tags", []string{"fast", "furious"}).
		GtInt("score", 100).
		LteInt("score", 200).
		Exists("active", true).
		Build()

	fmt.Println("filter:", filter)

	update := greenleaf.
		Update().
		SetBool("is_active", true).
		SetIntSlice("numbers", []int{1, 2, 3, 4}).
		Build()

	fmt.Println("update:", update)
}