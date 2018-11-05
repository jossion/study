package main

import "fmt"

func main() {
a := 10
fmt.Println("函数调用前的a的值",a)

b:=&a
change(b)
fmt.Println("函数调用后a的值",a)
}

func change(num *int) {
	*num = 20
}
