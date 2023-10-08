package main

import (
	"fmt"
	"sort"
)

type user struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	//var myOtehrMap map[string]string
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"

	fmt.Println(myMap["dog"])
	fmt.Println(myMap["other-dog"])

	myMap1 := make(map[string]int)

	myMap1["first"] = 1
	myMap1["second"] = 2

	fmt.Println(myMap1["first"])
	fmt.Println(myMap1["second"])

	myMap2 := make(map[string]user)
	myMap2["first"] = user{FirstName: "Diwash", LastName: "Mainali"}

	fmt.Println(myMap2["first"].FirstName, myMap2["first"].LastName)

	// Slices
	var mySlice []string
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Marry")

	sort.Strings(mySlice)

	fmt.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	fmt.Println(numbers[0:2]) // prints range
}
