//Breack and Continue
package main

import (
	"fmt"
)

/*
 */
func main() {
	//breackContinue()
	//eludeFourGoto()
	primeNumber()

}
func breackContinue() {
	fmt.Println("Break,continued的区别")
	fmt.Println("No.1 Break")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println("No.2 Continue")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	//	打印偶数
	fmt.Println()
	fmt.Println("No.3 打印偶数")
	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	//	去除包含4的数字
	fmt.Println()
	fmt.Println("No.4 剔除包含4的数字")
	num := 0
	for num < 50 {
		num++
		if num%10 == 4 || num/10%10 == 4 {
			continue
		}
		fmt.Printf("%d ", num)
	}

}

func eludeFourGoto() {
	//	goto 控制语句
	fmt.Println()
	fmt.Println("No.5 goto 控制语句")
	num := 0
Loop:
	for num < 50 {
		num++
		if num%10 == 4 || num/10%10 == 4 {
			goto Loop
		}
		fmt.Printf("%d ", num)
	}

}

//输出1~100素数 PrimeNumber
func primeNumber() {
	fmt.Println()
	fmt.Println("No.6 goto 控制语句输出1~100素数")
	num := 0
Loop:
	for num < 100 {
		num++
		//	素数：是只能被1和自身整除的数值
		for i := 2; i < num; i++ {
			if num%i == 0 {
				goto Loop
			}
		}
		fmt.Printf("%d\t ", num)
	}
}
