package main

import "fmt"

// 常量
const PI = 3.14159

//批量申明常量
const (
	statusOK = 200
	notFound = 404
)

//批量声明常量是，如果某一行后面没有赋值，默认和上一行的值一样
const (
	n1 = 100
	n2 //100
	n3 //100
)

//iota， 类似枚举
// iota在const第一次出现的时候初始为0，同一个const域，每新增一行
//常量声明，计数一次
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// 常见的几个iota示例
const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

const (
	c1 = iota //0
	c2 = 100
	c3 //100
	c4 //100
)

const (
	d1 = iota //0
	d2 = 100  // iota +1
	d3 = iota //2
	d4        //3
)

// 多个常量声明在一行

const (
	e1, e2 = iota + 1, iota + 2 // 1 2
	e3, e4 = iota + 1, iota + 2 // 2 3
)

// iota应用，定义数量级

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//PI = 123 //cosnt不可再赋值

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("e1:", e1)
	fmt.Println("e2:", e2)
	fmt.Println("e3:", e3)
	fmt.Println("e4:", e4)
}
