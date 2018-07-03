package main

import (
	"fmt"
)

func main() {
	createMaps()
	loopThroughMap()
	getMapValue()
	editMapValue()
	addIfNotExist()
	deleteFromMap()
}

func createMaps() {

	fmt.Println("Creating a map")

	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2

	anotherMap := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println(aMap)
	fmt.Println(anotherMap)
}

func loopThroughMap() {

	fmt.Println("Looping through a map")

	// Remember this will be unordered, and different results each time!!!
	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	for key, value := range testMap {
		fmt.Println("\nKey is: ", key, "Value is: ", value)
	}
}

func getMapValue() {

	fmt.Println("Getting value from a map")

	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println(testMap["b"])

}

func editMapValue() {

	fmt.Println("Editing a value")

	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println(testMap)
	testMap["b"] = 999
	fmt.Println(testMap)
}

func addIfNotExist() {

	fmt.Println("When trying to edit a key that doesn't exist, it will add that key and the value")

	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fmt.Println(testMap)
	testMap["f"] = 9
	fmt.Println(testMap)
}

func deleteFromMap() {

	fmt.Println("Deleting a key")

	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println(testMap)
	delete(testMap, "d")
	fmt.Println(testMap)
}
