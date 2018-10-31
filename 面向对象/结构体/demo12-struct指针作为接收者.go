package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

func main() {
	r1 := Rectangle{5, 8}
	r2 := r1
	fmt.Println("-----------原始定义的R1&R2----------")
	fmt.Printf("R1的内存地址：\t%p\t",&r1)
	fmt.Printf("R2的内存地址：\t%p\t\n",&r2)
	fmt.Printf("R1的高度：%.2f\t\n", r1.height)
	fmt.Printf("R2的高度：%.2f\t\n", r2.height)

	r1.setValue()
	//r2.setValueP()
	fmt.Println("------------setvalue后-------------")
	fmt.Printf("R1的高度：%.2f\t\n", r1.height)
	fmt.Printf("R2的高度：%.2f\t\n", r2.height)
	fmt.Println("------------------------------------")
	r1.setValueP()
	//r2.setValueP()
	fmt.Println("------------setvalueP后-------------")
	fmt.Printf("R1的高度：%.2f\t\n", r1.height)
	fmt.Printf("R2的高度：%.2f\t\n", r2.height)
	fmt.Println("------------------------------------")
}
func (r Rectangle) setValue() {
	fmt.Printf("setValue方法中r内存地址：\t%p\t\n",&r)
	r.height = 10
}
func (r *Rectangle) setValueP() {
	fmt.Printf("setValueP方法中r内存地址：\t%p\t\n",r)
	r.height = 20
}
