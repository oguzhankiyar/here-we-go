package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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
	Sample("Rel", Rel)
	Sample("Abs", Abs)
	Sample("Walk", Walk)
	Sample("WalkDir", WalkDir)
	Sample("VolumeName", VolumeName)
	Sample("FromSlash", FromSlash)
	Sample("ToSlash", ToSlash)
}

func Dir() {
	// Returns dir of path

	fn := func(p string) {
		dir := filepath.Dir(p)
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
		base := filepath.Base(p)
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
		ext := filepath.Ext(p)
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
		dir, file := filepath.Split(p)
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
		p := filepath.Join(items...)
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
		clean := filepath.Clean(p)
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
		matched, err := filepath.Match(pattern, name)
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
		isAbs := filepath.IsAbs(p)
		fmt.Printf("%q -> %t\n", p, isAbs)
	}

	fn("") // false
	fn("/") // true
	fn("c:/") // false
	fn("/dev/projects") // true
	fn("dev/projects") // false
}

func Rel() {
	// Returns relative to given path

	fn := func(base, p string) {
		rel, err := filepath.Rel(base, p)
		if err != nil {
			fmt.Printf("%q, %q -> %v\n", base, p, err)
		} else {
			fmt.Printf("%q, %q -> %q\n", base, p, rel)
		}
	}

	fn("/a", "/a/b/c")
	fn("/a", "/b/c")
	fn("/a", "./b/c")
}

func Abs() {
	// Returns absolute path

	fn := func(p string) {
		volume, err := filepath.Abs(p)
		if err != nil {
			fmt.Printf("%q -> %v\n", p, err)
		} else {
			fmt.Printf("%q -> %q\n", p, volume)
		}
	}

	fn("") // "C:\\dev\\projects\\here-we-go"
	fn("c:/dev") // "c:\\dev"
	fn("/dev") // "C:\\dev"
	fn("dev") // "C:\\dev\\projects\\here-we-go\\dev"
}

func Walk() {
	// Walks to given path

	prepareTestDirTree := func(tree string) (string, error) {
		tmpDir, err := os.MkdirTemp("", "")

		if err != nil {
			return "", fmt.Errorf("error creating temp directory: %v\n", err)
		}

		err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
		if err != nil {
			os.RemoveAll(tmpDir)
			return "", err
		}

		return tmpDir, nil
	}

	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}

	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}

func WalkDir() {
	// Walks to given path

	prepareTestDirTree := func(tree string) (string, error) {
		tmpDir, err := os.MkdirTemp("", "")

		if err != nil {
			return "", fmt.Errorf("error creating temp directory: %v\n", err)
		}

		err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
		if err != nil {
			os.RemoveAll(tmpDir)
			return "", err
		}

		return tmpDir, nil
	}

	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}

	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	err = filepath.WalkDir(".", func(path string, info os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}

func VolumeName() {
	// Returns volume of path

	fn := func(p string) {
		volume := filepath.VolumeName(p)
		fmt.Printf("%q -> %q\n", p, volume)
	}

	fn("") // ""
	fn("c:/dev") // "c:"
	fn("/dev") // ""
	fn("dev") // ""
}

func FromSlash() {
	// Converts slash to separator

	fn := func(p string) {
		volume := filepath.FromSlash(p)
		fmt.Printf("%q -> %q\n", p, volume)
	}

	fn("") // ""
	fn("c:/dev") // "c:\\dev"
	fn("/dev") // "\\dev"
	fn("dev") // "dev"
}

func ToSlash() {
	// Converts separator to slash

	fn := func(p string) {
		volume := filepath.ToSlash(p)
		fmt.Printf("%q -> %q\n", p, volume)
	}

	fn("") // ""
	fn("c:\\dev") // "c:/dev"
	fn("\\dev") // "/dev"
	fn("dev") // "dev"
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}