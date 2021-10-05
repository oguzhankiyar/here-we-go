package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile(os.TempDir() + "level.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data, err := db.Get([]byte("key"), nil)
	fmt.Println(data)

	err = db.Put([]byte("my-key"), []byte("my-value"), nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		fmt.Printf("%s: %s\n", key, value)
	}
	iter.Release()
	err = iter.Error()

	err = db.Delete([]byte("key"), nil)
	fmt.Println(data)
}