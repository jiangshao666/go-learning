package main

import "fmt"

// 运算符
func main() {
	var (
		a = 4
		b = 5
	)

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ //++和--是单独的语句,不可以作为右值付给其他变量， 类似 a= a++
	b--

	a += 1
	b -= 1
	fmt.Println(a)
	fmt.Println(b)

	// 关系运算符
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)

	// 逻辑运算符
	age := 30

	if age > 18 && age <= 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("不用上班了")
	}

	if age < 18 || age > 60 {
		fmt.Println("不用上班了")
	} else {
		fmt.Println("苦逼上班")
	}

	if !(age > 18 && age < 60) {
		fmt.Println("不用上班了")
	} else {
		fmt.Println("苦逼上班")
	}

	// 位运算
	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 2)
	fmt.Println(5 >> 2)

}
