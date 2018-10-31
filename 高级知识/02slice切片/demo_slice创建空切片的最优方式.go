package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var sa[]string
	sa := make([]string,0,20)
	if sa == nil {
		fmt.Printf("声明的切片为 nil  长度为：%d \n",len(sa))
	}else {
		fmt.Printf("声明的切片不为 nil 长度为：%d \n",len(sa))
	}
	for i :=0;i<15 ;i++  {
		sa = append(sa,strconv.Itoa(i))
		printSlice1("sa",sa)
	}
	printSlice1("sa",sa)
}
func printSlice1(name string, x []string) {
	fmt.Print(name, "\t")
	fmt.Printf("地址： %p \t len=%d \t cap=%d \t value=%v \n", x, len(x), cap(x), x)
}

/*
结论：
1.创建slice时，请使用make函数
2.设置make函数的第三个参数cap为一个相对合适的数值，以避免当slice容量不够时系统自动创建新的内存空间

 */