package main

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of the file.
func GetFloats(fileName string) ([3]float64, error) { // The function will return an array of numbers and any error encountered
	var numbers [3]float64         // Declare the array we'll be returning
	file, err := os.Open(fileName) // Open the provided filename
	if err != nil {                // If there was an error opening the file, return it
		return numbers, err
	}
	i := 0 // This variable will track which array index we should assign to
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) // Convert the file line string to float64
		if err != nil {                                          // If there was an error converting the line to a number, return it
			return numbers, err
		}
		i++ // Move to the next array index
	}
	err = file.Close()
	if err != nil { // If there was an error closing the file, return it
		return numbers, err
	}
	if scanner.Err() != nil { // If there was an error scanning the file, return it
		return numbers, scanner.Err()
	}
	return numbers, nil // If we got this far, there were no errors, so return the array of numbers and a "nil" error.
}
