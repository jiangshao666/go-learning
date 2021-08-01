package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("conn read failed, err:%v", err)
			return
		}
		str := string(tmp[:n])
		fmt.Println(str)
	}
}

func main() {
	listner, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen on 127.0.0.1:20000 failed, err:%v", err)
		return
	}

	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Printf("Accept failed, err:%v", err)
			return
		}
		go process(conn)
	}
}
