package main

import (
	"container/list"
	"fmt"
)

func main() {
	//optionList1()
	optList2()
}

func optionList1() {
	//	1.声明list方式
	var ListA list.List
	printListInfo("声明List方式一", ListA)

	//填充数据
	ListA.PushFront("xyz")
	ListA.PushFront("123")
	ListA.PushBack(3.145926)
	printListInfo("填充数据后", ListA)
	//	遍历list
	iterateList(ListA)
	iterateListReback(ListA)

}

func optList2() {
	listB := list.New()
	temp := list.New()
	temp.PushFront("a")
	temp.PushFront("b")
	temp.PushFront("c")
	printListInfo("声明listB", *listB)
	//尾部添加
	listB.PushBack("phone")
	//	头部添加
	listB.PushFront(123456)
	//	尾部添加后保存元素句柄
	element := listB.PushBack("来北京")
	//	insertAfter方法
	listB.InsertAfter("欢迎下次光临", element)
	listB.InsertBefore("欢迎您", element)
	//	遍历list2
	iterateList(*listB)
	iterateListReback(*listB)
	//	Remove元素
	listB.Remove(element)
	iterateList(*listB) //验证移除输出
	iterateListReback(*listB)
	//	PustFrontList
	listB.PushFrontList(temp)//输出验证在链表头部插入另一个链表
	iterateList(*listB)
	iterateListReback(*listB)
	listB.PushBackList(temp)//输出验证在链表尾部插入另一个链表
	iterateList(*listB)
	iterateListReback(*listB)
}
func printListInfo(info string, l list.List) {
	fmt.Println("-------------" + info + "-------------")
	fmt.Printf("类型：%T\t 值：%v\t 长度：%d \n", l, l, l.Len())
	fmt.Println("--------------------------------------")
}
//顺序遍历
func iterateList(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println("\n------------------------------------")
		fmt.Printf("%v \t ", e.Value)
		//fmt.Printf("第%d个元素是：%v\t \n",e,e.Value)

	}
	fmt.Println("\n------------------------------------")
}
//逆序遍历
func iterateListReback(l list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		//fmt.Println("\n------------------------------------")
		fmt.Printf("%v \t ", e.Value)
		//fmt.Printf("第%d个元素是：%v\t \n",e,e.Value)

	}
	fmt.Println("\n------------------------------------")
}
