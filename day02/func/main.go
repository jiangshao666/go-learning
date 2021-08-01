package main

import "fmt"

// 函数

// go语言中函数的格式: func 函数名(参数列表) (返回值列表) {}

func sum(x int, y int) int {
	return x + y
}

// 无参无返回值
func f1() {
	fmt.Println("no arg no ret")
}

// 有参无返回值
func f2(x int) {
	fmt.Println(x)
}

// 无参有返回值
func f3() int {
	return 0
}

// 返回值有命名，则return 后面可以不带ret
func f4(x int, y int) (ret int) {
	ret = x + y
	//return ret
	return
}

// 返回值可以命名也可以不命名
func f5(x int, y int) int {
	return x + y
}

// 返回值可以有多个
func f6(x int) (succ bool, value int) {
	return true, 1
}

// 参数和返回值中连续参数是相同类型，则除去最后一个外，前面的参数类型可以省略
func f7(x, y, z int, s1, s2 string, b1, b2 bool) (sum, div int) {
	return x + y, x / y
}

// 可变长参数, 注意必须放在最后
// 可变长参数本质上是一个slice
func f8(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

func main() {
	f8("name", 1, 2, 3)
}
