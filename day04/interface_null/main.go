package main

import "fmt"

// 空接口 type xx interface{} 空接口没有必要起名字
// 空接口表示所有的结构体都默认实现了该接口， 也就是任意类型变量都可以赋值给空接口
// 使用的例子，如fmt.Println的参数就是一个空接口, 可以包含各种类型的map等

// interface是关键字
// interface{}是空接口，是类型

func show(a interface{}) {
	fmt.Printf("type: %T value: %v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 100)
	m1["name"] = "lisi"
	m1["age"] = 20
	m1["married"] = true
	m1["hobby"] = [...]string{"eat", "sleep"}

	fmt.Println(m1)

	show(true)
	show(false)
	show(1)
	show(m1)
}
