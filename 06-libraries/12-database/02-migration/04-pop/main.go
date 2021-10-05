package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/soda/cmd"
	"pop-sample/models"
)

func main() {
	cmd.RootCmd.Use = "soda"
	cmd.Execute()

	fmt.Println("This is a Go database test program")

	db, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}

	user := models.User{Firstname: "Vincent", Lastname: "Vega"}
	_, err = db.ValidateAndSave(&user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Added %s to database\n", user.Firstname)

	user = models.User{Firstname: "Jules", Lastname: "Winnfield"}
	_, err = db.ValidateAndSave(&user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Added %s to database\n", user.Firstname)
}
