package main

import (
	"fmt"
	"math"
)

//定义匿名函数
func main() {
	//	1.定义时调用无参匿名函数
	func(data int) {
		fmt.Println("你的成绩：", data)
	}(98)
	//	2.定义时调用有参数的匿名函数
	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(250)
	fmt.Printf("平方根是：%.2f", result)
	//	匿名函数赋值给变量，需要时调用
	myFunc := func(data string) string {
		return data
	}
	fmt.Println()
	fmt.Println(myFunc("欢迎学习go语言"))
}
