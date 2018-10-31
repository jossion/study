package main

import "fmt"

/*
1.声明
var name [len} type
2.初始化



*/
var arr [3]int
var arr2 = [4]int{1, 2, 3, 4}

func main() {
	a := [4]float64{2.3, 4, 5, 6.6}
	fmt.Println(a)
	b := [...]int{2, 3, 4, 5, 6, 7}
	fmt.Println(len(b))
	//	遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")

	}
	fmt.Println()
	//	遍历数组2
	for _, value := range b {
		fmt.Print(value, "\t")
	}
}
