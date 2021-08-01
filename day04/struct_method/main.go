package main

import "fmt"

// 结构体方法

type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) dog {
	return dog{
		name,
		age,
	}
}

// 方法, 是作用于特定类型的函数,
// 接收者表示调用该方法的具体的类型变量，多用类型名首字母小写表示
// 该接收者为值接收者, 所以本质上是拷贝了一个dog对象
func (d dog) addAge() {
	d.age++
}

// 指针接收者, 本质上是传递一个对象的指针
/*
什么时候应该使用指针类型接收者?
	1)需要修改接收者中的值
	2)接收者是拷贝代价比较大的大对象
	3)保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
func (d *dog) subAge() {
	d.age--
}

func (d dog) wang() {
	fmt.Printf("%s 汪汪汪 \n", d.name)
}

//	在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
//  非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
type myInt int

func (i myInt) sayType() {
	fmt.Println("this is myInt")
}

func main() {
	d := newDog("李四", 18)
	d.wang()
	// 值接收者传参
	fmt.Println(d.age)
	d.addAge()
	fmt.Println(d.age)
	// 指针接收者传参
	d.subAge()
	fmt.Println(d.age)

	// 针对任意类型调用方法
	i := myInt(10)
	i.sayType()
}
