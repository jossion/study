package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4,}
	fmt.Println("调用函数前",a)
	modify(&a)
	fmt.Println("调用函数后",a)
}

func modify(arr *[]int) {
	(*arr) [0] = 250
}
