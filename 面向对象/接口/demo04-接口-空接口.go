/*
空接口
1.概念
	·空接口：接口中没有任何方法。任意类型都可以实现该接口
	·空接口的定义：interface{},也就是包含0个method的interface
	·用空接口表示任意数据类型 类似于Java中的object
	·空接口常用于一下情形：
		☆ pintln的参数就是空接口
		☆ 定义一个map：key是string，value是任意数据类型
		☆ 定义一个切片，其中存储任意类型的数据
*/

package main

import "fmt"

//定义空接口
type A interface {
}
type Cat struct {
	name string
	age  int
}
type Person struct {
	name string
	sex  string
}

func main() {
	var a1 A = Cat{"Mimi", 3}
	var a2 A = Person{"steven", "男"}
	var a3 A = "Learn Golang with me!"
	var a4 A = 111
	var a5 A = 3.14
	pringInfo(a1)
	pringInfo(a2)
	pringInfo(a3)
	pringInfo(a4)
	pringInfo(a5)
	fmt.Println("-------------------------")
	map1 := make(map[string]interface{})
	map1["name"] = "daniel"
	map1["age"] = 12
	map1["height"] = 1.72
	fmt.Println(map1)
	fmt.Println("-------------------------")
	//3定义切片存储任意类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
	fmt.Println("-------------------------")
	//tmep:=Cat{"死猫",3}
	TransInterface(slice1)

}

func pringInfo(a A) {
	fmt.Printf("%T,%v,\n ", a, a)
}
//接口对象转型
//接口对象.(type) ,配合switch  case语句
func TransInterface(s []interface{})  {
	for key:= range s {
		fmt.Println("第",key+1,"个数据")
		switch t:=s[key].(type) {
		case Cat:
			fmt.Printf("\t Cat对象，name属性：%s,age属性：%d \n",t.name,t.age)
		case Person:
			fmt.Printf("\t Person对象，name属性：%s,sex属性：%d \n",t.name,t.sex)
		case string:
			fmt.Printf("\t string类型，name属性：%s\n",t)
		case int:
			fmt.Printf("\t int类型，name属性：%v\n",t)
		case float64:
			fmt.Printf("\t 发loath4类型，name属性：%.2f\n",t)
		}

	}
}