package main

import "fmt"

func main() {
	a:= [...] string {"a","b","c","d"}
	b:=a
	b[0]="x"
	fmt.Println("a",a)//[a b c d ]
	fmt.Println("b",b)//[x b c d ]
}
//