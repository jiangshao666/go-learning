package main

import "fmt"

// 构造函数， go中没有内置的构造函数，但是可以自己封装
// 约定俗称的方式是通过newXXX的方式定义构造函数

type person struct {
	name string
	age int
}

// person构造函数,该方式没new一个person会拷贝一次对象
// func newPerson(name string, age int ) person {
// 	return person{
// 		name, age,
// 	}
// }

//构造函数方式2, 当结构体比较大的生活尽量使用结构体指针，减少程序的内存拷贝开销
func newPerson(name string, age int ) *person {
	return &person{
		name, age,
	}
}

func main() {
	p1 := newPerson("zhangsan", 19)
	p2 := newPerson("lisi", 30)
	fmt.Println(*p1, *p2)
}