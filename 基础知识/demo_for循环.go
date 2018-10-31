package main

import "fmt"

func main()  {
	//baseFor()
	//baseFor1()
	//baseFor2()
	//baseFor3()
	//TraverseString()

}
//基本形态
func baseFor()  {
	fmt.Println("\n-----------------------")
	for i:=0;i<10 ;i++  {
		fmt.Printf("%d ",i)

	}
	fmt.Println("\n------------------------")
}

func baseFor1()  {
	fmt.Println("\n------------------------")
	i:=0
	for ;i<10 ;i++  {
		fmt.Printf("%d ",i)

	}
	fmt.Println("\n------------------------")
}
func baseFor2()  {
	fmt.Println("\n-------------------------")
	i:=0
	for ; ;i++  {
		if i>=10{
			break
		}
		fmt.Printf("%d ",i)
	}
	fmt.Println("\n-------------------------")
}
func baseFor3()  {
	fmt.Println("\n-------------------------")
	i:=0

	for ; ; {
		if i>=10{
			break
		}
		i++
		fmt.Printf("%d ",i)
	}
	fmt.Println("\n-------------------------")
}

//For 关键字后没有 条件表达式

func TraverseString()  {
	str:="username=Liujihe"
	for i,value:=range str  {
	fmt.Printf("第%d位的字符值是：%v,字符是%c \n",i,value,value)
	}
}
