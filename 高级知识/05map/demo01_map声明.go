/*
map是Go中的内置类型，它将一个值和一个键关联起来
1.map是一种无序的键值对集合，map通过key来快速检索数据，key类似索引
2.map可以向遍历数组或切片那样去遍历它，不过map是无序的，说以无法决定他的展示顺序，这是因为map是使用hash表来实现的
使用map要注意的细节：
1.map是无序的，每次输出的map都会不一样，它不能通过index获取，必须通过key获取；
2.map的长度不固定，和slice一样可以扩展。内置的len（）函数同样适用于map没反悔map拥有的键值对的数量；但是map不能通过cap（）
  函数计算容量
3.同一个map中的key是唯一的；key的数据类型必须是可参与不叫运算的类型，也就是支持==或!=操作的类型，包括：bool int float string
array。对于slice 和function等引用类型则不能作为键的类型；
4.map的value可以是任何数据类型
5.和slice一样 map是引用类型
Map声明方式：
1.var name map[key type]value type 不要加多余的空格；
2.var声明的map 默认map是nil；nil map是不能用来存放键值对的，因为没有获得内存空间；
3.var声明后，要么生命是就初始化，要么需要make（）函数分配到内存空间，这样才能在其中存放键值对；使用make函数
  name:=make{map[key type]value type},这样声明的map即便不初始化map的值一不等于nil
往map中存放键值对（key-value pair）
（1）声明是初始化数值

*/
package main

import (
	"fmt"
)

func main() {
	//	1、声明方式1
	var map1 map[string]string
	printMapInfo(map1)
	//2、声明方式2
	map2 := make(map[string]string)
	printMapInfo(map2)
	//	3、 map中的key可以是：int\float\bool\string; 不能是：slice，func，map
	var (
		m1 map[int]string
		m2 map[float64]string
		m3 map[bool]string
		//m4 map[m1]string

	)
	fmt.Println(m1, m2, m3)
}
func printMapInfo(m map[string]string) {

	fmt.Printf("%T，%v，len(map)=%d,m==nil?%v \n", m, m, len(m), m == nil)
	fmt.Println("-----------------------------------------------")
}
