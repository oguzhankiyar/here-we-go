package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"ent-sample/ent"
	"ent-sample/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

var client *ent.Client

func main() {
	var err error
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	Sample("Schema", Schema)
	Sample("Create", Create)
	Sample("Query", Query)
	Sample("Order", Order)
	Sample("Limit", Limit)
	Sample("Offset", Offset)
	Sample("Update", Update)
	Sample("Delete", Delete)
}

func Schema() {
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("failed creating schema resources: %v\n", err)
		return
	}

	fmt.Println("schema was created")
}

func Create() {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(u)
}

func Query() {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		Only(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(u)
}

func Order() {
	users, err := client.User.Query().
		Order(ent.Asc(user.FieldName)).
		All(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(users)
}

func Limit() {
	users, err := client.User.
		Query().
		Limit(10).
		All(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(users)
}

func Offset() {
	users, err := client.User.
		Query().
		Offset(10).
		All(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(users)
}

func Update() {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		Only(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	a8m, err := u.Update().
		SetAge(30).
		Save(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(a8m)
}

func Delete() {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		Only(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = client.User.
		DeleteOne(u).
		Exec(context.Background())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("deleted")
}

func PrintJson(data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Printf("%s\n", j)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}