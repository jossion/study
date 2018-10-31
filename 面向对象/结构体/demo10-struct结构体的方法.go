
/*
（一）方法的本质是函数，但又和函数有不同点
1.含义不同
	·函数function是一段具有独立功能的胆码，可以被反复多次调用，，从而实现代码的复用，而方法method是一个类的行为功能，只有该类的对象才能调用
2.方法有接受者，而函数没有
	·Go语言的方法method是一种作用于特定类型的变量函数，这种特定类型变量叫做Receiver（接受者、接收者、接收器）
	·接收者的概念类似于传统面向对象语言中的this 或self关键字
	·一个方法就是一个包含了接收者的函数
	·Go语言中接收者的类型可以是任何类型，不仅仅是结构体，也可以是其他任何类型
3.函数不可以重命名，而方法名可以一样
（二）方法的语法格式：
	func（接收者 类型）方法名（参数列表）（返回值列表）{
		方法体
	}
	·接收者在func关键字和方法名之间编写，接收者可以使struct类型或非struct类型，可以是指针类型和非指针类型
	·接收者的变量在命名时，官方建议使用接收者类型的第一个小写字母
	·实例
*/


package main

import "fmt"

type Employee struct {
name,currency string
salary float64
}
func main() {
	emp1:=Employee{"Daniel","$",2000}
	emp1.printSalary()
	printSalary(emp1)
}
//printsalary 方法 method
func (e Employee)printSalary()  {
	fmt.Printf("员工姓名：%s，薪资%s%.2f \n",e.name,e.currency,e.salary)
}

//printsalary 函数 function
func printSalary(e Employee)  {
	fmt.Printf("员工姓名：%s，薪资%s%.2f \n",e.name,e.currency,e.salary)
}