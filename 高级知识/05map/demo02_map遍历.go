package main

import "fmt"

func main() {
	//1.定义时同时初始化
	var county = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Roma",
	}
	fmt.Println(county)
	//2.短变量声明初始化方式
	rating := map[string]float64{"C": 5, "Go": 4.5, "Python": 4.6, "C++": 3}
	fmt.Println(rating)
	//	创建map后再赋值
	countyMap := make(map[string]string)
	countyMap["China"] = "BeiJing"
	countyMap["Japan"] = "Tokyo"
	countyMap["India"] = "New Delhi"
	countyMap["France"] = "Paris"
	fmt.Println(countyMap)


//	遍历的三种方式
//1.普通遍历
	for key, value := range county {
	fmt.Println("国家",key,"首都",value)
	}
//2.只遍历值
	for key, value := range rating {
		fmt.Println("语言",key,"评分",value)
	}
//3.只遍历key
	for key := range countyMap {
		fmt.Println("国家",key,"首都")
	}

	/*
	查看元素在集合中是否存在：
		1.通过key获取map中对应的value；语法：map[key]
		2.当key不存在时，会得到value值类型的默认值，比如string类型得到空字符串，int类型得到0，但是程序不会报错
		3.所以可以通过value，ok:=map[key],获取key/value是否存在，ok是bool型，如果ok=true，则该键值对存在，否则不存在。
	 */
	 value:=countyMap["England"]
	 fmt.Printf("%q \n",value)
	 value,ok:=countyMap["England"]
	 fmt.Printf("%T,%v \n",ok,ok)
	 if ok{
	 	fmt.Println("首都",value)
	 }else{
	 	fmt.Println("首都信息不存在")
	 }
	 if value,ok:=countyMap["China"];ok{

		 fmt.Println("首都",value)
	 }else{
		 fmt.Println("首都信息不存在")
	 }
}
