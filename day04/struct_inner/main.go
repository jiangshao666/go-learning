package main

import "fmt"

// 匿名结构体字段
// 注意：匿名字段一般用于字段少且类型不太重复的情况
type monster struct {
	int    //等同于 int:int
	string //等同于 string:string
	hp     int
}

type address struct {
	province string
	city     string
}

type compony struct {
	name string
	mem  int
}

// 结构体嵌套
type person struct {
	name    string
	age     int
	address //匿名嵌套结构体，等同于 address: address
	comp    compony
}

func main() {
	mon := monster{
		10, "Boss", 100,
	}
	fmt.Println(mon, mon.int, mon.string, mon.hp)

	p := person{
		name:    "zhangsan",
		age:     25,
		address: address{"guangdong", "shenzhen"},
		comp:    compony{"tencent", 10000},
	}

	fmt.Println(p)
	fmt.Println(p.address.city)
	fmt.Println(p.city)      // 等同于(p.address.city), 因为address是一个匿名嵌套结构体
	fmt.Println(p.comp.name) //comp是一个非匿名嵌套结构体，访问属性必须通过comp层级访问
}
