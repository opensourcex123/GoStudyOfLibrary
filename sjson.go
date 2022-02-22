package main

import (
	"fmt"
	"github.com/tidwall/sjson"
)

// Go 每日一库之 sjson sjson快速设置 JSON 串中的值
// 同gjson一样，sjson同样不会检查传入的 JSON 串的合法性,如果传入一个非法的 JSON 串，程序输出不确定的结果

const json = `{"name":{"first":"li","last":"dj"},"age":18}`

func main() {
	value, _ := sjson.Set(json, "name.last", "nfx")
	fmt.Println(value)
}
