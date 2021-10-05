package main

import (
	"fmt"

	"github.com/joomcode/errorx"
)

func main() {
	err := someFunc()
	fmt.Println(err.Error())

	err = errorx.Decorate(err, "decorate")
	fmt.Println(err.Error())

	err = errorx.Decorate(err, "outer decorate")
	fmt.Println(err.Error())
}

func someFunc() error {
	return errorx.AssertionFailed.New("example")
}