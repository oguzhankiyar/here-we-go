package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	Sample("Plain", Plain)
	Sample("Object", Object)
	Sample("Loop", Loop)
}

func Plain() {
	t, err := template.New("Tmpl1").Parse(`Hello, {{.}}!`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = t.Execute(os.Stdout, "Gopher")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println()
}

func Object() {
	type Person struct {
		Name	string
		Age		uint
	}
	data := Person{"Gopher", 5}

	tmpl, err := template.New("test").Parse("{{.Name}} is {{.Age}} years old")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println()
}

func Loop() {
	type Person struct {
		Name	string
		Age		uint
	}
	data := []Person{
		{"Gopher", 5},
		{"Alice", 10},
		{"Bob", 11},
	}

	tmpl, err := template.New("test").Parse("{{range $val := .}} - {{$val.Name}}\n{{end}}")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println()
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}