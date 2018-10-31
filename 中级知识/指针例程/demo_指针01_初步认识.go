package main

import "fmt"

func main() {
	a:=10
	fmt.Printf("变量a：%x \n",a)
	fmt.Printf("变量a的地址：%x \n",&a)
	b:=[]int{1,3,5,7,}
	fmt.Printf("变量b：%x \n",b)
	fmt.Printf("变量b的地址：%x \n",&b)
}
