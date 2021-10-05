package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	u := uuid.New()
	fmt.Println("uuid:", u.String())
}