package main

import (
	"fmt"
	"path"
)

func main() {
	Sample("Dir", Dir)
	Sample("Base", Base)
	Sample("Ext", Ext)
	Sample("Split", Split)
	Sample("Join", Join)
	Sample("Clean", Clean)
	Sample("Match", Match)
	Sample("IsAbs", IsAbs)
}

func Dir() {
	// Returns dir of path

	fn := func(p string) {
		dir := path.Dir(p)
		fmt.Printf("%q -> %q\n", p, dir)
	}

	fn("") // "."
	fn("C:/dev/projects/here-we-go") // "C:/dev/projects"
	fn("C:/dev/projects/here-we-go/") // "C:/dev/projects/here-we-go"
	fn("C:/dev/projects/here-we-go/main.go.go") // "C:/dev/projects/here-we-go"
}

func Base() {
	// Returns last element of path

	fn := func(p string) {
		base := path.Base(p)
		fmt.Printf("%q -> %q\n", p, base)
	}

	fn("") // "."
	fn("C:/dev/projects/here-we-go") // "here-we-go"
	fn("C:/dev/projects/here-we-go/") // "here-we-go"
	fn("C:/dev/projects/here-we-go/main.go.go") // "main.go.go"
}

func Ext() {
	// Returns ext of path

	fn := func(p string) {
		ext := path.Ext(p)
		fmt.Printf("%q -> %q\n", p, ext)
	}

	fn("") // ""
	fn("C:/dev/projects/here-we-go") // ""
	fn("C:/dev/projects/here-we-go/") // ""
	fn("C:/dev/projects/here-we-go/main.go.go") // ".go"
}

func Split() {
	// Returns split of path

	fn := func(p string) {
		dir, file := path.Split(p)
		fmt.Printf("%q -> %q, %q\n", p, dir, file)
	}

	fn("") // "", ""
	fn("C:/dev/projects/here-we-go") // "C:/dev/projects/", "here-we-go"
	fn("C:/dev/projects/here-we-go/") // "C:/dev/projects/here-we-go/", ""
	fn("C:/dev/projects/here-we-go/main.go.go") // "C:/dev/projects/here-we-go/", "main.go.go"
}

func Join() {
	// Joins paths

	fn := func(items ...string) {
		p := path.Join(items...)
		fmt.Printf("%q -> %q\n", items, p)
	}

	fn("") // ""
	fn("", "dev") // "dev"
	fn("dev", "") // "dev"
	fn("c:/", "dev") // "c:/dev"
	fn("c:/", "dev", "/projects") // "c:/dev/projects"
	fn("c:/", "dev", "/projects/here-we-go/") // "c:/dev/projects/here-we-go"
	fn("/", "dev", "projects", "here-we-go", "main.go.go") // "/dev/projects/here-we-go/main.go.go"
	fn("dev", "projects", "../") // "dev"
	fn("dev", "projects", "./") // "dev/projects"
}

func Clean() {
	// Joins paths

	fn := func(p string) {
		clean := path.Clean(p)
		fmt.Printf("%q -> %q\n", p, clean)
	}

	fn("") // "."
	fn("/") // "/"
	fn("/dev") // "/dev"
	fn("//dev") // "/dev"
	fn("dev/") // "dev"
	fn("dev//") // "dev"
	fn("dev//projects") // "dev/projects"
	fn("/dev/../dev/projects") // "/dev/projects"
	fn("/dev/projects/..") // "/dev"
	fn("/dev/projects/main.go..go") // "/dev/projects/main.go..go"
}

func Match() {
	// Matches path with pattern

	fn := func(pattern, name string) {
		matched, err := path.Match(pattern, name)
		if err != nil {
			fmt.Printf("%q, %q -> %v\n", name, pattern, err)
		} else {
			fmt.Printf("%q, %q -> %t\n", name, pattern, matched)
		}
	}

	fn("abc", "abc")
	fn("a*", "abc")
	fn("a*/b", "a/c/b")
}

func IsAbs() {
	// Checks path is absolute

	fn := func(p string) {
		isAbs := path.IsAbs(p)
		fmt.Printf("%q -> %t\n", p, isAbs)
	}

	fn("") // false
	fn("/") // true
	fn("c:/") // false
	fn("/dev/projects") // true
	fn("dev/projects") // false
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}