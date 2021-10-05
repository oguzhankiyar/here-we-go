package main

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
)

func main() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set("my-key", "my-value", nil)
		return err
	})

	err = db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			fmt.Printf("key: %s, value: %s\n", key, value)
			return true
		})
		return err
	})
}