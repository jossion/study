package main

import (
	"fmt"
	"strings"
)


func main() {
	strA := "seafood"
	strB := "foo"
	strC := "bar"

	if TestContains(strA,strB){
		bh:="包含"
		fmt.Printf("字符串StrA中%s字符串StrB",bh)
	}else {
		bh:="不包含"
		fmt.Printf("字符串StrA中%s字符串StrB",bh)
	}
	fmt.Println()
	if TestContains(strA,strC){
		bh:="包含"
		fmt.Printf("字符串StrA中%s字符串StrC",bh)
	}else {
		bh:="不包含"
		fmt.Printf("字符串StrA中%s字符串StrC",bh)
	}


}

//测试strings包里的15个检索函数
func TestContains(str1,str2 string )bool{
	return strings.Contains(str1, str2)
}
func TestContainsAny(str1,str2 string)bool{
	return strings.ContainsAny(str1,str2)
}
func TestContainsRune(str1 string ,r rune )bool{
	return strings.ContainsRune(str1,r)
}
func TestCount(str1,str2 string)int{
	return strings.Count(str1,str2)
}
