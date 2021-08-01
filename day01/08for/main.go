package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种1
	// var i = 0
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//变种2

	// var i = 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 无限循环

	// for {

	// }

	// 跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break
			continue
		}
		fmt.Println(i)
	}
}
