package main

import (
	"fmt"
	"reflect"
)

// 反射是一种机制，在编译时不知道具体类型的情况下，可以透视结构的组成、更新值

type Cat struct {
	Name string
}

func main() {
	var f float64 = 3.5
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())

	c := Cat{
		Name: "kitty",
	}
	t2 := reflect.TypeOf(c)
	fmt.Println(t2)
}
