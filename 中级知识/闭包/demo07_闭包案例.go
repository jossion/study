package main

import "fmt"

func main() {
	res := Counter()
	fmt.Printf("%T \n", res)
	fmt.Println("Res地址是:\t", res)
	fmt.Println("第一次\t res值是:\t", res())
	fmt.Println("第二次\t res值是:\t", res())
	fmt.Println("第三次\t res值是:\t", res())
	res = Counter()
	fmt.Println("第四次\t res值是:\t", res())
	fmt.Println("第五次\t res值是:\t", res())
	fmt.Println("第六次\t res值是:\t", res())
}

//闭包函数，实现计数器功能
func Counter() func() int {
	i := 0
	res := func() int {
		i++
		return i
	}
	fmt.Println("Counter内部：", res)
	return res
}
