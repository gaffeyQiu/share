package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7000")
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("listener fail: %v\n", err)
	}
	defer listener.Close()

	fmt.Println("server start at :7000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
			continue
		}
		
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	for {
		buf := make([]byte, 0, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("server read error: %v\n", err)
			break
		}
		
		log.Printf("recived: %s", string(buf))

		writeMsg := "Hello" + string(buf)

		_, err = conn.Write([]byte(writeMsg))
		if err != nil {
			log.Printf("write msg error: %v\n", err)
			break
		}
	}	
}