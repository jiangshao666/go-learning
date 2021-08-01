package main

import "fmt"

// 切片
func main() {
	// 切片的定义 变量名 [] 类型
	var s1 []int //定义的切片没有初始化则为nil
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)

	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"海淀", "张江", "光谷"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s1 == nil)

	//长度和容量;
	// 切片是一个指向底层数组的一个引用类型的结构
	//长度是指真实的切片的数量
	// 容量指的是底层数组从切片的第一个元素下标到数组末尾的总的数量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组创建切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //基于数组下标切割，左开右闭
	fmt.Println(s3)
	fmt.Println("len(s3)=", len(s3), "cap(s3)", cap(s3))
	s4 := a1[1:]
	fmt.Println(s4)
	fmt.Println("len(s4)=", len(s4), "cap(s4)", cap(s4))
	s5 := a1[:5]
	fmt.Println(s5)
	fmt.Println("len(s5)=", len(s5), "cap(s5)", cap(s5))
	s6 := a1[:]
	fmt.Println(s6)
	fmt.Println("len(s6)=", len(s6), "cap(s6)", cap(s6))

	//切片再切片
	s8 := s5[1:5]
	fmt.Println(s8)
	fmt.Println("len(s8)=", len(s8), "cap(s8)", cap(s8))

	a1[1] = 100 //修改底层数组的值，则引用的切片的值也会修改
	fmt.Println(s8)

}
