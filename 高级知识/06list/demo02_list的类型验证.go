package main

import (
	"container/list"
	"fmt"
)

func main() {
	copyList()
	/*
	结论：
		list是值类型，但是创建时使用New方法创建，返回的是地址指针，copy或传参时操作的是地址指针，导致元素操作是就具有了引用类型的特征
	*/
}
func copyList() {
	listA := list.New()
	printListInfo2("声明list", listA)
	//	listA 赋值
	listA.PushBack("one")
	listA.PushBack(2)
	listA.PushBack("三")
	printListInfo2("赋值后的listA", listA)
	iterateList2(listA)
	//	将listA拷贝给listB
	listB := listA
	printListInfo2("刚拷贝的listB ", listB)
	iterateList2(listB)
	//修改listB
	listB.PushBack("250")
	listB.PushBack("350")
	listB.PushBack("450")
	printListInfo2("修改后的listB", listB)
	iterateList2(listB)
	//	测试输出ListA
	printListInfo2("修改后的listA", listA)
	iterateList2(listA)
	//	listB调用Init后
	listA.Init()
	printListInfo2("listB调用Init后", listA)
	iterateList2(listA)

}

func printListInfo2(info string, l *list.List) {
	fmt.Println("-------------" + info + "-------------")
	fmt.Printf("类型：%T\t 值：%v\t 长度：%d \n", l, l, l.Len())
	fmt.Println("--------------------------------------")
}

//顺序遍历
func iterateList2(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println("\n------------------------------------")
		fmt.Printf("%v \t ", e.Value)
		//fmt.Printf("第%d个元素是：%v\t \n",e,e.Value)

	}
	fmt.Println("\n------------------------------------")
}

//逆序遍历
func iterateListReback2(l *list.List) {
	for e := l.Back(); e != nil; e = e.Prev() {
		//fmt.Println("\n------------------------------------")
		fmt.Printf("%v \t ", e.Value)
		//fmt.Printf("第%d个元素是：%v\t \n",e,e.Value)

	}
	fmt.Println("\n------------------------------------")
}
