package main

import "fmt"

type ISayHello interface {
	SayHello() string
}
type Duck struct {
	name string
}
type Person struct {
	name string
}
type Bird struct {
	name string
}

func (d Duck) SayHello() string {
	return d.name + "我是大黄鸭，GaGaGa！"
}
func (p Person) SayHello() string {
	return p.name + "我是你哥，请叫---大爷"
}
func (b Bird) SayHello() string {
	return b.name + "我不是鸡，我是鸟！！！"
}
func main() {
	duck := Duck{"鸭子:"}
	person := Person{"习大大:"}
	bird := Bird{"鹰:"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println(bird.SayHello())
	fmt.Println("----------------------------")

	var i ISayHello
	i = duck
	fmt.Println(i.SayHello())
	i = person
	fmt.Println(i.SayHello())

}
