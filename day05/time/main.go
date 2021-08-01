package main

import (
	"fmt"
	"time"
)

// 时间相关的基础操作
func timeBaseFunc() {
	// .Now默认返回的是本地时间
	now := time.Now()

	fmt.Println(now.Year(), now.Month(), now.Day(), now.Weekday(),
		now.Hour(), now.Minute(), now.Second())
	// 时间戳
	fmt.Println(now.Unix())

	// 格式化
	s := now.Format("2006-01-02 03:04:05.000")
	fmt.Println(s)

	// 将时间戳转换成Time
	t := time.Unix(1624505775, 0)
	fmt.Println(t)

	// 讲字符串的时间转成time, 默认parse返回的是UTC时间
	s2t, err := time.Parse("2006-01-02 03:04:05.000", "2021-06-24 11:39:15.510")
	if err != nil {
		fmt.Printf("parse time failed, err:%v", err)
		return
	}
	fmt.Println(s2t)

	// 时间的加减, 参数是time.Duration
	nt := now.Add(time.Hour)
	fmt.Println(nt)

	// 返回时区信息
	location := now.Location()
	fmt.Println(location)
	location = s2t.Location()
	fmt.Println(location)

	// 返回UTC时间
	utcnow := now.UTC()
	fmt.Println(utcnow)

	// 采用指定时区返回时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("loadlocation failed, err:%v", err)
		return
	}
	s2tlocal, err := time.ParseInLocation("2006-01-02 03:04:05.000", "2021-06-24 11:39:15.510", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v", err)
		return
	}
	fmt.Println(s2tlocal)
}

// 时间相关的进阶操作
func timeStepinFunc() {
	// 睡眠, 阻塞当前go程
	fmt.Println(time.Now())
	time.Sleep(time.Second * 2)
	fmt.Println(time.Now())

	// 定时器
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v \n", now)
	}
}

func main() {
	//timeBaseFunc()
	timeStepinFunc()
}
