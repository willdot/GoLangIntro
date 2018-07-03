package main

import (
	"fmt"
)

func main() {
	simpleSlice()
	editASlice()
	changeSliceReference()
	appendASlice()
	loopThroughSlice()
	appendSliceToSlice()
}

func simpleSlice() {
	fmt.Println("Simple slice start")
	// Declare the slice
	shoppingList := make([]string, 5, 10)
	todoList := []string{"learn go", "make go api"}
	fmt.Println("\nlength is: ", len(shoppingList), "nCapacity is: ", cap(shoppingList))
	fmt.Println("\nlength is: ", len(todoList), "nCapacity is: ", cap(todoList))
}

func editASlice() {
	fmt.Println("Edit slice start")
	aSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(aSlice[4])
	aSlice[1] = 55
	fmt.Println(aSlice[1])
}

func changeSliceReference() {

	fmt.Println("Change a slice start")
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice[4])

	sliceOfSlice := aSlice[2:4]
	fmt.Println(sliceOfSlice)

	// Edit the 3rd value of the underlying array and then see it changed in both slices
	aSlice[2] = 22
	fmt.Println(aSlice)
	fmt.Println(sliceOfSlice)
}

func appendASlice() {

	fmt.Println("Append a slice start")
	aSlice := make([]int, 1, 4)
	fmt.Println("\nlength is: ", len(aSlice), "nCapacity is: ", cap(aSlice))

	// Once it tries to add in the 5th value (above the capacity of the array) it will double the capacity of the array. Then when it gets to the 9th value, the same thing happens
	for i := 1; i < 17; i++ {
		aSlice = append(aSlice, i)
		fmt.Println("\nCapacity is: ", cap(aSlice))
	}
}

func loopThroughSlice() {

	fmt.Println("Loop through slice start")

	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)

	for _, i := range aSlice {
		fmt.Println(i)
	}
}

func appendSliceToSlice() {

	fmt.Println("Append slice to slice start")
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)

	newSlice := []int{6, 7, 8, 9}

	aSlice = append(aSlice, newSlice...)
	fmt.Println(aSlice)
}
