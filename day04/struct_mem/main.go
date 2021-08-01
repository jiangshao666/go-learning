package main

import "fmt"

// 结构体占用一块连续的内存，需要考虑内存对齐，所以需要注意字段顺序

type s struct {
	a int8
	b int8
	c int8
	d int32
}

func main() {
	ss := s{
		a: int8(10),
		b: int8(20),
		c: int8(30),
		d: int32(40),
	}
	fmt.Printf("%p\n", &ss.a)
	fmt.Printf("%p\n", &ss.b)
	fmt.Printf("%p\n", &ss.c)
	fmt.Printf("%p\n", &ss.d)
}
