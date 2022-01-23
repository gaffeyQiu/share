package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name     string
	category string
}

func (d *Dog) Name() string {
	return d.name
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func (d *Dog) Category() string {
	return d.category
}

func main() {
	dog := Dog{"dog1", "dog"}
	var pet Pet = &dog
	pet.SetName("dog2")
	fmt.Println(pet.Name())
}

//func main() {
//	var dog *Dog
//	if dog == nil {
//		fmt.Println("dog is nil")
//	}
//	var pet Pet = dog
//	if pet == nil {
//		fmt.Println("pet is nil")
//	}
//
//	fmt.Println(pet)
//}
