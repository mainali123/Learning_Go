package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirhtDate   time.Time
}

type myStruct struct {
	name string
}

func (m *myStruct) printName() string { // ties function with the structure and can use as it is it's value
	return m.name
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "9863025000",
	}
	fmt.Println(user.FirstName, user.LastName, user.BirhtDate)

	var firstperson myStruct
	firstperson.name = "Diwash"

	secondPerson := myStruct{name: "Utsav"}

	fmt.Println("first person", firstperson.printName())
	fmt.Println("second person", secondPerson.printName())
}
