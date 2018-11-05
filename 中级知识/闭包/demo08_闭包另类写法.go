package main

import "fmt"

func main() {
	res := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}() //（）表示此函数为立即执行函数
	fmt.Println("第一次\t res值是:\t", res())
	fmt.Println("第二次\t res值是:\t", res())
	fmt.Println("第三次\t res值是:\t", res())
}
