package main

import (
	"fmt"
)

func main() {

	fmt.Println("\nAbout to run a normal if")
	normalIf()

	fmt.Println("\nAbout to run an initialized if")
	initializeIf()

	fmt.Println("\nAbout to switch things up")
	switching("Hello")
	fmt.Println("\nAbout to switch it up again")
	switching("Something")
	fmt.Println("\nAbout to switch it up once again")
	switching("")
	fmt.Println("\nAbout to switch it up for the last time")
	switching("Jibberish")
}

func normalIf() {
	firstValue := "2"
	secondValue := "1"

	if firstValue > secondValue {
		fmt.Println("\nfirstValue is the higher number")
	} else if firstValue < secondValue {
		fmt.Println("\nsecondValue is the higher number")
	} else {
		fmt.Println("\nBoth values are the same")
	}
}

func initializeIf() {

	// This initializes the variables in the scope of the if statement. Once the if statement is complete, they are disposed off and out of scope
	if firstValue, secondValue := 2, 1; firstValue > secondValue {
		fmt.Println("\nfirstValue is the higher number")
	} else if firstValue < secondValue {
		fmt.Println("\nsecondValue is the higher number")
	} else {
		fmt.Println("\nBoth values are the same")
	}
}

func switching(input string) {
	switch input {
	case "Hello":
		fmt.Println("\nYou said hello")
	case "Something":
		fmt.Println("\nYou said something")
	case "", " ":
		fmt.Println("\nYou didn't say anything")
	default:
		fmt.Println("\nYou said something we don't understand")
	}
}
