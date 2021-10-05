package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/badger"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions(os.TempDir() + "badger.temp"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("key-1"), []byte("val-1"))
		if err != nil {
			return err
		}

		fmt.Println("set key-1 success")

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("key-1"))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			fmt.Println("key-1:", string(val))
			return nil
		})
	})
	if err != nil {
		log.Fatal(err)
	}
}