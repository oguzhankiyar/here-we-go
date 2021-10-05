package main

import "fmt"

type Item interface {
	GetSize() int
}

type File struct {
	Name string
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

type Folder struct {
	Name 		string
	Children	[]Item
}

func (f Folder) GetSize() int {
	size := 0

	for _, child := range f.Children {
		size += child.GetSize()
	}

	return size
}

func main() {
	imgFile1 := &File{"image_1.jpg", 55}
	imgFile2 := &File{"image_2.jpg", 25}
	txtFile1 := &File{"todo.txt", 10}

	imgFolder := &Folder{"img", []Item{imgFile1, imgFile2}}
	txtFolder := &Folder{"txt", []Item{txtFile1}}
	rootFolder := &Folder{"assets", []Item{imgFolder, txtFolder}}

	fmt.Println("Size:", rootFolder.GetSize())
}