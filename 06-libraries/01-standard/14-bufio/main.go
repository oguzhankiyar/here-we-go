package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Sample("ScannerFirst", ScannerFirst)
	Sample("ScannerSecond", ScannerSecond)
	Sample("ScannerThird", ScannerThird)
}

func ScannerFirst() {
	scanner := bufio.NewScanner(strings.NewReader("I am Gopher!"))

	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 12) // true
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}

func ScannerSecond() {
	scanner := bufio.NewScanner(strings.NewReader("I am Gopher!"))

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err) // I am Gopher!
	}
}

func ScannerThird() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}

		if !atEOF {
			return 0, nil, nil
		}

		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err) // "1" "2" "3" "4" ""
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
