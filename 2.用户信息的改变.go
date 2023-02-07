package main

import (
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) UpdateAge(age int) {
	u.Age = age
}

func (u *User) UpdateAddress(address string) {
	u.Address = address
}

func main() {
	user := &User{
		Name:    "wyb",
		Age:     20,
		Address: "Nan Shan",
	}

	fmt.Printf("User: %+v\n", user)

	user.UpdateName("WYB")
	user.UpdateAge(22)
	user.UpdateAddress("CQUPT")

	fmt.Printf("User: %+v\n", user)
}
