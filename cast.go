package main

import (
	"fmt"
	"github.com/spf13/cast"
)

// cast是一个小巧、实用的类型转换库，用于将一个类型转为另一个类型
func main() {
	// i = indirect(i)	// 调用indirect函数将参数中可能的指针去掉。如果类型本身不是指针，那么直接返回。否则返回指针指向的值。循环直到返回一个非指针的值
	fmt.Println(cast.ToString(8))
	fmt.Println(cast.ToString(nil))

	fmt.Println(cast.ToInt(8.31))
}
