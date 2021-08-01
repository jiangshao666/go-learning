package main

import "fmt"

// 匿名函数

// 一般不把匿名函数用于全局
// var f1 = func(x int) {

// }

func main() {
	// 函数内部不可以什么带名字的函数，因此只能使用匿名函数
	// 你们函数
	sum := func(x int, y int) int {
		return x + y
	}
	fmt.Println(sum(1, 2))

	// 如果只是调用一次的函数，还可以简写成立即执行的函数
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

}
