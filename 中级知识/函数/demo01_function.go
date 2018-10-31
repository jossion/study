package main

import "fmt"

//声明全局变量
var a int= 7
var b = 9

func main()  {
	c:=sum(a,b)
	fmt.Printf("main()函数中 c = %d\n",c)
	a,b,c :=10,20,0
	fmt.Printf("main()函数中 a = %d\n",a)
	fmt.Printf("main()函数中 b = %d\n",b)
	fmt.Printf("main()函数中 c = %d\n",c)
	c=sum(a,b)
	fmt.Printf("main()函数中 c = %d\n",c)
}

func sum(a,b int)int  {
	return a+b
}