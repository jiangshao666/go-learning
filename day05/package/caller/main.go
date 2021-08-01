package main

import (
	"fmt"

	"github.com/jiangshao666/day05/package/importer" //引入第三方package
)

// 执行顺序: 首先是全局变量、然后是init函数，然后是main函数
// 如果有引入package, 则按照package引用顺序的逆序执行

var globali = 10

const pi = 3.14

func init() {
	fmt.Println("main init ...")
	fmt.Println("global var", globali, pi)
}

func main() {
	ret := importer.Add(10, 20)
	fmt.Println("Add result", ret)
}
