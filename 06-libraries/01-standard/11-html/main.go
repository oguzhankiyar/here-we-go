package main

import (
	"fmt"
	"html"
)

func main() {
	Sample("EscapeString", EscapeString)
	Sample("UnescapeString", UnescapeString)
}

func EscapeString() {
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	fmt.Println(html.EscapeString(s))
}

func UnescapeString() {
	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	fmt.Println(html.UnescapeString(s))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}