package main

import (
	"fmt"
	"log"
	"os"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open(os.TempDir() + "my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return err
		}

		b := tx.Bucket([]byte("MyBucket"))
		err = b.Put([]byte("answer"), []byte("42"))

		return err
	})

	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)

		return nil
	})
}