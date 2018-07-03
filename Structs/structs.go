package main

import (
	"fmt"
)

func main() {
	type person struct {
		name       string
		age        int64
		occupation string
	}

	// Different ways of initializing
	var Person1 person

	Person1.name = "William"
	Person1.age = 27
	Person1.occupation = "Software engineer"
	fmt.Println("\n", Person1)

	// This creates a new reference. So if you want to see it's values, you have to use * in front like with the println
	Person2 := new(person)

	Person2.name = "Willdot"
	Person2.age = 28
	Person2.occupation = "Software"
	fmt.Println("\n", *Person2)

	Person3 := person{
		name:       "Will",
		age:        29,
		occupation: "Software developer",
	}

	fmt.Println("\n", Person3)

}
