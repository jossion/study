package main

import "fmt"

type Human struct {
	name, phone string
	age         int
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
	s1 := Student{Human{"Daniel", "138000001", 12}, "北大附中"}
	e1 := Employee{Human{"Stever", "1788883388", 35}, "北航学院"}
	s1.SayHi()
	fmt.Println("---------------------------------------")
	e1.SayHi()
}
func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s，我%d岁，联系方式:%s,  ", h.name, h.age, h.phone)
}
func (s *Student) SayHi() {
	fmt.Printf("大家好！我是%s，我在%s上学,我%d岁，联系方式:%s,  ", s.name, s.school, s.age, s.phone)
}
func (e *Employee) SayHi() {
	fmt.Printf("大家好！我是%s，我在%s工作,我%d岁，联系方式:%s,  ", e.name, e.company, e.age, e.phone)
}
