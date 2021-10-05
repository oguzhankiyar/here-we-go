package main

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
)

type Person struct {
	Id 			string 			`json:"id,omitempty"`
	FirstName 	string 			`json:"firstname,omitempty"`
	LastName 	string 			`json:"lastname,omitempty"`
	Links		[]Link			`json:"social,omitempty"`
}

type Link struct {
	Title 		string 			`json:"title"`
	Address 	string			`json:"media"`
}

func main() {
	cluster, err := gocb.Connect("couchbase://127.0.0.1:9042", gocb.ClusterOptions{})
	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("example")
	collection := bucket.Collection("person")

	newPerson := Person{
		FirstName: "Gopher",
		LastName: "Go",
		Links: []Link{
			{"Web", "golang.org"},
		},
	}
	collection.Upsert("gopher", newPerson, &gocb.UpsertOptions{Expiry: 0})

	var findPerson Person
	result, err := collection.Get("gopher", &gocb.GetOptions{})
	if err != nil {
		panic(err)
	}

	err = result.Content(&findPerson)
	if err != nil {
		panic(err)
	}

	fmt.Println(findPerson)
}