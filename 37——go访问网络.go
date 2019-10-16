package main

import (
	"fmt"
	"io"
	"net"
	_ "reflect"
)

// http client demo

func main() {
	conn, err := net.Dial("tcp", "http://127.0.0.1:5000/updateHost")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n") // 发送请求

	var buf [8192]byte
	// 接收响应
	for {
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("get response failed, err:", err)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
