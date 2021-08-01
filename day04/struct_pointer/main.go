package main


import "fmt"

// 结构体是值类型

type person struct {
	name string
	age  int
}

// 值传递
func f(p person) {
	p.age = 20 //修改的是副本的值
}

// 传递指针值
func f1(p *person) {
	//(*p).age = 20
	p.age = 20 // 等同于(*p).age = 20 ， 是一个语法糖
}

func main() {
	var p person
	p.name = "Jam"
	p.age = 10
	f(p)
	fmt.Println(p.age)

	//　传指针地址
	f1(&p)
	fmt.Println(p.age)

	// 创建结构体指针， 方式1
	var pp = new(person)
	f1(pp)
	fmt.Printf("%T %p %p\n", pp, pp, &pp)
	fmt.Println(pp.age)

	// 创建结构体指针，方式2 key-value方式初始化 &pp2即为结构体指针
	var pp2 = person{
		name: "Sam",
		age:  20,
	}
	fmt.Printf("%v#\n", pp2)
	fmt.Printf("%T %T %p \n", pp2, &pp2, &pp2)

	// 创建结构体指针，方式3 使用值列表的方式，需要保持值和属性顺序对应
	// &pp3即为结构体指针
	pp3 := person{
		"Tom",
		30,
	}
	fmt.Printf("%v#\n", pp3)
	fmt.Printf("%T %T %p \n", pp3, &pp3, &pp3)

}
