package main

import (
	"fmt"

	"github.com/schollz/peerdiscovery"
)

func main() {
	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{Limit: 1})
	for _, d := range discoveries {
		fmt.Printf("discovered '%s'\n", d.Address)
	}
}