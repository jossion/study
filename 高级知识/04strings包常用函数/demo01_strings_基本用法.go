package main

import (
	"fmt"
	"unicode/utf8"
)

//字符串遍历
func main() {
	s := "我爱go语言"
	fmt.Println("字节长度是：", len(s))
	fmt.Println("---------------------")
	//	for ...range遍历字符串
	for i, ch := range s {
		fmt.Printf("%d:%c \n", i, ch)
	}
	fmt.Println("---------------------")
//	遍历字节数组
for i,ch:=range []byte(s){
	fmt.Printf("%d:%x \n",i ,ch)
}
	fmt.Println("---------------------")

//遍历所有字符
count :=0
	for i,ch:=range []rune(s) {
		fmt.Printf("%d:%c \n",i ,ch)
		count++
	}
	fmt.Println("字符串长度",count)
	fmt.Println("用RuneCountInString函数取字符串长度",utf8.RuneCountInString(s))
	fmt.Println("---------------------")

}
