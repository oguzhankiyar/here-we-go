package main

import "fmt"

func main() {
	fn := func(str string, separator rune) {
		fmt.Printf("%q with %q -> %q\n", str, separator, Split(str, separator))
	}

	fn("this.is.gopher!", '.')
	fn("my name is gopher!", ' ')
}

func Split(str string, separator rune) []string {
	items := make([]string, 0)

	temp := make([]rune, 0)

	for _, c := range str {
		if c ==  separator {
			items = append(items, string(temp))
			temp = temp[:0]
		} else {
			temp = append(temp, c)
		}
	}

	if len(temp) > 0 {
		items = append(items, string(temp))
		temp = temp[:0]
	}

	return items
}