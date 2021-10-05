package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// We can use third party packages that other developers created

	u := uuid.New()
	fmt.Println("uuid:", u.String())
}