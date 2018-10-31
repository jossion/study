package main

import "fmt"

func main() {
addcontroller()
}

//实现的累加器闭包函数
func adder()func(int)int{
	sum:=0
	res:=func(num int)int{
		sum+=num
		return sum
	}
	return res
}

func addcontroller()  {
	res:=adder()
	fmt.Printf("%T \n",res)
	for i:=0;i<=5 ;i++  {
		fmt.Printf("i=%d \t",i)
		fmt.Println(res(i))
	}
}