package main

import "fmt"

type Teacher struct {
	name string
	age  int8
	sex  byte
}

func main() {
	//	1'var声明方式实例化结构体
	var t1 Teacher

	fmt.Printf("t1%T,%v,%q\n", t1, t1, t1)
	t1.name = "stever"
	t1.age=35
	t1.sex=1
	fmt.Println(t1)
	fmt.Println("-------------------------")
//	2.变量剪短声明格式实例化结构体
	t2 :=Teacher{}
	t2.name="David"
	t2.age=40
	t2.sex=1
	fmt.Println(t2)
	fmt.Println("-------------------------")
//	3.变量简短声明格式实例化结构体,s声明时初始化
	t3:=Teacher{
		name:"Josh",
		age:28,
		sex:1,
	}
	fmt.Println(t3)
	fmt.Println("-------------------------")
//	4.变量简短声明格式实例化结构体
t4:=Teacher{"Jos",27,0}
	fmt.Println(t4)
	fmt.Println("-------------------------")
//5.创建指针类型的结构体
t5:=new(Teacher)
	//fmt.Printf("t5%T,%v,%p\n", t5, t5, t5)
	(*t5).name="Jihe"
	(*t5).age=22
	(*t5).sex=1
	fmt.Println(t5)
	fmt.Println(&t5)
	fmt.Println(*t5)
	fmt.Println("-------------------------")
	//语法简写形式---语法糖（用于指针类型）
	t5.name="Jihe2"
	t5.age=23
	t5.sex=1
	fmt.Println(t5)
	fmt.Println(&t5)
	fmt.Println(*t5)
	fmt.Println("-------------------------")
}
