package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	booksAPI := app.Party("/books")

	booksAPI.Use(iris.Compression)

	booksAPI.Get("/", list)
	booksAPI.Post("/", create)

	app.Listen(":1234")
}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)

	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		return
	}

	fmt.Println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}