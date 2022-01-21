package main

import "fmt"

func main() {
	s1 := make([]int, 0, 0)
	for i := 1; i <= 10; i++ {
		s1 = append(s1, i)
		fmt.Printf("第 %d 次循环 cap = %d\t s1 = %v\n", i, cap(s1), s1)
	}
	//s1 = append(s1, 1, 2, 3)
}
