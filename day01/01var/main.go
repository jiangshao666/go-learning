package main

import "fmt"

// GO的变量必须先申明再使用

// var name string
// var age int
// var isOk bool

var studentName string

//批量申明

var (
	name string
	age  int
	isOk bool //全局变量可以申明不使用，用警告
)

func main() {

	// 先声明后使用
	studentName = "hanmeimei"
	name = "jiangshao"
	age = 16

	//var localV string //局部变量申明不使用会报错

	fmt.Printf("name:%s age: %d", name, age)
	fmt.Println()
	fmt.Print(studentName)

	// 申明并同时赋值
	var s1 string = "zhansan"
	fmt.Println(s1)

	// 类型推导
	var id = 10
	fmt.Println(id)

	// 简短变量申明, 常用,只能再函数中使用
	meet := false
	fmt.Println(meet)

	// 匿名变量 _ ,用于占位
	x, _ := foo()

	fmt.Print(x)
}

func foo() (int, string) {
	return 10, "jiangshao"
}
