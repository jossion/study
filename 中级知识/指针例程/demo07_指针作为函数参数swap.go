package main

import "fmt"

func main() {
	//定义两个局部变量
	a, b := 100, 200
	//传统函数返回值方式
	a, b = swap0(a, b)
	fmt.Println("第一次交换swap0")
	fmt.Println(a, b)

	//使用指针的方式
	swap(&a, &b)
	fmt.Println("第er次交换swap")
	fmt.Println(a, b)

}

//数据交换传统写法
func swap0(x, y int) (int, int) {
	return y, x
}

//指针作为参数的写法
func swap(x, y *int) {
	*x, *y = *y, *x
}
