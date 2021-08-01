package main

import (
	"flag"
	"fmt"
)

// flag 包，用于接收命令行参数, 区别于os.Args的是可以接收k=v形式的参数

// flag的语法 flag.type("参数名", "默认值", "提示语") 注意返回的是一个指针类型
func main() {
	// fmt.Println(os.Args)
	// fmt.Printf("%T \n", os.Args)

	// 方式一 flag.Type()函数
	// name := flag.String("name", "zhangsan", "姓名")
	// age := flag.Int("age", 22, "年龄")
	// married := flag.Bool("married", false, "婚否")
	//flag.Parse()
	// fmt.Println(*name, *age, *married)

	// 方式二 flag.TypeVar()函数

	var name string
	var age int
	var married bool
	flag.StringVar(&name, "name", "zhangsan", "姓名")
	flag.IntVar(&age, "age", 22, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")

	flag.Parse()

	fmt.Println(name, age, married)
	fmt.Println(flag.NArg())  //获取标准参数个数 非kv形式
	fmt.Println(flag.NFlag()) // 获取kv传参个数
	fmt.Println(flag.Args())  //获取非kv参数slice
}
