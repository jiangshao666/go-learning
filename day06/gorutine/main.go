package main

import (
	"fmt"
	"time"
)

// gorutine, 使用go 关键字启动一个gorutine

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动后会启动一个main gorutine
// main函数如果提前退出，则由main函数启动的gorutine也会退出
func main() {
	for i := 0; i < 5; i++ {
		go hello(i)
	}
	fmt.Println("func 1")
	time.Sleep(time.Second) //防止main gorutine先退出

	// 匿名函数,本质上是一个闭包，会导致upvalue的i应用出错
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("upvalue", i)
		}()
	}
	/* 会导致如下输出
	upvalue 3
	upvalue 3
	upvalue 3
	upvalue 4
	upvalue 5
	*/
	time.Sleep(time.Second) //防止main gorutine先退出

	// 解决上述问题的方案，将upvalue作为参数传递进匿名函数
	for i := 0; i < 5; i++ {
		go func(upvalue int) {
			fmt.Println("argvalue", upvalue)
		}(i)
	}
	time.Sleep(time.Second) //防止main gorutine先退出
}
