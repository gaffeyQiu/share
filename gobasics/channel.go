package main

import (
	"fmt"
	"time"
)

// Producer 生产者: 生成 factor 整数倍的序列
func Producer(out chan<- int) {
	for i := 0; ; i++ {
		out <- i
	}
}

// Consumer 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(ch)
	go Consumer(ch)

	// 运行一定时间后退出
	time.Sleep(1 * time.Second)
}
