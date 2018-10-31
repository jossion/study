package main

import (
	"fmt"
	"time"
)

func main() {
	t1:=time.Now()
	testTimePakage()
	t2:=time.Now()
	fmt.Printf("程序完成用时：%v ms",t2.Sub(t1).Seconds()*1000)
}
func testTimePakage() {
	//	1.Now()返回当前本地时间
	t := time.Now()
	fmt.Println("1.", t)
	//2.Local（）将时间转成本地时区，但指向同一时间点的time
	fmt.Println("2.", t.Local())
	//3.UTC()将时间转成UTC和零时区，待指向同一时间点的time
	fmt.Println("3.", t.UTC())
	//	4.Date()可以根据指定数字，返回一个本地货国际标准的时间格式
	t = time.Date(2018, time.January, 1, 1, 1, 1, 0, time.Local)
	fmt.Printf("4. 本地时间%s，国际统一时间：%s \n", t, t.UTC())
	//	5.Parse()能将一个格式化的时间字符串解析成它所代表的时间。就是string转time
	/*
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	*/
	t, _ = time.Parse("2006-01-02 15:04:05", "2018-10-24 15:31:40")
	fmt.Println("5.", t)
	//	6.Format（）根据指定的时间格式，将时间格式化成文本，就是time-》string
	fmt.Println("6.", time.Now().Format("2006-01-02 15:04:05"))
	//	7.String()将时间格式化层字符串，格式为：“2006-01-02 15:04:05.999999999 -0700 MST”
	fmt.Println("7.", time.Now().String())
	//	8.Unix()将t 表示为Unix时间（时间戳，一个int64），即从时间点January1,1970 UTC到时间点t就过得时间（单位秒）
	fmt.Println("8.", time.Now().Unix())
	//	9.UnixNano()将t表示为unix时间（时间戳，int64），即从时间点January1 1970到时间点t说经过的时间（单位：纳秒）
	fmt.Println("9.", time.Now().UnixNano())
	//	10.Equal()判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息
	fmt.Println("10-1.", t.Equal(time.Now()))
	fmt.Println("10-2.", time.Now().Equal(time.Now()))
	//	11.Before()如果t代表的时间在u之前，返回真；否则为假
	// 		and After() 如果t代表的时间在u之后，返回真；否则为假
	fmt.Println("11-Before:", t.Before(time.Now()))
	fmt.Println("11-After:", t.After(time.Now()))
	//	12.Date()返回时间点t对应的年月日
	year, month, day := time.Now().Date()
	fmt.Println("12.", year, month, day)
	//	13.year() 返回t对应的年
	year = time.Now().Year()
	fmt.Println("13.", year)
	//14.month() 返回t 对应的月
	month = time.Now().Month()
	fmt.Println("14.", month)
	//	15.day()返回t对应的日
	day = time.Now().Day()
	fmt.Println("15.", day)
	//	16.weekday()返回t对应的星期几
	week := time.Now().Weekday()
	fmt.Println("16.", week)
	//	17.clock()返回t对应的时分秒
	h, m, s := time.Now().Clock()
	fmt.Println("17.", h, m, s)
	//	18.hour()Minute()Second()Nanosecond()
	h = time.Now().Hour()
	m = time.Now().Minute()
	s = time.Now().Second()
	ms := time.Now().Nanosecond()
	fmt.Printf("18.当前时间为:\t %d时\t%d分\t%d秒\t%d毫秒 ", h, m, s, ms)
	fmt.Println()
	fmt.Println("+++++++++++++++++++计算时间差值+++++++++++++++++++")
	//	19.sub() 返回一个时间段t-u。如果结果超出了Duration可以表示的最大/最小值，将返回最大或最小值，要回去时间点t-d（d 为Duration），可以使用t.ADD(-d)
	fmt.Println("19.", t.Sub(time.Now()))
	//	20.hours()menutes()seconds()Nanoseconds() 计算时间差值
	fmt.Println("20-h", t.Sub(time.Now()).Hours())
	fmt.Println("20-m", t.Sub(time.Now()).Minutes())
	fmt.Println("20-s", t.Sub(time.Now()).Seconds())
	fmt.Println("20-ns", t.Sub(time.Now()).Nanoseconds())
	fmt.Println("+++++++++++++++++++计算时间差值结束+++++++++++++++++++")

	//	21.string()返回时间段餐用的格式，用字符串
	fmt.Println("21.", "时间间距", t.Sub(time.Now()).String())
	//	22.ParseDuration 解析一个时间段字符串
	d, _ := time.ParseDuration("1h30m20s")
	fmt.Println("22.", d)
	//	23.add（）返回t+d
	fmt.Println("23", "交卷时间", time.Now().Add(d))
	//	24.AddDate()返回增加了给出的年份、月份和天数的时间点Time
	fmt.Println("24", "一年一月一天后的日期", time.Now().AddDate(0, 3, 0))

}
