package main

import (
	"fmt"
)

func main() {
	//math.MaxFloat32

	f1 := 1.2345 //默认是float64, 如果需要float32则明确指定
	f2 := float32(1.2345)

	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	//f1 = f2 float32的值不可以赋值给float64，是两种不同的类型

	var ret1 float32 = 10 / 4
	fmt.Println(ret1)

	var ret2 float32 = float32(10) / 4
	fmt.Println(ret2)
}
