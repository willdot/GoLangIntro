package main

import (
	"fmt"
	"os" // This gives access to environment variables
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// Get the username environment variable
	name := os.Getenv("USERNAME")

	fmt.Println("Hello ", name)

}
