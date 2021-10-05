package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	Sample("Select", Select)
	Sample("Insert", Insert)
	Sample("Update", Update)
	Sample("Delete", Delete)
}

func Select() {
	sql, args, err := sq.
		Select("*").
		From("users").
		Join("emails USING (email_id)").
		Where(sq.Eq{"deleted_at": nil}).
		ToSql()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func Insert() {
	sql, args, err := sq.
		Insert("users").
		Columns("name", "age").
		Values("moe", 13).
		Values("larry", sq.Expr("? + 5", 12)).
		ToSql()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func Update() {
	sql, args, err := sq.
		Update("users").
		Set("name", "gopher").
		Where(sq.Eq{"id": 1}).
		ToSql()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func Delete() {
	sql, args, err := sq.
		Delete("users").
		Where(sq.Eq{"id": 1}).
		ToSql()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}