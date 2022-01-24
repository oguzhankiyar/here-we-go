package main

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type Database struct {
	items map[int]string
}

func NewDatabase() *Database {
	return &Database{
		items: make(map[int]string),
	}
}

func (d *Database) Connect() {
	println("Connecting database")
}

func (d *Database) Disconnect() {
	println("Disconnecting database")
}

func (d *Database) GetAll() map[int]string {
	return d.items
}

func (d *Database) GetOne(id int) string {
	return d.items[id]
}

func (d *Database) Has(id int) bool {
	_, ok := d.items[id]
	return ok
}

func (d *Database) Set(id int, value string) {
	d.items[id] = value
}

func (d *Database) Del(id int) {
	delete(d.items, id)
}

func (d *Database) Clear() {
	d.items = make(map[int]string)
}

func HashItem(database *Database, id int) (string, error) {
	if !database.Has(id) {
		return "", errors.New("the item could not found")
	}

	value := database.GetOne(id)

	hash := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v%s", id, value)))[0:10]

	return hash, nil
}
