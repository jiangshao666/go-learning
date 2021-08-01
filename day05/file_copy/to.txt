package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(src string, dest string) bool {
	fileFrom, err := os.Open(src)
	if err != nil {
		fmt.Printf("file src open failed, err:%v", err)
		return false
	}
	fileDest, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 666)
	if err != nil {
		fmt.Printf("file dest open failed, err:%v", err)
		return false
	}
	defer fileDest.Close()
	defer fileFrom.Close()

	readBuffer := bufio.NewReader(fileFrom)
	writeBuffer := bufio.NewWriter(fileDest)
	for {
		line, err := readBuffer.ReadString('\n')
		if err == io.EOF {
			fmt.Println("copy done")
			return true
		}
		writeBuffer.WriteString(line)
		writeBuffer.Flush()
	}
}

func main() {
	ret := copyFile("./main.go", "./to.txt")
	fmt.Println("copy file ok?", ret)
}
