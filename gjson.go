package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

// Go 每日一库之 gjson
// gjson实际上是get + json的缩写，用于读取 JSON 串，同样的还有一个sjson（set + json）库用来设置 JSON 串
func main() {
	json := `{"name":{"first":"li","last":"dj"},"age":18}`
	lastName := gjson.Get(json, "name.last") // 返回gjson.Result类型
	fmt.Println("last name:", lastName.String())

	age := gjson.Get(json, "age")
	fmt.Println("age:", age.Int())
}
