package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读文件

// 手动指定读取长度读取
func readFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("file open failed err:%v \n", err)
		return
	}
	defer fileObj.Close()

	var tmp [128]byte

	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("read finish")
			return
		}
		if err != nil {
			fmt.Printf("file read failed err:%v \n", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			fmt.Println("read finish")
			return
		}
	}
}

// 带缓冲区的read 使用bufio包
func readFile2() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("file open failed err:%v \n", err)
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read finish")
			return
		}

		if err != nil {
			fmt.Printf("file read failed err:%v \n", err)
			return
		}

		fmt.Print(line)
	}
}

// ioutil包读取文件
func readFile3() {
	bytes, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("file read failed err:%v \n", err)
		return
	}
	fmt.Println(string(bytes))
	fmt.Println("read finish")
}

func main() {
	//readFile1()
	//readFile2()
	readFile3()
}
