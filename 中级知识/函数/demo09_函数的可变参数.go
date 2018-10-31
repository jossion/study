package main

import "fmt"

func main() {
	res := AddSum(1, 2, 3)

	fmt.Println("累加值为： ", res)
	//传slice作为可变参数集合
	arr := []int{60, 78, 90, 69, 73, 81}
	fmt.Println("切片的累加值：", AddSum(arr...))
	arr1 := []float64{60, 78, 90, 69, 73, 81}
	sum , avg , count := GetScore(arr1...)
	fmt.Printf("学员共有%d门成绩 \n总成绩:：%.2f \n 平均成绩为：%.2f",count,sum , avg)
}

//累加求和，参数个数不定
func AddSum(nums ...int) (sum int) {
	fmt.Printf("%T \n", nums)
	for _, value := range nums {
		sum += value
	}
	return
}

//累加求平均值
func GetScore(scores ... float64) (sum , avg float64 , count int) {
	for _,value:=range scores {
		sum += value
		count++
	}
	avg = sum/float64(count)
	return
}
