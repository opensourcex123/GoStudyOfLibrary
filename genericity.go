package main

import "fmt"

// Print 泛型
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// Stringer 接口类型约束
type Stringer interface {
	String() string
}

func Stringify[T Stringer](s []T) (res []string) {
	for _, v := range s {
		res = append(res, v.String())
	}
	return res
}

// AnyString 近似约束元素
type AnyString interface {
	~string
}

// Desensitization 这里的T可以代表各种以string为别名的类型
func Desensitization[T AnyString](str T) string {
	var newStr string
	// 逻辑
	//newStr = desensitization(str)
	return newStr
}

// PrintInt64OrFloat64 只能用确定的类型进行联合类型约束
func PrintInt64OrFloat64[T int64 | float64](t T) {
	fmt.Printf("%#v", t)
}

// Filter 过滤函数，对src中的每个元素执行f函数，如果返回true，则加入到结果中
func Filter[T any](f func(T) bool, src []T) []T {
	var res []T
	for _, v := range src {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	s := []int{-1, -2, 1, 2, 3}
	Print(s)

	// 过滤掉负数
	res := Filter(func(t int) bool {
		return t >= 0
	}, s)
	fmt.Println(res)
}
