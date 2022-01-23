package main

import "fmt"

type Any interface{}

func main() {
	var T Any
	x := 1
	str := "str"
	T = x
	T = str

	fmt.Println(T.(string))
}
