package main

import (
	"fmt"
	"time"
)

func main() {
	arrA := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	BubbleSortA(arrA)
	fmt.Println(arrA)
	arrB := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	BubbleSortB(arrB)
	fmt.Println(arrB)
}

func BubbleSortA(arr []int) {
	//x:=time.Date(2017,12,22,24,59,59,88,time.Local)
	//fmt.Println(x.Format("2006-01-02 15:04:05"))
	//starttime := time.Now()
	iCount := 0 // 外部循环次数
	jCount := 0 // 内部循环次数
	//双层For循环实现
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//判断相邻两个数大小
			if arr[j] > arr[j+1] {
				//如果前者大，则换位
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			jCount++
		}
		iCount++
	}
	fmt.Println("外部循环（i）的次数：", iCount)
	fmt.Println("内部循环（j）的次数：", jCount)
	//stoptime := time.Now()
	//fmt.Printf("开始时间：%d \t 结束时间：%d \t 用时：%n", starttime, stoptime, stoptime.Sub(starttime))

}

func BubbleSortB(arr []int) {
	fmt.Println(time.Now())
	iCount := 0 // 外部循环次数
	jCount := 0 // 内部循环次数
	//双层For循环实现
	for i := 0; i < len(arr)-1; i++ {
		//定义标记，判断本次循环是否需要换位，如果没有换位说明排序正常，可以用break跳出循环
		flag := true
		for j := 0; j < len(arr)-1-i; j++ {
			//判断相邻两个数大小
			if arr[j] > arr[j+1] {
				//如果前者大，则换位
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//	如果换位说明排序没结束，循环继续
				flag = false
			}
			jCount++
		}
		iCount++
		if flag {
			break
		}
	}
	fmt.Println("外部循环（i）的次数：", iCount)
	fmt.Println("内部循环（j）的次数：", jCount)
	fmt.Println(time.Now())
}

/*
冒泡排序的原则：小的数字排列在左边大的排列在右边
*/
