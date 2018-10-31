/* 概念：
·继承是面向对象编程中三大特征之一。用于描述两个结构体之间的关系（类）。一个类（子类或派生类）继承于另一个类（父类或超类）
·子类可以有自己的属性和方法也可以重写父类已有的方法
·子类可以直接方位父类所有的属性和方法
·提升字段：
	△在结构体中属于匿名结构体的字段为提升字段，因为它们可以被方位，就好像它们属于拥有匿名结构体字段的结构一样
	△换句话说：父类中的字段就是提升字段
·继承的意义：
	△避免重复代码
	△扩展类的功能
·采用匿名字段的形式就是模拟继承关系，而模拟聚合关系时 一定要用有名字的结构体作为字段*/

package main

import "fmt"

type Person struct {
	name string
	age  int8
	sex  string
}
type Student struct {
	Person
	schoolName string
}

func main() {
	//1.实例化初始化Person
	p1 := Person{"Stever", 35, "男"}
	fmt.Println(p1)
	fmt.Println("-------------------------------")
	//2.实例化初始化student
	//写法1
	s1 := Student{p1, "北航软件学院"}
	printStudentInfo(s1)
	//	写法2
	s2 := Student{Person{"Josh", 39, "男"}, "清华大学"}
	printStudentInfo(s2)
	//	写法3
	s3 := Student{Person: Person{
		name: "Penn",
		age:  19,
		sex:  "男",
	},
		schoolName: "北大",
	}
	printStudentInfo(s3)
//	写法4
s4:=Student{}
s4.name="Jossion"
s4.age=100
s4.sex="男"
s4.schoolName="家里蹲大学"
printStudentInfo(s4)
}

func printStudentInfo(s1 Student) {
	fmt.Println("-----------------------------Start-------------------------------")
	fmt.Println("\n", s1)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("姓名:\t%s\t，年龄:\t%d\t,性别:\t%s\t,学校:\t%s\t \n", s1.name, s1.age, s1.sex, s1.schoolName)
	fmt.Println("------------------------------End--------------------------------")
}
