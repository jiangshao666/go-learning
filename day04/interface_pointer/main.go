package main

import "fmt"

// 接口使用值接收者和指针接收者的区别
// 使用值接收者，结构体类型和结构体指针类型的变量都可以赋值给相应的interface类型的变量
// 使用指针接收者，则只能使用结构体指针类型赋值给interface类型的变量

type speaker interface {
	speak(words string)
	move()
}

type person struct {
	name string
}

// 使用值接收者实现了接口
// func (p person) speak(words string) {
// 	fmt.Println("person speak " + words)
// }

// func (p person) move() {
// 	fmt.Println("person move")
// }

//使用指针接收者
func (p *person) speak(words string) {
	fmt.Println("person speak " + words)
}

func (p *person) move() {
	fmt.Println("person move")
}

func main() {
	var sp speaker

	p1 := person{"zhangsan"}
	p2 := &person{"lisi"}

	sp = &p1
	fmt.Println(sp)
	sp = p2
	fmt.Println(sp)

}
