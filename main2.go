package main

import (
	"fmt"
)

func main() {
	fmt.Println("Find min and max from input.")

	var numberOfInputInt int
	fmt.Scanln(&numberOfInputInt)

	var lineInt int
	fmt.Scanln(&lineInt)

	var min = lineInt
	var max = lineInt

	numberOfInputInt--

	for numberOfInputInt > 0 {
		fmt.Scanln(&lineInt)
		if lineInt < min {
			min = lineInt
		} else if lineInt > max {
			max = lineInt
		}

		numberOfInputInt--
	}

	fmt.Println(min)
	fmt.Println(max)
}
