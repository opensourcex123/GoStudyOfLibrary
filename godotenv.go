package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// godotenv库从.env文件中读取配置， 然后存储到程序的环境变量中。在代码中可以使用读取非常方便。
func main() {
	//err := godotenv.Load()	// 导入github.com/joho/godotenv/autoload，配置会自动读取
	//if err != nil {			// Load接收多个文件名作为参数，如果不传入文件名，默认读取.env文件的内容。 如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效。
	//	log.Fatal(err)
	//}

	fmt.Println("name:", os.Getenv("name"))
	fmt.Println("age:", os.Getenv("age"))
	fmt.Println(os.Getenv("GOPATH"))
}
