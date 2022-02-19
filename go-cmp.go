package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

// Go 每日一库之 go-cmp 是一个比较库，提供多种选项，Google开源，最初用在测试中

type Contact struct {
	Phone string
	Email string
}

type User struct {
	Name    string
	Age     int
	Contact *Contact
}

func main() {
	u1 := User{Name: "dj", Age: 18}
	u2 := User{Name: "dj", Age: 18}

	fmt.Println("u1 == u2?", u1 == u2)
	fmt.Println("u1 equals u2?", cmp.Equal(u1, u2))

	c1 := &Contact{Phone: "123456", Email: "dj@example.com"}
	c2 := &Contact{Phone: "123456", Email: "dj@example.com"}

	u1.Contact = c1
	u2.Contact = c1
	fmt.Println("u1 == u2 with same pointer", u1 == u2)
	fmt.Println("u1 equals u2 with same pointer", cmp.Equal(u1, u2))

	u2.Contact = c2
	fmt.Println("u1 == u2 with different pointer", u1 == u2)              // 只比较指针，不去比较指向的值
	fmt.Println("u1 equals u2 with different pointer", cmp.Equal(u1, u2)) // 比较指向的值
}
