package main

import (
	"fmt"
	"math"
)

//形状接口 包括：矩形、圆形、三角形
//1.定义接口
type Shape interface{
	perimeter() float64
	area() float64
}
//2.定义结构体
type Rectangle struct {
	a,b float64
}
type Circel struct {
	radius float64
}
type Triangle struct {
a,b,c float64
}
//实现perrmeter和area方法
func (r Rectangle)perimeter()float64  {
	return (r.a+r.b)*2
}
func (r Rectangle)area() float64 {
return r.a*r.b
}
func (c Circel)perimeter()  float64{
return 2*c.radius*math.Pi
}
func (c Circel)area()float64  {
	return c.radius*c.radius*math.Pi
}
func (t Triangle)perimeter() float64 {
return t.a+t.b+t.c
}
func (t Triangle)area() float64{
//海伦公司
p:=t.perimeter()/2//半周长
return math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
}
func main() {
//	创建形状
//矩形
r1:=Rectangle{10.4,20.3}
c1:=Circel{5.5}
t1:=Triangle{3,4,5}
show(r1)
show(c1)
show(t1)
}

func show(s Shape)  {
	fmt.Printf("周长：%.2f \t面积：%2.f\t",s.perimeter(),s.area())
}