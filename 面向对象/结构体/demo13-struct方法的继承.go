package main

import "fmt"

type Human struct {
	name, phone string
	age      int
}
type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func main() {
s1:=Student{Human{"Daniel","138000001",12},"北大附中"}
e1:=Employee{Human{"Stever","1788883388",35},"北航学院"}
s1.SayHi()
fmt.Printf("我是%s的人",s1.school)
fmt.Println()
e1.SayHi()
	fmt.Printf("我是%s的人",e1.company)
}
func (h *Human)SayHi()  {
	fmt.Printf("大家好！我是%s，我%d岁，联系方式:%s,  ",h.name,h.age,h.phone)
}