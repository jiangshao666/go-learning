package main

import "fmt"

// go 模拟类的继承

type animal struct {
	kind int
}

func (a animal) move() {
	fmt.Printf("类型:%d 可以移动\n", a.kind)
}

type person struct {
	name string
	animal
}

func (p person) talk() {
	fmt.Println("人可以谈话, 人的类型:", p.kind)
}

func main() {
	p := person{
		"zhangsan",
		animal{1},
	}

	fmt.Println(p)
	p.move()
	p.talk()
}
