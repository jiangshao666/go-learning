package main

import "fmt"

//panic和recover
// panic可以在任何地方调用，用于排除错误中断程序继续执行
// recover必须在defer里面才能执行，用于捕获错误，以及部分资源释放

func funcA() {
	fmt.Println("A")
}

// 注意 defer必须在panic之前定义
func funcB(x int) {
	defer func() {
		err := recover()
		fmt.Println(err)
		//同时用于释放资源
	}()
	//panic("发生严重错误")

	f := 1 / (x - 1)
	fmt.Println("B", f)
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB(1)
	funcC()
}
