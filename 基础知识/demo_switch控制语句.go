package main

import "fmt"

func main() {
	//ScoreGrade()
	//ScoreGrade1()
	Month()
}

//判断成绩
func ScoreGrade() {
	score := 90.5
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("你的等级是：%s \n", grade)
	fmt.Print("你的评价是：")
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "":
		fmt.Println("良好")
	case "C":
		fmt.Println("中等")
	case "D":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}
func ScoreGrade1() {
	score := 78.5
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("你的等级是：%s \n", grade)
	fmt.Print("你的评价是：")
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("中等")
	case "D":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}

//判断月份的天数

func Month() {
	//	定义年 月 日
	year := 2018
	month := 2
	day := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
	case 4, 6, 9, 11:
		day = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			day = 29
		} else {
			day = 28
		}

	default:
		day = -1
	}
	fmt.Printf("%d年%d月的天数为：%d", year, month, day)
}
