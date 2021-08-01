package main

import "fmt"

//结构体

type person struct {
	name   string
	age    int
	gender int
	hobby  []string
}

func main() {
	var p person
	p.name = "zhangsan"
	p.age = 20
	p.gender = 1
	p.hobby = []string{"吃", "喝"}

	fmt.Println(p)
	fmt.Printf("%T \n", p)

	// 匿名结构体, 多用于临时场景
	var s struct {
		x int
		y int
	}

	s.x = 10
	s.y = 1

	fmt.Println(s)
	fmt.Printf("%T \n", s)

}
