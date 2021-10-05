package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	Sample("First", First)
	Sample("Second", Second)
}

func First() {
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 1, '-', tabwriter.Debug)
	w.Write([]byte("\tHello, there!\n\t\tMy name is Gopher."))
	w.Flush()

	fmt.Println()
}

func Second() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}