package main

import "fmt"

var a = 21.0
var b = 5.0
var c float64

func main() {
	Arithmetic()
	Relational()
	Logical()
	Bitwise()
	Assignment()
}

//运算符Arithmetic
func Arithmetic() {
	c = a + b
	fmt.Printf("第一行 - C的值为：%.2f\n", c)
	c = a - b
	fmt.Printf("第二行 - C的值为：%.2f\n", c)
	c = a * b
	fmt.Printf("第三行 - C的值为：%.2f\n", c)
	c = a / b
	fmt.Printf("第四行 - C的值为：%.2f\n", c)
	//c = a % b
	d := fmt.Sprintf("取模值：%d\n", int(a)%int(b))
	fmt.Printf("第五行 - C的值为： %d\n", int(a)%int(b))
	fmt.Printf("第五行 - C的值为： %s\n", d)
	a++
	fmt.Printf("第六行 -A的值为：%.2f\n", a)
	a = 21
	a--
	fmt.Printf("第七行 -A的值为：%.2f\n", a)
}

//关系运算符Relational
func Relational() {
	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}
	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	} else {
		fmt.Printf("第四行 - b小于等于 a\n")
	}
	if a >= b {
		fmt.Printf("第五行 - a 大于等于 b\n")
	} else {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}

}

//逻辑运算符Logical Operator
func Logical() {
	a := false
	b := true
	if a && b {
		fmt.Printf("第一行逻辑值为：true\n")
	} else {
		fmt.Printf("第一行逻辑值为：false\n")
	}
	if a || b {
		fmt.Printf("第二行逻辑值为：true\n")
	} else {
		fmt.Printf("第二行逻辑值为：false\n")
	}
}

//位运算符 Bitwise operator
//包括：按位与（&），按位或（|），按位异或（^）,按位左移（<<）,按位右移（>>）
func Bitwise() {
	fmt.Println(252 & 63)
	fmt.Println(178 | 94)
	fmt.Println(20 ^ 5)
	fmt.Println(3 << 4)
	fmt.Println(11 >> 2)
}

//赋值运算符 Assignment operator包括：（=、 +=、 -=、*=、/=、%=、<<=、>=、&=、^=、|=、）
func Assignment() {
	c = a
	fmt.Printf("第1个C等于：%f\n", c)
	c += a
	fmt.Printf("第2个C等于：%f\n", c)
	c -= a
	fmt.Printf("第3个C等于：%f\n", c)
	c *= a
	fmt.Printf("第4个C等于：%f\n", c)
	c /= a
	fmt.Printf("第5个C等于：%f\n", c)
	d := 200
	d <<= 2
	fmt.Printf("第1个d= %d\n", d)
	d >>= 2
	fmt.Printf("第2个d= %d\n", d)
	d &= 2
	fmt.Printf("第3个d= %d\n", d)
	d |= 2
	fmt.Printf("第4个d= %d\n", d)
	d ^= 2
	fmt.Printf("第5个d= %d\n", d)
}

//指针运算

