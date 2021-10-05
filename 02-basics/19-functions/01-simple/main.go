package main

import "fmt"

func main() {
	doWithoutParam()
	doWithParam("Hi!")
	fmt.Println(doWithReturnWithoutParam())
	fmt.Println(doWithReturnWithParam("Hi!"))
}

func doWithoutParam() {
	fmt.Println("Hey!")
}

func doWithParam(str string) {
	fmt.Println(str)
}

func doWithReturnWithoutParam() string {
	return "Hey"
}

func doWithReturnWithParam(str string) string {
	return str
}