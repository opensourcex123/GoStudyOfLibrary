package main

import (
	"fmt"
	"github.com/scylladb/termtables"
)

// termtables处理表格形式数据的输出。适用于随时随地的输出一些状态或统计数据，便于观察和调试
func main() {
	t := termtables.CreateTable()
	t.AddHeaders("User", "Age")
	t.AddRow("dj", 18)
	t.AddRow("nfx", 21)
	fmt.Println(t.Render())
}
