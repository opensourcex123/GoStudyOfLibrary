package main

import (
	"fmt"
	"github.com/valyala/fasttemplate"
)

// fasttemplate是一个比较简单、易用的小型模板库。
// fasttemlate只专注于一块很小的领域——字符串替换
func main() {
	template := `name: {{name}}
age: {{age}}`
	t := fasttemplate.New(template, "{{", "}}")
	s1 := t.ExecuteString(map[string]interface{}{
		"name": "dj",
		"age":  "18",
	})
	s2 := t.ExecuteString(map[string]interface{}{
		"name": "nfx",
		"age":  "20",
	})
	fmt.Println(s1)
	fmt.Println(s2)
}
