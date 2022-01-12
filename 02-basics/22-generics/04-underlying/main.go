package main

import "fmt"

func main() {
	type OtherInt64Type int64

	var intValue int = 1
	var int32Value int32 = 2
	var int64Value int64 = 3
	var otherInt64Value OtherInt64Type = 4

	// This is not acceptable because it is int
	//WithoutLike(intValue)

	// This is ok
	WithoutLike(int32Value)

	// This is ok
	WithoutLike(int64Value)

	// This is not acceptable because it is underlying int type
	//WithoutLike(otherInt64Value)

	// This is ok
	WithLike(int32Value)

	// This is ok
	WithLike(int64Value)

	// This is ok
	WithLike(otherInt64Value)

	_ = intValue
}

type IntType interface {
	int32 | int64
}

type IntLikeType interface {
	~int32 | ~int64
}

func WithoutLike[T IntType](v T) {
	fmt.Println(v)
}

func WithLike[T IntLikeType](v T) {
	fmt.Println(v)
}