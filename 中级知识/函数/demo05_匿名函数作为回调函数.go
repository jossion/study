package main

import (
	"fmt"
	"math"
)

//遍历切片
type myFuncs func(float64) string

func main() {
	//	定义一个slice，对其中数据求平方根和求求平方
	arr := []float64{1, 9, 16, 25, 30}
	result := Filterslice(arr, func(v float64) string {
		v = math.Sqrt(v)
		return fmt.Sprintf("%.2f", v)
	})
	fmt.Print(result)
	//求平方
	result = Filterslice(arr, func(v float64) string {
		v = math.Pow(v,2)
		return fmt.Sprintf("%.2f", v)
	})
	fmt.Print(result)
}

//遍历切片，对其中元素进行运算
func Filterslice(arr []float64, f myFuncs) []string {
	var result []string
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}
