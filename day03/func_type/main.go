package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 1
}

//函数可以作为参数
func f3(f func() int) {
	ret := f()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数也可以作为返回值

func ff(a, b int) int {
	return a + b
}

func f5(x func() int) func(int, int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(b)
	//f3(f4) 错误 f4的类型和f3所需参数类型不一样

	fret := f5(f2)
	fmt.Printf("%T\n", fret)
}
