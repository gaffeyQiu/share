package main

import (
	"fmt"
)

var x string = "a"

func main() {

	fmt.Printf("The x is %v\n", x)

	var x int
	x = 1
	fmt.Printf("The x is %v\n", x)
	{
		var x float32
		x = 0.1
		fmt.Printf("The x is %v\n", x)
	}
}
