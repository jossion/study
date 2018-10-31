package main

import "fmt"

func main() {
	//	先定义一个实际变量
	a := 123
	//指针变量
	var ip *int
	//	给指针变量赋值
	ip = &a
	fmt.Printf("a的类型是%T \t，值是%v\t \n",a,a)
	fmt.Printf("&a的类型是%T\t，值是%v \t\n",&a,&a)
	fmt.Printf(" ip的类型是%T\t，值是%v\t \n",ip,ip)
	fmt.Printf(" *ip的类型是%T\t，值是%v\t \n",*ip,*ip)
	fmt.Printf("&ip的类型是%T\t,值是%v \t \n",&ip,&ip)
}
