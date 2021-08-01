package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Dial 127.0.0.1:20000 failed, err:%v", err)
		return
	}

	//var msg string
	// if len(os.Args) < 2 {
	// 	msg = "hello world"
	// } else {
	// 	msg = os.Args[1]
	// }

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("请输入发送内容:")
		msg, _ := reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	conn.Close()
}
