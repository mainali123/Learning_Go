package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") // Opening the file that returns: pointer to the file and an error value
	if err != nil {                  // If there was error opening the file
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // Creating a scanner to read the file
	for scanner.Scan() {              // Scanning the context of the file line by line (Similar to for loop)
		fmt.Println(scanner.Text())
	}
	err = file.Close() // Closing the file, returns an error value if it fails
	if err != nil {    // If closing the file encountered an error
		log.Fatal(err)
	}
	if scanner.Err() != nil { // If an error is encountered while scanning the file
		log.Fatal(scanner.Err())
	}
}
