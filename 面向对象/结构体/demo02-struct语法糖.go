package main

import "fmt"

type Emp struct {
name string
age int8
sex byte
}
func main() {
//	实例化Emp
emp1:=new(Emp)
fmt.Printf("emp1:%T ,%v,%p \n",emp1,emp1,emp1)

//不用取指针类型结构体取地址，直接使用结构体的实例即可
emp1.name="David"
emp1.age=30
emp1.sex=1
fmt.Println("------------------------")
SyntacticSugar()
}

func SyntacticSugar()  {
	arr:=[4]int{1,2,3,4}
	arr2:=&arr
	fmt.Println((*arr2)[len(arr)-1])
	fmt.Println(arr2[3])
//	slice中没有语法糖的用法，必须对指针操作
	arr3:=[]int{10,20,30,40}
	arr4:=&arr3
	fmt.Println((*arr4)[len(arr)-1])
	fmt.Println((*arr4)[0])
}