package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// cli是一个用于构建命令行程序的库
// cli非常简洁，所有的初始化操作就是创建一个cli.App结构的对象。通过为对象的字段赋值来添加相应的功能
func main() {
	app := &cli.App{
		Name:  "hello",
		Usage: "hello world example",
		Action: func(context *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
