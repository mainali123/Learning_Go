package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println(sum / float64(len(numbers)))

	ans, err := GetFloats("data.txt")
	if err == nil {
		for _, an := range ans {
			fmt.Println(an)
		}
	}
}
