package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	Start(os.Args[1])
}

func Start(addr string) {
	// 向服务器发起连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Dial to server failed: %v\n", err)
	}

	// 向服务器发消息
	go SendMsg(conn)
	// 接收来自服务器消息
	go ReadMsg(conn)
	for {
	}
}

// SendMsg 向服务器端发消息
func SendMsg(conn net.Conn) {
	for {
		var input string
		// 接收输入消息，放到input变量中
		fmt.Scanln(&input)

		// 退出
		if input == ":q" {
			fmt.Println("Bye bye ...")
			conn.Close()
			os.Exit(0)
		}

		// 内容为空
		if len(input) == 0 {
			continue
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			break
		}
	}
}

func ReadMsg(conn net.Conn) {
	// 接收来自服务器端的广播消息
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("recv server msg failed: %v\n", err)
			conn.Close()
			os.Exit(0)
			break
		}
		fmt.Println(string(buf[0:length]))
	}
}
