package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

func main() {
	Sample("Marshal", Marshal)
	Sample("MarshalWithTags", MarshalWithTags)
	Sample("Unmarshal", Unmarshal)
	Sample("UnmarshalWithTags", UnmarshalWithTags)
}

func Marshal() {
	type Post struct {
		Title string
	}

	post := Post{
		Title: "Post 1",
	}

	data, err := xml.Marshal(&post)

	fmt.Printf("%q %v\n", string(data), err) // "<Post><Title>Post 1</Title></Post>" <nil>
}

func MarshalWithTags() {
	type Post struct {
		Title string    `xml:"Title"`
		Date  time.Time `xml:"CreatedAt"`
	}

	post := Post{
		Title: "Post 1",
		Date:  time.Date(2021, 01, 01, 01, 00, 00, 00, time.UTC),
	}

	data, err := xml.Marshal(&post)

	fmt.Printf("%q %v\n", string(data), err) // "<Post><Title>Post 1</Title><CreatedAt>2021-01-01T01:00:00Z</CreatedAt></Post>" <nil>
}

func Unmarshal() {
	xmlString := `<Post><Title>Post 1</Title></Post>`

	type Post struct {
		Title string
	}

	var post Post

	err := xml.Unmarshal([]byte(xmlString), &post)

	fmt.Printf("%#v %v\n", post, err) // main.go.Post{Title:"Post 1"} <nil>
}

func UnmarshalWithTags() {
	xmlString := `<Post><Title>Post 1</Title><CreatedAt>2021-01-01T01:00:00Z</CreatedAt></Post>`

	type Post struct {
		Title string    `xml:"Title"`
		Date  time.Time `xml:"CreatedAt"`
	}

	var post Post

	_ = xml.Unmarshal([]byte(xmlString), &post)

	fmt.Printf("title: %q date: %s\n", post.Title, post.Date) // title: "Post 1" date: 2021-01-01 01:00:00 +0000 UTC
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
