package main

import (
	"fmt"
	"log"

	rethinkdb "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func main() {
	session, err := rethinkdb.Connect(rethinkdb.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := rethinkdb.Expr("Hello World").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
}