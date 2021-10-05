package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = os.TempDir() + "nuts.db"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Put("bck", []byte("my-key"), []byte("my-value"), 1000)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get("bck", []byte("my-key"))
		fmt.Printf("%s\n", entry.Value)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}