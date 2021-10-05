package main

import "fmt"

type Post struct {
	Id		string
	Title 	string
}

func (p Post) String() string {
	return p.Id + " " + p.Title
}

func main() {
	post := Post{
		Id: "1",
		Title: "Golang",
	}

	fmt.Println(getString(post))
}

func getString(x interface{}) string {
	switch x.(type) {
	case Post:
		return x.(Post).String()
	default:
		return "unknown"
	}
}