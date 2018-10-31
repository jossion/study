package main

import "fmt"

/*
1.定义一个函数的类型
2.实现定义的函数类型
3.作为参数调用
*/
type myFunc func(int) bool

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 34, 23, 4, 123}
	fmt.Println("slice= ", arr)
	//	获取奇数元素
	odd := Filter(arr, isodd)
	fmt.Println("奇数元素是：", odd)
	//	获取偶数元素
	even := Filter(arr, isEven)
	fmt.Println("偶数元素是：", even)
}

//判断整形元素是偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

//判断整形元素是奇数
func isodd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

//根据函数处理切片，实现奇数偶数分组，返回新的切片
func Filter(arr []int, f myFunc) []int {
	var result []int
	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
