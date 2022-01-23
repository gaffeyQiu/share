package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

func HappenErr() bool {
	return false
}

func returnsMyError() error {
	var p *MyError = nil

	if HappenErr() {
		p = &MyError{error: errors.New("happen error")}
	}

	return p
}
func main() {
	err := returnsMyError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}
