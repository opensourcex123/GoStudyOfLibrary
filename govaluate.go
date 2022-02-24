package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
)

// govaluate与 JavaScript 中的eval功能类似，用于计算任意表达式的值
func main() {
	expr, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Fatal("syntax error", err)
	}

	res, err := expr.Evaluate(nil)
	if err != nil {
		log.Fatal("evaluate error", err)
	}

	fmt.Println(res)

	// 使用函数
	functions := map[string]govaluate.ExpressionFunction{
		"strlen": func(arguments ...interface{}) (interface{}, error) {
			length := len(arguments[0].(string))
			return length, nil
		},
	}

	exprString := "strlen('teststring')"
	expr2, _ := govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	res, _ = expr2.Evaluate(nil)
	fmt.Println(res)
}
