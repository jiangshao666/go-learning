package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 21000,
	})
	if err != nil {
		fmt.Printf("listen udp faild, err:%v", err)
		return
	}

	defer conn.Close()
	var buffer [1024]byte
	for {
		n, fromAddr, err := conn.ReadFromUDP(buffer[:])
		if err == io.EOF {
			fmt.Println("read udp over")
			return
		}
		if err != nil {
			fmt.Printf("read udp  faild, err:%v", err)
			return
		}
		str := strings.ToUpper(string(buffer[:n]))
		_, err = conn.WriteToUDP([]byte(str), fromAddr)
		if err != nil {
			fmt.Printf("write udp  faild, err:%v", err)
			return
		}
	}
}
