package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ozzo-validation是一个非常强大的，灵活的数据校验库 ozzo-validation提倡用代码指定规则来进行校验
func main() {
	name := "nfx"

	err := validation.Validate(name,
		validation.Required,
		validation.Length(2, 10),
		is.URL)
	fmt.Println(err)
}
