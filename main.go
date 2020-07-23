package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Find min and max from input.")

	var numberOfInput string
	fmt.Scanln(&numberOfInput)

	numberOfInputInt, _ := strconv.Atoi(numberOfInput)
	// fmt.Println(numberOfInputInt)

	var min int
	var max int
	var line string
	for i := 1; i <= numberOfInputInt; i++ {
		fmt.Scanln(&line)
		lineInt, _ := strconv.Atoi(line)
		if i == 1 {
			min = lineInt
			max = lineInt
		}

		if lineInt < min {
			min = lineInt
		} else if lineInt > max {
			max = lineInt
		}
	}

	fmt.Println(min)
	fmt.Println(max)
}
