package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	Sample("Drivers", Drivers)
	Sample("Open", Open)
	Sample("Ping", Ping)
	Sample("Create", Create)
	Sample("Insert", Insert)
	Sample("Prepare", Prepare)
	Sample("Transaction", Transaction)
	Sample("Select", Select)
	Sample("Drop", Drop)
	Sample("Close", Close)
}

func Drivers() {
	drivers := sql.Drivers()
	fmt.Println(drivers)
}

func Open() {
	file, err := os.CreateTemp("", "*.db")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	db, err = sql.Open("sqlite3", file.Name())
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Open Success!")
}

func Ping() {
	err := db.Ping()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Ping Success!")
}

func Create() {
	_, err := db.Exec(`create table "users" ("id" integer primary key, "name" varchar(64) null)`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Create Success!")
}

func Insert() {
	_, err := db.Exec(`insert into "users" ("id", "name") values (1, "Alice"), (2, "Bob")`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Insert Success!")
}

func Prepare() {
	stmt, err := db.Prepare(`insert into "users" ("id", "name") values (?, ?)`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	users := map[int]string{
		4: "Alex",
		5: "Martin",
	}

	for id, name := range users {
		_, err := stmt.Exec(id, name)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
	}

	stmt.Close()
}

func Transaction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(`insert into "users" ("id", "name") values (?, ?)`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer stmt.Close()

	users := map[int]string{
		7: "Steve",
		8: "Bill",
	}

	for id, name := range users {
		_, err = stmt.Exec(id, name)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func Select() {
	rows, err := db.Query(`select "id", "name" from "users"`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer rows.Close()

	var (
		id int
		name string
	)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Printf("%v - %s\n", id, name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func Drop() {
	_, err := db.Exec(`drop table "users"`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Drop Success!")
}

func Close() {
	err := db.Close()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Close Success!")
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}