package main

import (
	"bufio"
	"fmt"
	"os"
)

// 采用fmt.scan读取内容有一个问题就是一旦内容带空格，则空格后面的内容将无法读取
func scanRead() {
	fmt.Print("请输入内容:")
	var s string
	fmt.Scanln(&s)
	fmt.Println("输入的内容为:" + s)

}

// 采用bufio的形式可以解决该问题
func bufioRead() {
	fmt.Print("请输入内容:")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("读取异常 err:%v", err)
		return
	}
	fmt.Println("输入的内容为:" + s)

}

func main() {
	//scanRead()
	bufioRead()
}
