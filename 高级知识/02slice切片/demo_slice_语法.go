package main

import "fmt"

/*
1.slice （“动态数组”），长度不固定，可以追加或删除元素
2.Slice本身没有任何数据，它是对现有数组的引用
3.从概念上讲：包括三个元素：（指针 长度 最大长度）
*/
func main() {
	arr:=[5]int{1,2,3,4,5}
	//Slice初始化
	//直接声明
	S:=[]int{1,2,3}
	fmt.Printf("直接声明的Slice S的内容是%x，\n 长度是：%d \t，最大长度是：%d\t\n",S,len(S),cap(S))
	s_All:=arr[:]//截取数组的全部元素
	s_LeftToCenter:=arr[0:2]             //截取部分数组 从前左向右截取
	s_CentertoRight:=arr[2:]//截取部分数组 从中间到最后向右截取
	fmt.Printf("截取数组全部元素：%d；类型是%T\n",s_All,s_All)
	fmt.Printf("截取数组从前到中间某：%d\n",s_LeftToCenter)
	fmt.Printf("截取数组从中间某处到后：%d\n",s_CentertoRight)

//	使用make函数创建
	s_new:=make([]byte,5,10)
	fmt.Printf("\n Make出的Slice的内容是%d；\n 长度是:%d \t最大长度是：%d\t",s_new,len(s_new),cap(s_new))
}
