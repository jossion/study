package main

import "fmt"

func main() {
	eval()
}

func eval() {
	num1, num2, result := 12, 4, 0
	operation := "+"
	switch operation {
	case "+":
		result = num1 + num2
		//fallthrough
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	default:
		result = -1
	}
	fmt.Println(result)
}
