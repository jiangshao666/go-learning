package main

import (
	"fmt"
	"strconv"
)

// strconv标准库，用于字符串和其他类型的转换

func main() {
	// 其他类型转字符串, 可以通过Sprinf格式化实现
	// str := string(int) //强转的话是把int的值当初unicode值转对应的符号
	var b bool = false
	s := fmt.Sprintf("%t", b)
	fmt.Printf("%#v \n", s)
	// 整型转字符串的方式
	s = strconv.Itoa(100)

	fmt.Printf("%#v \n", s)
	//================================
	// 字符串转其他 使用strconv.ParseXXX系列
	str := "1000"
	//i := int32(str) //不可以直接强转
	// 通过strconv库实现
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("parse int failed")
	}
	fmt.Printf("%T \n", i)
	// 如果是转换成整数，可以使用strconv.Atoi
	intv, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("parse Atoi failed")
	}
	fmt.Printf("%T \n", intv)
}
