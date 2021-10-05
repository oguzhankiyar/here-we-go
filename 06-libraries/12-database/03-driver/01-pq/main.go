package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:123456@localhost:5432/here_we_go?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE pq_sample (Id int, Name text)")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO pq_sample VALUES (1, 'GOPHER')")
	if err != nil {
		panic(err)
	}
	defer func() {
		_, err = db.Exec("DROP TABLE pq_sample")
		if err != nil {
			panic(err)
		}
	}()

	rows, err := db.Query("SELECT Id, Name FROM pq_sample WHERE Id = $1", 1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username string

		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v - %s\n", id, username)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}