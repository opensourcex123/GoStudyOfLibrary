package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
)

// Go 每日一库之 buntdb 用 Go 语言编写的内存键值数据库,支持 ACID、并发读、自定义索引和空间信息数据
func main() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *buntdb.Tx) error {
		oldValues, replaced, err := tx.Set("testKey", "testValue", nil)
		if err != nil {
			return err
		}

		fmt.Printf("old value:%q replaced:%t\n", oldValues, replaced)
		return nil
	})

	db.View(func(tx *buntdb.Tx) error {
		value, err := tx.Get("testKey")
		if err != nil {
			return err
		}

		fmt.Println("value is:", value)
		return nil
	})
}
