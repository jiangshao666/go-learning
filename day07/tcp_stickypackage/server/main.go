package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/jiangshao666/day07/tcp_stickypackage/proto"
)

// tcp 粘包
// 因为发送方开启Nagle算法或者因为接收端缓冲区堆积且接送速度不够快导致粘包

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buff [1024]byte
	for {
		// n, err := reader.Read(buff[:])
		// if err == io.EOF {
		// 	break
		// }
		// if err != nil {
		// 	fmt.Printf("read failed, err:%v", err)
		// 	return
		// }
		//str := string(buff[:n])
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		fmt.Println("receive content:", msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen on 127.0.0.1:20000 failed, err:%v", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept failed, err:%v", err)
			return
		}
		process(conn)
	}

}
