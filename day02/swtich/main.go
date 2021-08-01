package main

import "fmt"

func main() {
	// var caser int = 2
	// switch caser {
	// case 0:
	// 	fmt.Println("大拇指")
	// case 1:
	// 	fmt.Println("食指")
	// case 2:
	// 	fmt.Println("中指")
	// case 3:
	// 	fmt.Println("无名指")
	// case 4:
	// 	fmt.Println("小拇指")
	// default:
	// 	fmt.Println("无效")
	// }

	//变种1
	switch caser2 := 4; caser2 {
	case 0, 2, 4, 6, 8:
		fmt.Println("偶数")
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	default:
		fmt.Println(caser2)
	}

	//变种2
	switch caser3 := 30; {
	case caser3 < 18:
		fmt.Println("未成年")
	case caser3 < 28 && caser3 <= 40:
		fmt.Println("青年")
	case caser3 < 60:
		fmt.Println("中年")
	default:
		fmt.Println("老年")

	}

	//变种3, 兼容C语言中的case
	var s string = "a"

	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough //后续的紧邻的一个case会无条件执行
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	}
}
