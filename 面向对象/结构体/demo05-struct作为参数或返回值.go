package main

import "fmt"

type Flower struct {
	name  string
	color string
}

func main() {
	//	1.结构体作为参数
	f1 := Flower{"玫瑰", "红"}
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1) //main.Flower,{玫瑰 红},0xc000044400
	//结构体对象作为参数
	changeFlower1(f1)
	//结构体指针作为对象
	changeFlower2(&f1)

	//	2.结构体作为返回值
	//对象作为返回值
	f2:=getFlower1()
	f3:=getFlower1()
	fmt.Println("更改前",f2,f3)
	f2.name="杏花"
	fmt.Println("更改后",f2,f3)
	fmt.Println("------------------------------")
		//结构体指针作为返回值
	f4:=getFlower2()
	f5:=getFlower2()
	fmt.Println("更改前",f4,f5)
	f4.name="桃花"
	fmt.Println("更改后",f4,f5)
	fmt.Println("------------------------------")
}
func changeFlower1(f Flower) {
	f.name = "月季"
	f.color = "粉色"
	fmt.Printf("函数changeFlower1内f:%T,%v,%p \n", f, f, &f) //main.Flower,{月季 粉色},0xc000044480
}
func changeFlower2(f *Flower) {
	f.name = "百合花"
	f.color = "白色"
	fmt.Printf("函数changeFlower2内f:%T,%v,%p ，%p\n", f, f, f, &f) //*main.Flower,&{百合花 白色},0xc000044400
}
func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白色"}
	fmt.Printf("函数getFlower1内的F：%T,%v,%p\n", f, f, &f)
	return
}
func getFlower2() (f *Flower) {
	f = &Flower{"芙蓉", "红色"}
	fmt.Printf("函数getFlower2内的F：%T,%v,%p，%p\n", f, f,f, &f)
	return
}

/*面向对象编程时要注意：返回值尽量不要返回对象（以免过多的性能资源损耗，）尽可能的返回对象指针*/