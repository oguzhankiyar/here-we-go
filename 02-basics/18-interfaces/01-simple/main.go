package main

import "fmt"

type Writer interface {
	Write(str string)
}

type SomeWriter struct {

}

func (writer SomeWriter) Write(str string) {
	fmt.Println(str)
}

func Write(w Writer, str string)  {
	w.Write(str)
}

func main() {
	w := &SomeWriter{}

	Write(w, "Hey!")
}