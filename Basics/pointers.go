package main

import "fmt"

func main() {
	var myString string = "Green"

	fmt.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	fmt.Println("myString is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
