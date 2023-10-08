package main

import "fmt"

func main() {
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	for _, anima := range animals {
		fmt.Println(anima)
	}
}
