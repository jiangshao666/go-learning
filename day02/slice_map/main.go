package main

import "fmt"

// slice和map的结合
func main() {

	//01 元素类型为map的slice
	// 定义一个元素类型为map的切片，但是切片长度为0, 容量为10
	//var s1 = make([]map[int]string, 0, 10)
	//s1[0][10] = "beijing" //直接赋值会报index out of range，原因是切片长度为0，则s[0]是nil
	var s1 = make([]map[int]string, 1, 10)
	//s1[0][10] = "beijing" //修改s1长度为1以后，会报assignment to entry in nil map， 原因是s1[0]的map没有初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "beijing"
	fmt.Println(s1)

	//02 元素类型为slice的map
	var m1 = make(map[string][]int, 10)
	m1["beijing"] = []int{10, 20, 30}
	fmt.Println(m1)
}
