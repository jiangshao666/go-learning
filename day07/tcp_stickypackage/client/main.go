package main

import (
	"fmt"
	"net"

	"github.com/jiangshao666/day07/tcp_stickypackage/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("connect to 127.0.0.1:20000 failed, err:%v", err)
		return
	}

	for i := 0; i < 20; i++ {
		msg := "I am fine, How are you?"
		//conn.Write([]byte(msg))
		bytes, err := proto.Encode(msg)
		if err != nil {
			fmt.Printf("Encode failed, err:%v", err)
			return
		}
		conn.Write(bytes)
	}
	conn.Close()

}
