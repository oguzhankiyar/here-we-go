package main

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	query	string
	values map[string]string
}

func New() *QueryBuilder {
	return &QueryBuilder{"", make(map[string]string)}
}

func (builder *QueryBuilder) Add(key, value string) *QueryBuilder {
	builder.values[key] = value

	return builder
}

func (builder *QueryBuilder) Build() string {
	query := ""

	for k, v := range builder.values {
		query += fmt.Sprintf("%s=%s&", k, v)
	}

	return strings.TrimRight(query, "&")
}

func main() {
	builder := New()

	query := builder.
		Add("id", "28").
		Add("page", "1").
		Add("size", "10").
		Build()

	fmt.Println("query:", query)
}