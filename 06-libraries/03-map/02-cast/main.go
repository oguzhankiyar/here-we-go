package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	Sample("ToString", ToString)
	Sample("ToInt", ToInt)
	Sample("ToFloat", ToFloat)
	Sample("ToBool", ToBool)
	Sample("ToTime", ToTime)
	Sample("ToDuration", ToDuration)
	Sample("ToSlice", ToSlice)
}

func ToString() {
	fn := func(input interface{}) {
		output := cast.ToString(input)
		fmt.Printf("%#v -> %q\n", input, output)
	}

	fn("gopher")
	fn(5)
	fn(10.21)
	fn([]byte("golang"))
	fn(nil)

	var some interface{} = "this is go"
	fn(some)

	fmt.Println()
}

func ToInt() {
	fn := func(input interface{}) {
		output := cast.ToInt(input)
		fmt.Printf("%#v -> %v\n", input, output)
	}

	fn(8)
	fn(8.31)
	fn("8")
	fn(true)
	fn(false)
	fn(nil)

	var some interface{} = 5
	fn(some)
}

func ToFloat() {
	fn := func(input interface{}) {
		output := cast.ToFloat64(input)
		fmt.Printf("%#v -> %v\n", input, output)
	}

	fn(8)
	fn(8.31)
	fn("8")
	fn(true)
	fn(false)
	fn(nil)

	var some interface{} = 5
	fn(some)
}

func ToBool() {
	fn := func(input interface{}) {
		output := cast.ToBool(input)
		fmt.Printf("%#v -> %v\n", input, output)
	}

	fn(8)
	fn("1")
	fn(true)
	fn(false)
	fn(nil)

	var some interface{} = 5
	fn(some)
}

func ToTime() {
	fn := func(input interface{}) {
		output := cast.ToTime(input)
		fmt.Printf("%#v -> %v\n", input, output)
	}

	fn(60)
	fn("2021-01-01")
	fn(true)
	fn(nil)

	var some interface{} = "2021-01-01T01:00:00Z"
	fn(some)
}

func ToDuration() {
	fn := func(input interface{}) {
		output := cast.ToDuration(input)
		fmt.Printf("%#v -> %v\n", input, output)
	}

	fn(60)
	fn("1250")
	fn("10h")
	fn(true)
	fn(nil)

	var some interface{} = "5m"
	fn(some)
}

func ToSlice() {
	fn := func(input interface{}) {
		output := cast.ToStringSlice(input)
		fmt.Printf("%#v -> %#v\n", input, output)
	}

	fn([]int{1, 2})
	fn(1250)
	fn("gopher")
	fn(true)
	fn(nil)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}