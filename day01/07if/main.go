package main

import "fmt"

func main() {

	// 方式一
	// age := 19

	// if age > 18 {
	// 	fmt.Println("成年了")
	// } else {
	// 	fmt.Println("好好学习")
	// }

	//方式二

	if age := 15; age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Printf("才%d岁好好学习", age)
	}
	//fmt.Println(age)
}
