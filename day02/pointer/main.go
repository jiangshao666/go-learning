package main

import "fmt"

// 指针
func main() {

	n := 10
	p := &n //&是取地址符号，返回的是n的内存地址，一个十六进制表示的值
	fmt.Printf("%T %v \n", p, p)
	m := *p // *表示获取指针所指的地址的数值
	fmt.Printf("%T %v \n", m, m)

	// 注意指针的定义需要初始化，否则容易导致空指针
	// var a *int; //定义了一个int类型的指针，但是并没有初始化地址，因此地址其实是nil
	// fmt.Println(a)
	// *a = 10; //对地址为nil的指针赋值会导致空指针
	// fmt.Println(*a)

	// 初始化指针可以用new 或者 make
	// new(Type t) 函数通常用于基本数据类型string, int等类型内存申请,返回的指一个类型为Type的指针
	// make(Type t) 指能用于slice, map, channel的初始化，并且返回的就是这几个类型的对象，因为这几个类型都是引用类型的
	
	// 使用new初始化指针
	var n1 = new (int) //定义一个int类型的指针并初始化指针地址，指针所指向地址的指为默认0值
	fmt.Println(n1,*n1) 
	*n1 = 10
	fmt.Println(*n1)

	// 使用make初始化对象

	var mm map[string]int; //定义一个map，但是没有分配内存
	mm = make(map[string]int, 10) //真正的内存分配
	mm["key"] = 100;
	fmt.Println(mm)

}