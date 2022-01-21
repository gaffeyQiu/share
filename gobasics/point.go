package main

import "fmt"

func main() {
	a := 1
	b := &a
	*b = 10
	fmt.Println(a)
}
