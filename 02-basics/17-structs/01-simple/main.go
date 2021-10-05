package main

import "fmt"

func main() {
	type Post struct {
		Title 		string
		Category 	string
	}

	post := Post{
		Title: "Go",
		Category: "Programming Language",
	}

	fmt.Printf("post: %+v\n", post)
	fmt.Printf("post.Title: %v\n", post.Title)
	fmt.Printf("post.Category: %v\n", post.Category)
}