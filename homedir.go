package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
)

// homedir库的学习

//go-homedir有两个功能：
//	Dir：获取用户主目录；
//	Expand：将路径中的第一个~扩展成用户主目录。
func main() {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Home dir:", dir)

	dir = "~/golang/src"
	expandedDir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Expand of %s is: %s\n", dir, expandedDir)
}
