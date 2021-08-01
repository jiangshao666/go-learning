package main

import "fmt"

// 数组
func main() {

	// 语法 变量名 [长度]类型
	// go种数值的长度是数组类型的组成部分,
	// 如果不初始化，数组的默认值是对应类型的零值： bool是false, 字符串是"", 整数和浮点数是0
	var a1 [3]int
	var a2 [4]int

	fmt.Printf("%v %T\n", a1, a1)
	fmt.Printf("%v %T\n", a2, a2)

	// 数组初始化1
	a1 = [3]int{1, 2, 3}
	fmt.Println(a1)
	// 初始化方式2, 根据数据元素个数推导数组长度
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a3)
	// 初始化方式3， 初始化部分值，剩余部分回补齐默认值
	a10 := [10]int{0, 1}
	fmt.Println(a10)
	// 初始化方式4, 通过指定下标对应值初始化
	a4 := [5]int{0: 1, 4: 5}
	fmt.Println(a4)

	// 数组遍历
	// 方式1， 通过索引便利
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	// 方式2， for range
	for i, v := range a1 {
		fmt.Printf("%d=%d\n", i, v)
	}

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(a11)

	//多维数组遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型，不可以修改
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100 //修改值不影响原来的值
	fmt.Println(b1)
	fmt.Println(b2)

	// 练习1， 数组求和 [1,3,5,7,8]
	sa1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for i := 0; i < len(sa1); i++ {
		sum += sa1[i]
	}
	fmt.Println(sum)
	// 练习2，找出数组中两个数的和为给定数值N的下标 i, j
	for i := 0; i < len(sa1); i++ {
		for j := i + 1; j < len(sa1); j++ {
			if sa1[i]+sa1[j] == 8 {
				fmt.Printf("i=%d j=%d\n", i, j)
			}
		}
	}
}
