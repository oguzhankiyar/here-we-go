package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	Sample("Encode", Encode)
	Sample("Decode", Decode)
	Sample("Valid", Valid)
	Sample("Marshal", Marshal)
	Sample("MarshalWithTags", MarshalWithTags)
	Sample("Unmarshal", Unmarshal)
	Sample("UnmarshalWithTags", UnmarshalWithTags)
	Sample("Compact", Compact)
	Sample("Indent", Indent)
	Sample("HtmlEscape", HtmlEscape)
}

func Encode() {
	type Message struct {
		Author, Text string
	}

	enc := json.NewEncoder(os.Stdout)

	message := Message{
		Author: "Gopher",
		Text:   "I am Gopher!",
	}

	enc.Encode(&message) // {"Author":"Gopher","Text":"I am Gopher!"}
}

func Decode() {
	const jsonStream = `
	{"Author": "Alice", "Text": "Hi, I am Alice."}
	{"Author": "Bob", "Text": "Hi, I am Bob."}
	{"Author": "Alice", "Text": "Nice to meet you, Bob!"}
`
	type Message struct {
		Author, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var m Message

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s: %s\n", m.Author, m.Text)
	}
}

func Valid() {
	fn := func(data string) {
		bytes := []byte(data)
		valid := json.Valid(bytes)
		fmt.Printf("Valid(%q) -> %t\n", data, valid)
	}

	fn(`{ "Text": "Hello" }`)   // true
	fn(`"Text": "Hello" }`)     // false
	fn(`{ }`)                   // true
	fn(`[{ }]`)                 // true
	fn(`[{ "Text": "Hello" }]`) // true
}

func Marshal() {
	type Post struct {
		Title string
	}

	post := Post{
		Title: "Post 1",
	}

	jsonBytes, err := json.Marshal(&post)
	fmt.Printf("%s %v\n", jsonBytes, err) // {"Title":"Post 1"} <nil>
}

func MarshalWithTags() {
	type Post struct {
		Title string    `json:"title"`
		Date  time.Time `json:"created_at"`
	}

	post := Post{
		Title: "Post 1",
		Date:  time.Date(2021, 01, 01, 01, 00, 00, 00, time.UTC),
	}

	jsonBytes, err := json.Marshal(&post)
	fmt.Printf("%s %v\n", jsonBytes, err) // {"title":"Post 1","created_at":"2021-01-01T01:00:00Z"} <nil>
}

func Unmarshal() {
	jsonString := `{ "title": "Golang" }`

	type Post struct {
		Title string
	}

	var post Post

	json.Unmarshal([]byte(jsonString), &post)
	fmt.Printf("%#v\n", post) // main.go.Post{Title:"Golang"}
}

func UnmarshalWithTags() {
	jsonString := `{ "title": "Golang", "created_at": "2021-01-01T01:00:00Z" }`

	type Post struct {
		Title string    `json:"title"`
		Date  time.Time `json:"created_at"`
	}

	var post Post

	json.Unmarshal([]byte(jsonString), &post)
	fmt.Printf("title: %q date: %s\n", post.Title, post.Date) // title: "Golang" date: 2021-01-01 01:00:00 +0000 UTC
}

func Compact() {
	jsonString := `
{
	"Title": "Post 1"
}
`

	var dest bytes.Buffer

	_ = json.Compact(&dest, []byte(jsonString))

	fmt.Println(dest.String())
}

func Indent() {
	jsonString := `
{ "Title": "Post 1" }
`

	var dest bytes.Buffer

	_ = json.Indent(&dest, []byte(jsonString), "", "\t")

	fmt.Println(dest.String())
	/*
		{
			"Title": "Post 1"
		}
	*/
}

func HtmlEscape() {
	jsonString := `{ "Body": "<div>Post 1</div>" }`

	var dest bytes.Buffer

	json.HTMLEscape(&dest, []byte(jsonString))

	fmt.Println(dest.String()) // { "Body": "\u003cdiv\u003ePost 1\u003c/div\u003e" }
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
