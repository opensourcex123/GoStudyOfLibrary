package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

// Go 每日一库之 jennifer 代码生成库,可以用它来生成任何 Go 语言代码。
func main() {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello World!")),
	)
	fmt.Printf("%#v", f)
}
