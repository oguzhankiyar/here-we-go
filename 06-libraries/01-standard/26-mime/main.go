package main

import (
	"fmt"
	"mime"
)

func main() {
	Sample("AddExtensionType", AddExtensionType)
	Sample("FormatMediaType", FormatMediaType)
	Sample("ParseMediaType", ParseMediaType)
	Sample("ExtensionByType", ExtensionByType)
	Sample("TypeByExtension", TypeByExtension)
}

func AddExtensionType() {
	err := mime.AddExtensionType(".go", "app/go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func FormatMediaType() {
	mediatype := "text/html"
	params := map[string]string{
		"charset": "utf-8",
	}

	result := mime.FormatMediaType(mediatype, params)
	fmt.Println("result:", result)
}

func ParseMediaType() {
	mediatype, params, err := mime.ParseMediaType("text/html; charset=utf-8")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("MediaType:", mediatype)
	fmt.Println("Params:", params)
}

func ExtensionByType() {
	fn := func(typ string) {
		ext, err := mime.ExtensionsByType(typ)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Printf("%s -> %s\n", typ, ext)
	}

	fn("application/json")
	fn("text/html")
	fn("app/go")
	fn("text/gopher")
}

func TypeByExtension() {
	fn := func(ext string) {
		typ := mime.TypeByExtension(ext)
		fmt.Printf("%s -> %s\n", ext, typ)
	}

	fn(".json")
	fn(".html")
	fn(".go")
	fn(".gopher")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}