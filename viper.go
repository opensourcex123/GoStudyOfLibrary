package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// go每日一库之viper库。viper 是一个配置解决方案，拥有丰富的特性,viper库可以监听配置文件修改，不需要重启服务器就可以让配置生效
//	设置文件名时不要带后缀；
//	搜索路径可以设置多个，viper 会根据设置顺序依次查找；
//	viper 获取值时使用section.key的形式，即传入嵌套的键名；
//	默认值可以调用viper.SetDefault设置。
func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed", err)
	}

	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
}
