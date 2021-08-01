package importer

import "fmt"

// package被加载的时候自动调用，不能手动调用
func init() {
	fmt.Println("importer package init...")
}

// 函数首字母大写，才能被外部调用的package访问
func Add(x int, y int) int {
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}
