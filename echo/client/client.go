package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7000")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalf("dial tcp error: %v\n", err)
	}
	defer conn.Close()

	fmt.Println("connect!")
	
	
	
	for {
		reader := bufio.NewReader(os.Stdin)
		typed, _, _ := reader.ReadLine()
		fmt.Println(typed)
		conn.Write(typed)
	}
}