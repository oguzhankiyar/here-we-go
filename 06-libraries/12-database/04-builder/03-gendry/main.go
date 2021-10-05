package main

import (
	"fmt"

	"github.com/didi/gendry/builder"
)

func main() {
	where := map[string]interface{}{
		"city": []string{"beijing", "shanghai"},
		"score": 5,
		"age >": 35,
		"address": builder.IsNotNull,
		"_or": []map[string]interface{}{
			{
				"x1":    11,
				"x2 >=": 45,
			},
			{
				"x3":    "234",
				"x4 <>": "tx2",
			},
		},
		"_orderby": "bonus desc",
		"_groupby": "department",
	}
	table := "some_table"
	selectFields := []string{"name", "age", "sex"}
	cond, values, err := builder.BuildSelect(table, where, selectFields)

	if err != nil {
		panic(err)
	}

	fmt.Println("cond:", cond)
	fmt.Println("values:", values)
}