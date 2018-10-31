/*结构体匿名字段：
	（1）概念
			·结构体中的字段没有名字，只包含一个没有字段名的类型，这些字段被称为你名字段。
			·如果字段没有名字，那么默认使用类型名作为字段名
			·注意：同一个类型只能写一个
			·结构体潜逃中采用匿名结构体字段可以模拟集成关系
	（2）示例代码
			type User struct{
				string
				byte
				int8
				float64
			}
*/

package main

import "fmt"

type noNameField struct {
	string
	byte
	int8
	float64
}
//匿名字段局限性：不易阅读，
func main() {
	user:=noNameField{"steven",'m',35,178.5 }
	fmt.Println(user)
	fmt.Printf("名字：%s \t性别：%c\t 年龄：%v 身高：%.2f \n",user.string,user.byte,user.int8,user.float64)
}
