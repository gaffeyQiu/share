package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	Client  string
	Content string
}

func main() {
	addr := "localhost:9090"
	Start(addr)
}

// Start 启动服务器
func Start(addr string) {
	// 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("listen tcp port failed: %v\n", err)
		return
	}

	// 客户端列表
	clientList := make(map[string]net.Conn)
	// 消息通道
	messageChan := make(chan Message, 10)

	// 广播消息
	go BroadMessages(clientList, messageChan)

	fmt.Printf("server start at %s ...\n", addr)
	// 启动
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept failed:%v\n", err)
			continue
		}
		fmt.Printf("new client online: %s\n", conn.RemoteAddr().String())

		clientList[conn.RemoteAddr().String()] = conn

		// 处理消息
		go Handler(conn, clientList, messageChan)
	}
}

// BroadMessages 向所有在线的客户端发广播
func BroadMessages(clientList map[string]net.Conn, messages chan Message) {
	for {
		// 不断从通道里读取消息
		msg := <-messages

		for key, conn := range clientList {
			// 不广播给自己
			if conn.RemoteAddr().String() == msg.Client {
				continue
			}
			writeMsg := fmt.Sprintf("[%s]:  %s", msg.Client, msg.Content)
			_, err := conn.Write([]byte(writeMsg))
			if err != nil {
				log.Printf("broad message to %s failed: %v\n", key, err)
				delete(clientList, key)
			}
		}
	}
}

// Handler 处理客户端发到服务端的消息，将其扔到通道中
func Handler(conn net.Conn, clientList map[string]net.Conn, messages chan Message) {
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			log.Printf("read client message failed:%v\n", err)
			delete(clientList, conn.RemoteAddr().String())
			conn.Close()
			break
		}

		// 把收到的消息写到通道中
		reciveMsg := Message{
			Client:  conn.RemoteAddr().String(),
			Content: string(buf[0:length]),
		}
		messages <- reciveMsg
	}
}
