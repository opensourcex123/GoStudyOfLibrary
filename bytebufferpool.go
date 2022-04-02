package main

import (
	"fmt"
	"github.com/valyala/bytebufferpool"
)

// bytebufferpool实现了自己的Buffer类型，并使用一个简单的算法降低扩容带来的性能损失
func main() {
	b := bytebufferpool.Get()
	b.WriteString("hello")
	b.WriteByte(',')
	b.WriteString(" world!")

	fmt.Println(b.String())

	bytebufferpool.Put(b)
}
