package main

import "fmt"

// 整型

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)

	// 八进制

	i2 := 0777
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0xf
	fmt.Printf("%d\n", i3)

	// 查看变量类型
	fmt.Printf("%T\n", i3)

	// 声明int8类型, 明确指定类型，否则默认为int
	i4 := int8(8)
	fmt.Printf("%T\n", i4)
}
