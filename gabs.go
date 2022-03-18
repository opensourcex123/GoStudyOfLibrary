package main

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

// gabs是一个用来查询和修改 JSON 串的库。
// 它使用encoding/json将一般的 JSON 串转为map[string]interface{}，并提供便利的方法操作map[string]struct{}。
func main() {
	jobj, _ := gabs.ParseJSON([]byte(`{
    "info": {
      "name": {
        "first": "lee",
        "last": "darjun"
      },
      "age": 18,
      "hobbies": [
        "game",
        "programming"
      ]
    }
    }`))

	fmt.Println("first name: ", jobj.Search("info", "name", "first").Data().(string))
	fmt.Println("last name: ", jobj.Path("info.name.last").Data().(string))
	gObj, _ := jobj.JSONPointer("/info/age") // JSONPointer()参数必须以/开头
	fmt.Println("age: ", gObj.Data().(float64))
	fmt.Println("one hobby: ", jobj.Path("info.hobbies.1").Data().(string))
}
