package main

import (
	"fmt"
	"strings"
)

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 跟add的差别在于 x是共享引用的，每次调用都会累加x
func add2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// f := add()
	// fmt.Println(f(10)) //10
	// fmt.Println(f(20)) //20

	f := add2(0)
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30

	jpgFunc := makeSuffix(".jpg")
	pdfFunc := makeSuffix(".pdf")

	fmt.Println(jpgFunc("beauty"))
	fmt.Println(pdfFunc("golang"))

	fAdd, fSub := calc(10)
	fmt.Println(fAdd(1), fSub(2))
	fmt.Println(fAdd(3), fSub(4))
	fmt.Println(fAdd(5), fSub(6))
}
