package main

import "fmt"

//func main() {
//	s1 := []int{1, 2, 3, 4}
//	s2 := s1
//	for i := 0; i < len(s1); i++ {
//		s2[i]++
//	}
//	fmt.Println(s1)
//}

func main() {
	s1 := []int{1, 2, 3, 4}
	for _, v := range s1 {
		v += 1
	}

	fmt.Println(s1)
}

//func main() {
//	s1 := s1
//
//	s1 = []int{1, 2, 3, 4}
////	s2 := append(s1, 5)
//	fmt.Println(s2)
//}
