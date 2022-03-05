package main

import (
	"fmt"
	"github.com/xujiajun/nutsdb"
	"log"
)

// nutsdb与我们之前介绍过的buntdb有些类似，但是支持List、Set、Sorted Set这些数据结构。
func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		val := []byte("dj")
		if err := tx.Put("", key, val, 0); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *nutsdb.Tx) error {
		key := []byte("name")
		if e, err := tx.Get("", key); err != nil {
			return err
		} else {
			fmt.Println(string(e.Value))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
