package main

import (
	"fmt"
	"unsafe"
)

func main() {
	user := User{
		ID:   1,
		Name: "Alan",
	}
	fmt.Printf("userID = %d, userName = %s\n", user.ID, user.Name)

	fmt.Printf("User struct size: %d\n", unsafe.Sizeof(user))
	setName(&user, "Ben")

	fmt.Printf("userID = %d, userName = %s\n", user.ID, user.Name)
}

func setName(user *User, newName string) {
	fmt.Printf("User struct point size: %d\n", unsafe.Sizeof(user))
	user.Name = newName
}
