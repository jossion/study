/*
1.概念：
	·事物的多种形态
	·Go中的多态性是在接口的帮助下实现的。定义接口类型，创建实现该接口的结构体对象
	·定义接口类型的对象，可以保存实现该接口的任何类型的值。Go语言接口变量的这个特性实现了Go语言中的多态
	·接口类型的对象，不能访问其实现类中的属性字段 
*/

package main

import "fmt"

type Income interface {
	//计算收入总额
	calculate() float64 //计算总额
	source() string     //收入来源
	//收入来源
}
type FixedBill struct {
	projectName  string  //项目名称
	beddedAmount float64 //项目金额
}

//定时产生的收入
type TimeAndMaterial struct {
	projectName string
	workHours   float64 //时常
	hourlyRate  float64 //小时收入
}

//固定收入项目收入
func (f FixedBill) calculate() float64 {
	return f.beddedAmount
}
func (f FixedBill) source() string {
	return f.projectName
}

//定时项目收入
func (t TimeAndMaterial) calculate() float64 {
	return t.workHours * t.hourlyRate
}
func (t TimeAndMaterial) source() string {
	return t.projectName
}

//广告收入
type Advertisement struct {
	adName         string
	clickCount     int
	incomePerClick float64
}

func (a Advertisement) calculate() float64 {
	return float64(a.clickCount) * a.incomePerClick
}
func (a Advertisement) source() string {
	return a.adName
}
func main() {
	p1 := FixedBill{"proj1", 5000}
	p2 := FixedBill{"proj2", 5000}
	p3 := TimeAndMaterial{"proj3", 100, 40}
	p4 := TimeAndMaterial{"proj4", 20, 100}
	p5 := Advertisement{"广告1", 100, 30}
	p6 := Advertisement{"广告2", 300, 80}
	ic := []Income{p1, p2, p3, p4, p5, p6}
	fmt.Println("收入总额：", calculateNetIncome(ic))
}

//计算总收入
func calculateNetIncome(ic []Income) float64 {
	netincome := 0.0
	for _, income := range ic {
		fmt.Printf("收入来源：%s，收入金额%.2f \n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	return netincome
}
