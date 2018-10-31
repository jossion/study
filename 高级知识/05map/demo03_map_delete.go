package main

import (
	"fmt"
)

func main() {
	//	声明并初始化Map
	map1 := map[string]string{
		"element":    "div",
		"width":      "100px",
		"height":     "200px",
		"border":     "solid",
		"background": "none",
	}
	//2.根据Key删除map中的某个人元素
	fmt.Println("删除前", map1)
	if _, ok := map1["background"]; ok {

		fmt.Println("找到了，删除中....")
		delete(map1, "background")
	}
	fmt.Println("删除后", map1)
	/*
	   3.清空map中所有元素
	   	Go语言没有为map提供清空所有元素的函数，清空map的唯一方法是重新make一个新的map赋值给原map
	*/
	map1 = map[string]string{}
	fmt.Println("清空后的Map：", map1)
	map1 = map[string]string{
		"element":    "div",
		"width":      "100px",
		"height":     "200px",
		"border":     "solid",
		"background": "none",
	}
	map1 = make(map[string]string)
	fmt.Println("清空后的Map：", map1)
}
