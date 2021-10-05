package main

import (
	"log"
	"os"

	"github.com/google/wire"
)

type MongoDataStore struct{}

type DataStore interface{}

type PersonResource struct {
	Store  *DataStore
	Logger *log.Logger
}

func NewPersonResource(store *DataStore, logger *log.Logger) *PersonResource {
	return &PersonResource{Store: store, Logger: logger}
}

func main() {
	e := InitializePersonResource()
	e.createItem("")
}

func InitializePersonResource() *PersonResource {
	wire.Build(NewDatastore, NewLogger, NewPersonResource)
	return &PersonResource{}
}

func (pr *PersonResource) createItem(string string) {
}

func NewLogger() *log.Logger {
	return log.New(os.Stderr, "exp", log.LstdFlags|log.Lshortfile)
}

func NewDatastore() *DataStore {
	var datastore DataStore = &MongoDataStore{}
	return &datastore
}