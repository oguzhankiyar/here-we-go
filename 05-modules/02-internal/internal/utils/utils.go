package utils

import (
	"math/rand"
	"time"
)

func Rand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}