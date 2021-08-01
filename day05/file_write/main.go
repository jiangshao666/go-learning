package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 使用标准库写文件
func fileWrite() {
	fileWriter, err := os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("file open failed, err:%v", err)
		return
	}
	defer fileWriter.Close()

	fileWriter.Write([]byte("just coding"))
}

// 使用bufio写文件

func fileWrite2() {
	fileWriter, err := os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("file open failed, err:%v", err)
		return
	}
	defer fileWriter.Close()
	buffer := bufio.NewWriter(fileWriter)
	buffer.WriteString("i love NBA\n")
	buffer.Flush()
}

// 使用ioutil写文件
func fileWriter3() {
	s := "but i don't like NBA, I like football\n"

	err := ioutil.WriteFile("./out.txt", []byte(s), 666)
	if err != nil {
		fmt.Printf("file write failed, err:%v", err)
		return
	}

}

func main() {
	//fileWrite()
	//fileWrite2()
	fileWriter3()
}
