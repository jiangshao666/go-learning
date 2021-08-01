package main

import "fmt"

// 接口 type usb interface
// 接口是一种类型

type speaker interface {
	speak(words string)
	move()
}

type cat struct {
}

func (c cat) speak(words string) {
	fmt.Println("cat speak " + words)
}

type dog struct {
}

func (d dog) speak() {
	fmt.Println("dog speak ")
}

type wood struct {
}

type person struct {
	name string
}

func (p person) speak(words string) {
	fmt.Println("person speak " + words)
}

func (p person) move() {
	fmt.Println("person move")
}

func hurt(s speaker, words string) {
	s.speak(words)
}

func main() {
	//var c cat
	//var d dog
	//var w wood
	var p person
	//hurt(c, "miaomiao~") //只有实现了interface中的索引函数的，才能赋值给interface类型
	//hurt(d, "wangwang~") //实现interface函数的时候必须是同样的函数类型（参数、返回值等）
	//hurt(w)              // 只有实现了interface的函数的struct才能赋值给对应的interface类型对象
	hurt(p, "aouch") //正确的示例

	// interface变量赋值
	var sp speaker
	//sp = c1
	//sp = dog
	//sp = wood
	sp = p
	sp.move()
}
