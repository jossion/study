package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Println("--------------------------")
	fmt.Printf("b的数据类型：%T ,数值：%v \n", b, b)
	fmt.Println("b的地址", b)
	fmt.Println("*b的地址", *b)

	*b++
	fmt.Println("a的地址", a)
	var c *int = &a
	fmt.Println("--------------------------")
	fmt.Printf("c的数据类型：%T ,数值：%v \n", c, c)

}
