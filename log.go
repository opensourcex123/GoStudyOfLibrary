package main

import "log"

// Go标准库提供了一个日志库log
// log默认输出到标准错误（stderr），每条日志前会自动加上日期和时间。如果日志不是以换行符结尾的，那么log会自动加上换行符。即每条日志会在新行中输出。

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "dj",
		Age:  18,
	}

	log.SetPrefix("Login:") // 为每一条日志增加前缀
	log.Printf("%s login, age:%d", u.Name, u.Age)
	log.Panicf("Oh, system error when %s login", u.Name)
	log.Fatalf("Danger! hacker %s login", u.Name)
}
