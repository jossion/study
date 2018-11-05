package main

import (
	"fmt"
)

func main() {
	var ptr *int

	fmt.Printf("ptr类型为%T，值为%v \n", ptr, ptr)

	if ptr == nil {
		fmt.Println("空指针")
	} else {
		fmt.Println("非空指针")
	}
}
