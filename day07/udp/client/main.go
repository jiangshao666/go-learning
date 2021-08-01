package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 21000,
	})
	if err != nil {
		fmt.Printf("udp connect failed, err:%v", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var buff [1024]byte
	for {
		fmt.Printf("请输入:")
		str, _ := reader.ReadString('\n')
		if str == "exit" {
			break
		}
		conn.Write([]byte(str))

		// 收包
		n, _, err := conn.ReadFromUDP(buff[:])
		if err != nil {
			fmt.Printf("ReadFromUDP failed, err:%v", err)
		}
		fmt.Printf("recive resp: %s", string(buff[:n]))
	}
}
