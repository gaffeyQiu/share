package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func main() {
	user := User{
		ID:   1,
		Name: "Alan",
	}
	fmt.Printf("userID = %d, userName = %s\n", user.ID, user.Name)

	fmt.Printf("userName = %s\n", user.GetName())
}

func (u User) GetName() string {
	return u.Name
}
