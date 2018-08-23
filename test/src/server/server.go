package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("开始一个TCP服务")

	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	fmt.Println("Listening")
	for true {
		conn, err := listener.Accept()
		
		if err != nil {
			fmt.Println("Error accepting", err.Error())
		}

		fmt.Println("Accept")
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for true {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Println("len", len)
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}
