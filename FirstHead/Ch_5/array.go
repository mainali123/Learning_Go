package main

import "fmt"

func main() {
	// Declaration of array if the value is unknown
	var notes [4]string
	notes[0] = "Sa"
	notes[1] = "Re"
	notes[2] = "Ga"
	notes[3] = "Ma"

	for i := 0; i < 4; i++ {
		fmt.Println(notes[i])
	}

	// Array literals
	allNOtes := [8]string{"Sa", "Re", "Ga", "Ma", "Pa", "Dha", "Ni", "Sa"}

	fmt.Println(allNOtes)

	// looping safely using for...range
	for index, value := range allNOtes {
		fmt.Println(index, value)
	}
}
