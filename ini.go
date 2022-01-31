package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
)

// go-ini库的学习 go-ini是 Go 语言中用于操作 ini 文件的第三方库。

// 步骤
//		首先调用ini.Load加载文件，得到配置对象cfg；
//		然后以分区名调用配置对象的Section方法得到对应的分区对象section，默认分区的名字为""，也可以使用ini.DefaultSection；
//		以键名调用分区对象的Key方法得到对应的配置项key对象；
//		由于文件中读取出来的都是字符串，key对象需根据类型调用对应的方法返回具体类型的值使用，如上面的String、MustInt方法。
func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal("Fail to read file", err)
	}

	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

	fmt.Println("MySql IP:", cfg.Section("mysql").Key("ip").String())
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
}
