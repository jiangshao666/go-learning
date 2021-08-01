package main

import "fmt"

// 阶乘
func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波拉数列
func fab(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

// 面试题: 一个N级台阶，每次只能走1步或者2步，请问一共有几种走法
func step(n int) uint32 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return step(n-1) + step(n-2)
}

// 递归
func main() {
	// ret := factorial(5)
	// fmt.Println(ret)

	// for i := 1; i < 10; i++ {
	// 	ret := fab(i)
	// 	fmt.Println(ret)
	// }

	stepN := step(4)
	fmt.Println(stepN)
}
