package main

import (
	"GoStudyOfLibrary/quicktemplate/templates"
	"fmt"
)

// 一个生成代码框架的工具，最大程度地降低重复劳动
// quicktemplate功能强大，语法简单，使用方便
func main() {
	fmt.Println(templates.Greeting("dj", 5))
}
