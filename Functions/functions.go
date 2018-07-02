package main

import (
	"fmt"
	"strings"
)

func main() {

	module := "Function basics"
	author := "Will"

	fmt.Println(converter(module, author))

	highestNumber := getHighestNumber(10, 29, 1, 56, 99)
	fmt.Println("Highest number is: ", highestNumber)
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}

// Variadic function that takes an unspecified amount of parameters
func getHighestNumber(inputs ...int) int {
	highest := inputs[0]

	for _, i := range inputs {
		if i > highest {
			highest = i
		}
	}

	return highest
}
