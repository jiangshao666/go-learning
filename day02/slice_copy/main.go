package main

import "fmt"

// 切片 copy, 本质上是重新创建的一个独立的地城数组
func main() {
	s1 := []int{1, 3, 5}
	s2 := s1
	var s3 = make([]int, 3)
	copy(s3, s1)
	s1[0] = 100
	fmt.Println(s1, s2, s3)

	// 只copy所需部分，剩余部分为0值
	var s4 = make([]int, 10)
	copy(s4, s1)
	fmt.Println(s4)

	// 拷贝是空间不足，但是不会出错，只拷贝能够拷贝的部分
	var s5 = make([]int, 1)
	copy(s5, s1)
	fmt.Println(s5)

}
