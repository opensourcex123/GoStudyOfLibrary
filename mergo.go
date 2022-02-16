package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
)

// Go 每日一库之 mergo 一个合并结构体字段的库
// 可以在相同的结构体或map之间赋值，可以将结构体的字段赋值到map中，可以将map的值赋值给结构体的字段

type redisConfig struct {
	Address string
	Port    int
	DB      int
}

var defaultConfig = redisConfig{
	Address: "127.0.0.1",
	Port:    6381,
	DB:      1,
}

func main() {
	var config redisConfig
	// 结构体赋值给结构体
	if err := mergo.Merge(&config, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println("redis address: ", config.Address)
	fmt.Println("redis port: ", config.Port)
	fmt.Println("redis db: ", config.DB)
	// 结构体赋值给map
	var m = make(map[string]interface{})
	if err := mergo.Map(&m, defaultConfig); err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
