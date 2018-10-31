/*
（一）概念
	1.面向对象语言中，接口用于定义对象的行为。接口只定义对象应该做什么，实现这种行为的方法（实现细节）是由对象来决定
	2.在Go语言中，接口是一组方法的签名：
		·接口只指定了类型应该具有的方法，类型决定了如何实现这些方法。
		·当某个类型为接口中的所有方法提供了具体实现细节时，这个类型就被称为实现了该接口。
		·接口定义了一组方法，如果某个对象实现了该接口的所有方法，则此对象就是实现了该接口
	3.Go语言的类型都是隐式实现接口的。任何定义了接口中所有方法的类型都被称为隐式的实现了该接口

（二）接口的定义语法及示例
	1.定义接口
		type name interface {
			method1([参数列表]) [返回值]
			method2([参数列表]) [返回值]
			。。。
			methodN([参数列表]) [返回值]

		}
	2.定义结构体
		type name struct {
		//属性
		}
	3.实现接口方法
		func(name struct)method1([参数列表])[返回值] {

		}
		func(name struct)method2([参数列表])[返回值] {

		}
		........
		func(name struct)methodN([参数列表])[返回值] {

		}
*/

//
package main

import "fmt"

type Phone interface {
	call()
}
type AndroidPhone struct {
}
type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}
func (a IPhone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}
func main() {
	//	定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)//注意赋值方式不同时 结构性质的不同
	fmt.Printf("类型：%T,值：%v,内存地址：%p\n", phone, phone, phone)
	fmt.Println("---------------------------------------------------")
	phone.call()
	phone = IPhone{}//注意赋值方式不同时 结构性质的不同
	fmt.Printf("类型：%T,值：%v,内存地址：%p\n ", phone, phone, &phone)
	fmt.Println("---------------------------------------------------")
	phone.call()
}
