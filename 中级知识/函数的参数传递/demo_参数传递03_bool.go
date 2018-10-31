package main

import (
	"fmt"
)

/*
1.值传递：是指在调用函数时将实参赋值一份传递到函数中，这样在函数中如果对参数进行修改，不会影响到实际参数。
2.引用传递：是指在调用函数时将实参的地址传递到函数中，那么在函数中对参数所进行的修改将影响到实参。
3.Go语言中可以借助指针来实现引用传递的效果。函数参数使用指针参数，传参数时其实是在拷贝一份指针参数，也就是copy了变量（实参）地址
*/
func main() {
	a := true
	fmt.Printf("\n1.变量a的内存地址：%p，值：%v \n\n ", &a, a)
	//	传值
	changeBoolVal(a)
	fmt.Printf("\n2.changeBoolVal函数调用后：\n 变量a的内存地址：%p，值：%v \n\n ", &a, a)
	//	通过传指针打到传引用的效果
	changeStrinPtr(&a)
	fmt.Printf("\n3.changeBoolPtr函数调用后：\n 变量a的内存地址：%p，值：%v \n\n ", &a, a)
}

func changeBoolVal(a bool) {
	fmt.Printf("------------ChangeBoolVal函数内：值参数a的内存地址：%p,值为%v \n", &a, a)
	a = false
}

func changeStrinPtr(a *bool) {
	fmt.Printf("------------ChangeBoolVal函数内：指针参数a的内存地址：%p,值为:%v \n", &a, a)
	*a = false
}
