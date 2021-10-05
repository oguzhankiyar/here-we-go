package main

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func main() {
	Sample("Select", Select)
	Sample("SelectWhere", SelectWhere)
	Sample("InsertWithVals", InsertWithVals)
	Sample("InsertWithRecord", InsertWithRecord)
	Sample("InsertWithStruct", InsertWithStruct)
	Sample("InsertWithCols", InsertWithCols)
	Sample("UpdateWithRecord", UpdateWithRecord)
	Sample("UpdateWithStruct", UpdateWithStruct)
	Sample("UpdateWhere", UpdateWhere)
	Sample("Delete", Delete)
	Sample("DeleteWhere", DeleteWhere)
}

func Select() {
	sql, args, err := goqu.
		From("test").
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func SelectWhere() {
	sql, args, err := goqu.
		From("test").
		Where(goqu.Ex{
			"d": []string{"a", "b", "c"},
		}).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func InsertWithVals() {
	sql, args, err := goqu.Insert("user").
		Cols("first_name", "last_name").
		Vals(
			goqu.Vals{"Greg", "Farley"},
			goqu.Vals{"Jimmy", "Stewart"},
			goqu.Vals{"Jeff", "Jeffers"},
		).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func InsertWithRecord() {
	sql, args, err := goqu.
		Insert("user").
		Rows(
			goqu.Record{"first_name": "Greg", "last_name": "Farley"},
			goqu.Record{"first_name": "Jimmy", "last_name": "Stewart"},
			goqu.Record{"first_name": "Jeff", "last_name": "Jeffers"},
		).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func InsertWithStruct() {
	type User struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}

	sql, args, err := goqu.
		Insert("user").
		Rows(
			User{FirstName: "Greg", LastName: "Farley"},
			User{FirstName: "Jimmy", LastName: "Stewart"},
			User{FirstName: "Jeff", LastName: "Jeffers"},
		).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func InsertWithCols() {
	sql, args, err := goqu.
		Insert("user").
		Prepared(true).
		Cols("first_name", "last_name").
		FromQuery(goqu.
			From("other_table").
			Select("fn", "ln")).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func UpdateWithRecord() {
	sql, args, err := goqu.
		Update("items").
		Set(
			goqu.Record{"name": "Test", "address": "111 Test Addr"},
		).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func UpdateWithStruct() {
	type item struct {
		Address string `db:"address"`
		Name    string `db:"name" goqu:"skipupdate"`
	}

	sql, args, err := goqu.
		Update("items").
		Set(
			item{Name: "Test", Address: "111 Test Addr"},
		).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func UpdateWhere() {
	sql, args, err := goqu.
		Update("test").
		Set(goqu.Record{"foo": "bar"}).
		Where(goqu.Ex{
			"a": goqu.Op{"gt": 10},
		}).
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func Delete() {
	sql, args, err := goqu.
		Delete("items").
		ToSQL()

	if err != nil {
		panic(err)
	}

	fmt.Println("sql:", sql)
	fmt.Println("args:", args)
}

func DeleteWhere() {
	sql, args, err := goqu.
		Delete("items").
		Where(
			goqu.Ex{"c": nil},
		).
		ToSQL()

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