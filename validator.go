package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Go 每日一库之 validator
// validator用于对数据进行校验。在 Web 开发中，对用户传过来的数据我们都需要进行严格校验，防止用户的恶意请求。

type User struct {
	Name string `validate:"min=6,max=10"`  // 约束字符串的长度
	Age  int    `validate:"min=1,max=100"` // 约束整形的大小
}

func main() {
	validate := validator.New()

	u1 := User{Name: "lidajun", Age: 18}
	err := validate.Struct(u1)
	fmt.Println(err)

	u2 := User{Name: "dj", Age: 10}
	err = validate.Struct(u2)
	fmt.Println(err)
}
