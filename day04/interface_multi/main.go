package main

import "fmt"

// 一个结构体可以实现多个接口
// 接口还可以嵌套

type mover interface {
	move()
}

type speaker interface {
	speak()
}

// 接口嵌套
type animal interface {
	mover
	speaker
}

type person struct {
	name string
}

// person同时实现了两个接口
func (p *person) move() {
	fmt.Println("person move")
}

func (p *person) speak() {
	fmt.Println("person speak")
}

func main() {
	p := person{"lisi"}
	p.move()
	p.speak()
}
