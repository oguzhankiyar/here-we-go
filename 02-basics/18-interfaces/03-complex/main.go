package main

func main() {
	word := Word{
		Text: "This is word doc",
	}

	pdf := Pdf{
		Text: "This is pdf doc",
	}

	var list List

	list = append(list, word, pdf)

	list.Print()
}