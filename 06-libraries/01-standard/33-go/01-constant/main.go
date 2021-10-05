package main

import (
	"fmt"
	"go/constant"
	"math"
)

func main() {
	i := constant.MakeInt64(math.MaxInt64)
	fmt.Printf("%v\n", constant.Val(i))

	e := constant.MakeFloat64(math.E)
	fmt.Printf("%v\n", constant.Val(e))

	b := constant.MakeBool(true)
	fmt.Printf("%v\n", constant.Val(b))

	b = constant.Make(false)
	fmt.Printf("%v\n", constant.Val(b))
}