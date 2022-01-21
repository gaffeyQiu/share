package main

import "fmt"

func main() {
	nameAgeMap := make(map[string]int)

	nameAgeMap["Alan"] = 10
	nameAgeMap["Ben"] = 20
	nameAgeMap["Carl"] = 30
	nameAgeMap["David"] = 40

	for k, v := range nameAgeMap {
		fmt.Printf("key: %s, val: %d\n", k, v)
	}

	//delete(nameAgeMap, "Ben")
	//fmt.Println(nameAgeMap)

	//alanAge := &nameAgeMap["Alan"]
}
