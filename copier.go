package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

// Go 每日一库之 copier copier库能处理不同类型之间的赋值
// copier还能：
//	调用同名方法为字段赋值；
//	以源对象字段为参数调用目标对象的方法，从而为目标对象赋值（当然也可以做其它的任何事情）；
//	将切片赋值给切片（可以是不同类型哦）；
//	将结构体追加到切片中。

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	user := User{Name: "dj", Age: 10}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
