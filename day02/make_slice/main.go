package main

import "fmt"

func main() {
	// make创建切片 make([]类型，长度 容量)

	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len=%d cap=%d\n", s2, len(s2), cap(s2))

	// 切片之间是不能直接用==来比较的，唯一可以==比较是跟nil比较
	// 一个nil值的切片的长度和容量都是0, 但是不能说长度和容量都是0的切片一定是nil
	// 因此如果要判断一个切片是不是空的，应该用len来判断
	var snil []int
	fmt.Printf("snil=%v len=%d cap=%d isnil=%t\n", snil, len(snil), cap(snil), (snil == nil))
	semtpy := []int{}
	fmt.Printf("semtpy=%v len=%d cap=%d isnil=%t\n", semtpy, len(semtpy), cap(semtpy), (semtpy == nil))
	s0 := make([]int, 0)
	fmt.Printf("s0=%v len=%d cap=%d isnil=%t\n", s0, len(s0), cap(s0), (s0 == nil))

	// 切片的赋值拷贝，实际上是引用赋值,底层数组是同一个
	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 100
	fmt.Println(s3, s4)

	// 切片的遍历，同数组
}
