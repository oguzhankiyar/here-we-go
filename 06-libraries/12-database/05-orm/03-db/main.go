package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

var session db.Session

type Book struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	AuthorID    uint   `db:"author_id"`
	SubjectID   uint   `db:"subject_id"`
}

func main() {
	var err error
	path := fmt.Sprintf("%s%v.db", os.TempDir(), time.Now().Unix())
	session, err = sqlite.Open(sqlite.ConnectionURL{Database: path})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SQL().Exec(`
CREATE TABLE "books" (
  "id" INTEGER NOT NULL,
  "title" VARCHAR NOT NULL,
  "author_id" INTEGER,
  "subject_id" INTEGER,
  CONSTRAINT "books_id_pkey" PRIMARY KEY ("id")
);
`)

	fmt.Printf("Connected to %q:\n", session.Name())

	Sample("Collections", Collections)
	Sample("Collection", Collection)
	Sample("Insert", Insert)
	Sample("InsertWithSQL", InsertWithSQL)
	Sample("Next", Next)
	Sample("Find", Find)
	Sample("Cond", Cond)
	Sample("Count", Count)
	Sample("OrderBy", OrderBy)
	Sample("Paginate", Paginate)
	Sample("Update", Update)
	Sample("UpdateWithSQL", UpdateWithSQL)
	Sample("Delete", Delete)
	Sample("DeleteWithSQL", DeleteWithSQL)
}

func Collections() {
	collections, err := session.Collections()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if len(collections) == 0 {
		fmt.Println("no collection")
		return
	}

	for i := range collections {
		fmt.Printf("-> %s\n", collections[i].Name())
	}
}

func Collection() {
	collection := session.Collection("books")
	exists, err := collection.Exists()
	if !exists && errors.Is(err, db.ErrCollectionDoesNotExist) {
		fmt.Printf("Collection does not exist: %v\n", err)
		return
	}

	fmt.Println(collection.Name())
}

func Insert() {
	collection := session.Collection("books")

	book1 := Book{
		ID: 1,
		Title: "book-1",
		AuthorID: 1,
		SubjectID: 1,
	}

	book2 := Book{
		ID: 2,
		Title: "book-2",
		AuthorID: 2,
		SubjectID: 1,
	}

	book3 := Book{
		ID: 3,
		Title: "book-3",
		AuthorID: 3,
		SubjectID: 2,
	}

	res, err := collection.Insert(&book1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res, err = collection.Insert(&book2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res, err = collection.Insert(&book3)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("inserted with id:", res.ID())
}

func InsertWithSQL() {
	book := Book{
		Title:    "The Gopher",
		AuthorID: 1,
	}

	res, err := session.SQL().
		InsertInto("books").
		Values(book).
		Exec()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	id, _ := res.LastInsertId()

	fmt.Println("inserted with id:", id)
}

func Next() {
	collection := session.Collection("books")

	res := collection.Find()
	defer res.Close()

	var book Book

	for res.Next(&book) {
		PrintJson(book)
	}

	if err := res.Err(); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("completed")
}

func Find() {
	collection := session.Collection("books")

	var books []Book
	err := collection.Find().All(&books)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(books)
}

func Cond() {
	collection := session.Collection("books")

	var book Book
	err := collection.Find(db.Cond{"id": 1}).One(&book)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(book)
}

func Count() {
	collection := session.Collection("books")

	total, err := collection.Find().Count()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("total:", total)
}

func OrderBy() {
	collection := session.Collection("books")

	var books []Book
	err := collection.Find().OrderBy("-title").All(&books)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(books)
}

func Paginate() {
	collection := session.Collection("books")

	res := collection.
		Find().
		Paginate(2).
		Cursor("id")

	var books []Book
	err := res.All(&books)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(books)

	res = res.NextPage(books[len(books)-1].ID)

	err = res.All(&books)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(books)
}

func Update() {
	collection := session.Collection("books")

	var book Book
	res := collection.Find(1)
	err := res.One(&book)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	book.Title = "New title"

	err = res.Update(book)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	PrintJson(book)
}

func UpdateWithSQL() {
	q := session.SQL().
		Update("books").
		Set("Title = ?", "The Golang").
		Where("id = ?", 1)

	res, err := q.Exec()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	count, _ := res.RowsAffected()
	fmt.Println("affected:", count)
}

func Delete() {
	collection := session.Collection("books")

	res := collection.Find(3)

	err := res.Delete()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("deleted")
}

func DeleteWithSQL() {
	q := session.SQL().
		DeleteFrom("books").
		Where("title", "The Gopher")

	res, err := q.Exec()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	count, _ := res.RowsAffected()
	fmt.Println("affected:", count)
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