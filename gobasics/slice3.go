package main

import "fmt"

func main() {
	s1 := make([]uint8, 1023, 1024)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s1 = append(s1, 1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	s1 = append(s1, 1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
}
