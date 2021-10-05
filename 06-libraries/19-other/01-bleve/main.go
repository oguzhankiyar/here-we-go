package main

import (
	"fmt"
	"os"
	"time"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	// open a new index
	path := fmt.Sprintf("%sexample-%v.bleve", os.TempDir(), time.Now().Unix())
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(path, mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	type Item struct {
		Id		string
		Name 	string
	}

	items := []Item{
		{
			Id: "1",
			Name: "the golang language",
		},
		{
			Id: "2",
			Name: "the gopher",
		},
	}

	// index some data
	for _, item := range items {
		index.Index(item.Id, item)
	}

	// search for some text
	query := bleve.NewRegexpQuery("go.*")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}