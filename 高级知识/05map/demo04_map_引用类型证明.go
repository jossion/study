
/*

 */
package main

import (
	"fmt"
)

func main() {
//	初始薪资
salary:= map[string]float64{
	"Steven": 15000,
	"Josh":20000,
	"Daniel":5000,
}
fmt.Println("初始薪资",salary)
newSalary:=salary
newSalary["Daniel"]=8000
fmt.Println("Daniel调整后：",newSalary)//8000
	fmt.Println("Daniel调整前：",salary)//8000
//	说明新的变量的改变同时影响到原始变量的值
//涨薪后
changeSalary(salary)

fmt.Println("原始薪资salary是否受涨薪函数影响",salary)
}
//函数传参的方式会怎样？
func changeSalary(m map[string]float64)  {
	for key := range m {
		m[key]*=1.1
	}
}
/*
说明map不管是何种使用方式（直接引用或是函数传参）都是引用类型
 */