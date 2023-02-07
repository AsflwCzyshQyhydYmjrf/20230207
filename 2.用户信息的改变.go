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
		Name:    "John",
		Age:     30,
		Address: "New York",
	}

	fmt.Printf("User: %+v\n", user)

	user.UpdateName("Jack")
	user.UpdateAge(35)
	user.UpdateAddress("Los Angeles")

	fmt.Printf("User: %+v\n", user)
}
