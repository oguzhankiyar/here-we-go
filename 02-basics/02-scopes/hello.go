package main

// file scope
import "fmt"

// package scope
func hello(name string) { // block scope start
	fmt.Println(helloMessage + messageSeparator, name + messageSuffix)
} // block scope end