package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
width,height float64
}
type Circle struct {
radius float64
}
func main() {
	r1:=Rectangle{10,4}
	r2:=Rectangle{12,5}
	c1:=Circle{2}
	c2:=Circle{10}
	fmt.Printf("R1的面积：%.2f,\t%+v \t\n",r1.Area(),r1)
	fmt.Printf("R2的面积：%.2f,\t%+v \t\n",r2.Area(),r2)
	fmt.Printf("C1的面积：%.2f,\t%+v \t\n",c1.Area(),c1)
	fmt.Printf("C2的面积：%.2f,\t%+v \t\n",c2.Area(),c2)
}

//定义Rectangle的方法
func(r Rectangle) Area()float64  {
	return r.width*r.height
}
//定义Circle的方法
func(c Circle) Area()float64  {
	return  c.radius*c.radius*math.Pi
}