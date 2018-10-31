//函数变量
//函数作为值传递
package main

import (
	"fmt"
	"strings"
)
type JOPL func(string)string

func main() {
	str:="asdfbgkhiruenfovnrfb"
	//函数调用方式
	fmt.Println(processLetter(str))
	//函数作为参数传入
	fmt.Println()
	fmt.Println(StringTocase(str,processLetter))
//	函数作为type自定义类型
	fmt.Println()
	fmt.Println(StringRocase2(str,processLetter))

}
//处理字符串，实现大小写奇偶交替
func processLetter(str string)string{
	fmt.Println("基本的函数调用")
	result:=""
	for i,value:=range  str{
		if i%2==0{
			result+=strings.ToUpper(string(value))
		}else {
			result+=strings.ToLower(string(value))
		}
	}
	return  result
}

func StringTocase(str string,fixstr func(string)string) string {
	fmt.Println("函数作为参数传递")
	return fixstr(str)
}
func StringRocase2(str string,fixstr JOPL) string  {
	fmt.Println("自定义函数为数据类型作为参数传递")
	return fixstr(str)
}