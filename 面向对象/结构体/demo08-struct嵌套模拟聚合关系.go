/*
结构体嵌套模拟实现功能：
	聚合关系：一个类作为另一个类的附属
	继承关系：一个类作为另一个类的子类。
*/

package main

import "fmt"

type Address struct {
	province, city string
}
type Person struct {
	name    string
	age     int8
	address *Address
}

func main() {
	//	结构体对象模拟聚合关系案例
	p:=Person{}
	p.name="stever"
	p.age=35

	//第一种赋值
	addr:=Address{}
	addr.province="北京"
	addr.city="海淀"
	p.address=&addr
	fmt.Println(p)
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：城市：%s\t 区：%s\t \n",p.name,p.age,p.address.province,p.address.city)
	//修改Person对象
	p.address.city="昌平"
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：城市：%s\t 区：%s\t \n",p.name,p.age,p.address.province,p.address.city)
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：城市：%s\t 区：%s\t \n",p.name,p.age,p.address.province,addr.city)
	//修改Address对象
	addr.city  ="大兴"
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：城市：%s\t 区：%s\t \n",p.name,p.age,p.address.province,p.address.city)
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：城市：%s\t 区：%s\t \n",p.name,p.age,p.address.province,addr.city)
	//第二种赋值
	p.address=&Address{
		province:"陕西省",
		city:"西安",
	}
	fmt.Printf("姓名：%s\t 年龄：%d\t 地址：省：%s\t 市：%s\t \n",p.name,p.age,p.address.province,p.address.city)

}
/*
总结：
使用struct嵌套方式时，如果实现的是结构体聚合，存在两种情况：
	（1）嵌入的结构体数据改变时，原结构体数据不改变
			嵌入时使用结构体对象类型嵌入
	（2）嵌入的结构体数据改变时，原结构体数据改变
			嵌入时使用结构体对象指针类型嵌入
*/