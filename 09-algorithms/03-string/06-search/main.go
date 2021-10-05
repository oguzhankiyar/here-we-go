package main

import "fmt"

func main() {
	fn := func(str string, substr string) {
		start, end := Search(str, substr)
		fmt.Printf("%q in %q -> [%v:%v]\n", substr, str, start, end)
	}

	fn("this is gopher", "go")
	fn("this is golang", "gopher")
}

func Search(str string, substr string) (int, int) {
	for i := 0; i < len(str) - len(substr) + 1; i++ {
		if str[i:i + len(substr)] == substr {
			return i, i + len(substr)
		}
	}

	return -1, -1
}