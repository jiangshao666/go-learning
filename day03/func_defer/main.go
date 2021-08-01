package main

import "fmt"

// go中的关键字defer用于延迟执行指定代码段,必须放在func内部
// 在defer归属的函数即将返回时，讲延迟处理的语句按照defer声明是顺序逆序执行，也就是先defer的语句后执行，后defer的语句先执行
// 通常的应用是用于资源释放、如文件句柄，数据库、网络连接等的回收

func deferDemo() {
	fmt.Println("start")
	//defer后面的语句延迟到函数即将返回时再执行
	// 一个函数中可以定义多个defer
	defer fmt.Println("defer01")
	defer fmt.Println("defer02")
	fmt.Println("end")
}

// go语言中的return语句不是原子操作，会拆分成两步，第一步时返回值赋值，第二步时执行RET指令，而defer语句的内容执行时间正好介于这两部之间

func f1() int { //匿名返回值
	x := 5 //第一步 返回值赋值 = 5
	defer func() {
		x++ //修改的时x不是返回值
	}()
	return x
}

func f2() (x int) { // 返回值 = x = 5
	defer func() {
		x++ //修改x，又因为返回值=x
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x, 不影响y和返回值
	}()
	return x // 返回值 = y =x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ //这个地方的参数x时一个拷贝的副本，因此不影响返回值
	}(x)
	return 5 // 返回值 = x = 5
}

// 面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	// 分析执行过程
	/*
		1) a=1 b=2
		2) defer calc("1", 1, calc("10", 1, 2))
		3) 执行calc("10", 1, 2) 打印 10 1 2 3
		4) 入栈 calc("1", 1, 3)
		5) a = 0
		6) defer calc("2", 0, calc("20", 0, 2))
		7) 执行calc("20", 0, 2) 打印 20 0 2 2
		8) 入栈 calc("2", 0, 2)
		9) b=1
		10) 出栈执行 calc("2", 0, 2) 打印 2 0 2 2
		11) 出栈执行 calc("1", 1, 3) 打印 1 1 3 4
	*/
}
