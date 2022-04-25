package main

import (
	"fmt"
	"github.com/araddon/dateparse"
	"log"
)

// 处理时间
func main() {
	t1, err := dateparse.ParseAny("3/1/2014")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1.Format("2006-01-02 15:04:05"))

	t2, err := dateparse.ParseAny("mm/dd/yyyy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
}
