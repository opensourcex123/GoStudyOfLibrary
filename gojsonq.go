package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
)

// Go 每日一库之 gojsonq 它可以帮助我们很方便的操作 JSON。
func main() {
	content := `{
  "user": {
    "name": "dj",
    "age": 18,
    "address": {
      "provice": "shanghai",
      "district": "xuhui"
    },
    "hobbies":["chess", "programming", "game"]
  }
}`
	gq := gojsonq.New().FromString(content)
	district := gq.Find("user.address.district")
	fmt.Println(district)

	gq.Reset() // 如果不重置，下一个查询就会在上一次查找的节点开始
	// gpCopy := gp.Copy() 可以使用此代替，下面同样改成gpCopy

	hobby := gq.Find("user.hobbies.[0]")
	fmt.Println(hobby)
}
