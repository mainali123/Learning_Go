package main

import "fmt"

func main() {
	var notes []string        // Declaration of slice
	notes = make([]string, 8) // Create a slice with eight strings

	notes[0] = "Sa"
	notes[1] = "Re"
	notes[2] = "Ga"

	fmt.Println(notes[0])
	fmt.Println(notes[1])

	primes := make([]int, 5) // Create a slice with five integers, and set up a variable to hold it.
	primes[0] = 2
	primes[1] = 3

	// Slice literals
	notes1 := []string{"Sa", "Re", "Ga", "Ma", "Pa", "Dha", "Ni", "Sa"} // Assign values using a slice literal
	fmt.Println(notes1)

	primes1 := []int{ // A multi-line slice literal
		2,
		3,
		5,
		7,
	}
	fmt.Println(primes1)

	// Slicing
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3] // From index 0 to 2 (excluding 3)
	fmt.Println(slice1)

	// Append function

	slice := []string{"a", "b"} // Creating a slice
	fmt.Println(slice, len(slice))
	// Assign the return value of "append" back to the same slice variable
	slice = append(slice, "c") // Append an element to the end of the slice
	fmt.Println(slice, len(slice))
	// Assign the return value of "append" back to the same slice variable
	slice = append(slice, "d", "e") // Append two elements to the end fo the slice
	fmt.Println(slice, len(slice))

}
