package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Random number part
	seconds := time.Now().Unix() // Get the current date and time, as an integer
	rand.Seed(seconds)           // Seed the random number generator.
	target := rand.Intn(100) + 1 // Generate an integer between 1 and 100

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // Create a bufio.reader, which lets us read keyboard input

	guessed := false // Setup to print a failure message by default

	for i := 0; i < 10; i++ {
		fmt.Printf("Make a guess (You have %d remaining lives): \n", 10-i)
		input, err := reader.ReadString('\n') // Read what the user types, up until they press Enter
		if err != nil {                       // If there's an error, print the message and exit
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // Removing all the white-spaces
		guess, err := strconv.Atoi(input) // Converting a string to Integer
		if err != nil {                   // If there's an error, print the message and exit
			log.Fatal(err)
		}

		if guess == target { // If user guessed the number
			fmt.Println("Good job! You guess it!") // Printing success
			guessed = true                         // Prevent the failure message from displaying
			break                                  // Exit the loop
		} else if guess > target { // If the guess was too high, say so.
			fmt.Println("Oops. Your guess was HIGH.")
		} else { // If the guess was too low, say so.
			fmt.Println("Oops. Your guess was LOW.")
		}
	}
	if !guessed { // if `success` is false, tell player what the target was
		fmt.Println("Sorry. You didn't guess my number. It was:", target)
	}
}
