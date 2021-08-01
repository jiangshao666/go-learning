package main

import "fmt"

// 变量作用域

var name = "lisi" //全局变量,全局可访问

// 变量查找顺序: 首先查找块内，如果没有则往上查找函数内局部变量，再没有则往上查找全局变量
func f() {
	fmt.Println(name)

	// x为局部变量，只能再函数内使用
	x := 10
	fmt.Println(x)

	// age为块作用域内变量，只能再if块内使用
	if age := 10; age < 18 {
		fmt.Println("好好学习")
	}

	sum := 0
	//i为块作用域内变量，只能再for块内使用
	for i := 0; i < 10; i++ {
		sum += i
	}
}

func main() {
	f()
}
