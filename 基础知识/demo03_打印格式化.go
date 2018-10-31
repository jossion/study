package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	//通用
	str := "stever"
	fmt.Printf("%T,%v\n", str, str)
	p := Point{1, 2}
	fmt.Printf("%T,%v\n", p, p)

	var a rune = '一'
	fmt.Printf("%T,%v\n", a, a)
	//布尔
	fmt.Printf("%T,%t\n", true, true)
	//	整形
	fmt.Printf("%T,%d \n", 123, 123.234) //错误
	fmt.Printf("%T,%5d \n", 123, 123)
	fmt.Printf("%T,%0d \n", 123, 12356)
	fmt.Printf("%T,%b \n", 123, 123)
	rst := fmt.Sprintf("%b,\n", 123)
	fmt.Println(rst)
	fmt.Printf("%x \n", 123) //转换成16进制
	fmt.Printf("%X \n", 123) //转换成16进制
	fmt.Printf("%U \n", 'a') //转换成unicode ASXII
	fmt.Printf("%U \n", '一') //转换成unicode ASXII
	//	浮点
	fmt.Printf("%f \n", 123.1)        //转换成unicode ASXII
	fmt.Printf("%.2f \n", 123.1)      //转换成unicode ASXII
	fmt.Printf("%e \n", 123.1)        //转换成科学计数法
	fmt.Printf("%10e \n", 123.123567) //转换科学计数法
	//	字符串
	fmt.Printf("%s \n", "学习Go语言")
	fmt.Printf("%q \n", "学习Go语言")
	arr:=[3]byte{65,66,67}
	fmt.Printf("%T,%s\n",arr,arr)
	fmt.Printf("%T,%p\n",arr,&arr)
}