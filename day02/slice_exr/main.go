package main

import "fmt"

// slice是引用类型，参数实际上是传的引用地址
func test(s []int) {
	s[0] = 100
}

// 切片练习, 请说出下述打印信息
func main() {
	var a = make([]int, 5, 10)
	//
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a) //[0,0,0,0,0,0,1,2,3,4,5,6,7,8,9]

	test(a)
	fmt.Println(a) //[100,0,0,0,0,0,1,2,3,4,5,6,7,8,9]

}
