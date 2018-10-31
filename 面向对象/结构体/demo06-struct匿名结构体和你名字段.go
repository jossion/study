/*
匿名结构体：
	（1）概念
			·没有名字的结构体。无需通过type关键字定义就可以直接使用
			·在创建匿名结构体时，同时要创建对象
			·匿名结构体由结构体定义和键值对初始化两个部分组成
			·语法格式：
				name:=struct{
					定义成员属性
				}{初始化成员属性}

*/

package main

import "fmt"

func main() {
noNameStruct()
}

func noNameStruct()  {
	//	匿名结构体
	stru := struct {
		province, city string
	}{"陕西", "西安"}
	fmt.Println(stru)
}

