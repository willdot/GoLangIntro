package main

import (
	"fmt"
)

func main() {
	name := "Will"
	course := "Go"

	fmt.Println("\nHi ", name, "you're currently doing ", course)

	// Pass by value
	changeCourseByValue(course)
	fmt.Println("You're now doing ", course)

	// Pass by reference
	changeCourseByReference(&course)
	fmt.Println("You're now doing ", course)
}

func changeCourseByValue(course string) string {
	course = "Go Fundaments"

	fmt.Println("\nTrying to change your course to ", course)

	return course
}

// The asterisk means we are using a reference to a variable
func changeCourseByReference(course *string) string {
	*course = "Go Fundamentals"

	fmt.Println("\nTrying to change your course to ", *course)

	return *course
}
