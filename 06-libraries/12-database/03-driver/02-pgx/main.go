package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	connStr := "postgres://postgres:123456@localhost:5432/here_we_go"
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	ctx := context.Background()

	_, err = conn.Exec(ctx, "CREATE TABLE pgx_sample (Id int, Name text)")
	if err != nil {
		panic(err)
	}
	_, err = conn.Exec(ctx, "INSERT INTO pgx_sample VALUES (1, 'GOPHER')")
	if err != nil {
		panic(err)
	}
	defer func() {
		_, err = conn.Exec(ctx, "DROP TABLE pgx_sample")
		if err != nil {
			panic(err)
		}
	}()

	var id int
	var name string
	err = conn.QueryRow(context.Background(), "SELECT Id, Name FROM pgx_sample WHERE Id = $1", 1).Scan(&id, &name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, name)
}