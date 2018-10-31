package main

import "fmt"

//append copy
//
func main() {
	fmt.Println("1.------------append------------")
	var numbers []int
	printSlice("numbers: \t", numbers)
	//append 一个元素
	numbers = append(numbers, 0)
	printSlice("numbers:  \t", numbers)
	//append 多个元素
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7)
	printSlice("numbers:  \t", numbers)

	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...)
	printSlice("numbers: \t", numbers)
	//切片的删除
	fmt.Println("2.------------Del------------")
	//删除第一个
	numbers = numbers[1:]
	printSlice("numbers: \t", numbers)
	//删除最后一个
	numbers = numbers[:len(numbers)-1]
	//删除中间一个，比如 删除中间
	a := int(len(numbers) / 2)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlice("numbers: \t", numbers)
	//创建目标切片
	fmt.Println("3.------------make------------")
	numbers1 := make([]int, len(numbers), cap(numbers)*2)
	//	copy numbers to numbers1
	count := copy(numbers1, numbers)
	fmt.Println("拷贝的个数", count)
	printSlice("numbers1", numbers1)
	//验证copy的两个切片是不关联的
	numbers[0] = 99
	numbers1[len(numbers1)-1] = numbers[6]
	printSlice("numbers",numbers)
	printSlice("numbers1",numbers1)
}

func printSlice(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("地址： %p \t len=%d \t cap=%d \t value=%v \n", x, len(x), cap(x), x)
}
