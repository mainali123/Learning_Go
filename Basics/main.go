package main

import "fmt"

func main() {
	fmt.Println("Hey there!!") // Print line

	var whatToSay string = "Diwash Mainali"
	var i int = 34

	fmt.Println(whatToSay)
	fmt.Println("is is set to", i)

	whatWasSaid, whatWasNotSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, whatWasNotSaid)
}

func saySomething() (string, string) {
	return "something", "nothing"
}
