package main

import "fmt"

type Human struct {
	name string
	age  int8
	sex  byte
}

func main() {
	h1 := Human{"stever", 35, 1}
	fmt.Printf("h1:%T,%v,%p\n", h1, h1, &h1)
	fmt.Println("------------------------------")
	//	copy结构体
	h2 := h1
	h2.name = "David"
	h2.age = 18
	fmt.Printf("h2:%T,%v,%p\n", h2, h2, &h2) //h2:main.Human,{David 18 1},0xc000044480
	fmt.Printf("h1:%T,%v,%p\n", h1, h1, &h1) //h1:main.Human,{stever 35 1},0xc000044400
	fmt.Println("------------------------------")
	//将结构体作为参数传递
	changeName(h1)
	fmt.Printf("h1:%T,%v,%p\n", h1, h1, &h1) //h1:main.Human,{stever 35 1},0xc000078060
}

func changeName(h Human) {
	h.name = "Daniel"
	h.age = 13
	fmt.Printf("函数体内修改后：%T,%v,%p\n", h, h, &h)//函数体内修改后：main.Human,{Daniel 13 1},0xc000078180
}
