package main

import (
	"fmt"

	"github.com/kelindar/column"
)

func main() {
	data := []map[string]interface{}{
		{"name": "gopher"},
		{"name": "alice"},
		{"name": "bob"},
	}

	players := column.NewCollection()
	players.CreateColumnsOf(data[0])

	for _, v := range data {
		players.Insert(v)
	}

	err := players.Query(func(txn *column.Txn) error {
		txn.WithString("name", func(v string) bool {
			return v == "gopher"
		}).Select(func(v column.Selector) {
			name := v.ValueAt("name")
			fmt.Printf("%s\n", name)
		})

		return nil
	})

	if err != nil {
		panic(err)
	}
}