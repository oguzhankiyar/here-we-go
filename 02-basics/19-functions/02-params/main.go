package main

import "fmt"

func main() {
	do("Hello!", "Hi!", "Hey!")
}

func do(strings ...string) {
	for _, v := range strings {
		fmt.Println(v)
	}
}