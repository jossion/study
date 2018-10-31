package main

import "fmt"

type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

func main() {
	//深拷贝，struct是值类型，默认就是深拷贝
	d1 := Dog{"doudou", "黑的", 2, "二哈"}
	fmt.Printf("h1:%T,%v,%p\n", d1, d1, &d1)
	d2 := d1
	fmt.Printf("h2:%T,%v,%p\n", d2, d2, &d2)
	d2.name = "maomao"
	fmt.Printf("h2修改后:%T,%v,%p\n", d2, d2, &d2)
	//	2.浅拷贝，直接赋值指针地址
	d3 := &d1
	fmt.Printf("h1:%T,%v,%p\n", d3, d3, d3)
	d3.name = "qiuqiu"
	d3.color = "白色"
	d3.kind = "萨摩耶"
	fmt.Println("d3修改后", d3)
	fmt.Println("原来的d1:", d1)
	fmt.Println("------------------------------")
//	3.浅拷贝，通过New（）函数实例化
d4:=new(Dog)
	d4.name = "多多"
	d4.color = "棕色"
	d4.kind = "八哥犬"
	d5:=d4
	fmt.Printf("h4:%T,%v,%p\n", d4, d4, d4)
	fmt.Printf("h5:%T,%v,%p\n", d5, d5, d5)
	fmt.Println("------------------------------")
	d5.color="金色"
	d5.kind="金毛犬"
	fmt.Printf("h4:%T,%v,%p\n", d4, d4, d4)
	fmt.Printf("h5:%T,%v,%p\n", d5, d5, d5)
	fmt.Println("------------------------------")
}
/*
总结：
struct操作时，copy指针是浅拷贝，操作副本时原始数据也会被修改，因为正副本指向的内存地址是同一个，值copy是深拷贝（系统默认方式）
正副本数据指向不同的内存空间，所以操作任何一个时不会影响另一个的值
*/