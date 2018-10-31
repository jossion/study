package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randTest()
}

func randTest() {
	//	1.通过默认的随机数种子获取
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Float64())
	//2.动态种子：先生成随机资源，实例化成对象，

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(10)
	fmt.Println(randnum)
//	3.简写形式：通过动态种子
//0-10的随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	//[5,11]
	num:=rand.Intn(11-5+1)+5
	fmt.Println(num)
	}
