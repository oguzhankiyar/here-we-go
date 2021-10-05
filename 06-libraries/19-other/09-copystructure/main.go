package main

import (
	"fmt"

	"github.com/mitchellh/copystructure"
)

func main() {
	input := map[string]interface{}{
		"bob": map[string]interface{}{
			"emails": []string{"a", "b"},
		},
	}

	dup, err := copystructure.Copy(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", dup)
}