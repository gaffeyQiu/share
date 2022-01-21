package main

import "fmt"

func main() {
	x := 3
	if x == 1 {
		fmt.Println(1)
	} else if x == 2 {
		fmt.Println(2)
	} else {
		fmt.Println(x)
	}

	//for i := 0; i < 3; i++ {
	//	fmt.Println(x)
	//}

	//for x > 0 {
	//	fmt.Println(x)
	//	x--
	//}

	//for {
	//	fmt.Println(1)
	//}

	//switch x {
	//case 1:
	//	println(1)
	//	fallthrough
	//case 2:
	//	println(2)
	//	fallthrough
	//default:
	//	fmt.Println("default")

	//for i, v := range [...]int{1, 2, 3} {
	//	fmt.Println(i, v)
	//}
}
