package main

import (
	"fmt"
	"time"
)

func main() {
	countdown()

	rangeFor()
}

func countdown() {
	for timer := 5; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("\nFINISHED!!!!!!")
			break
		} else if timer == 1 {
			fmt.Println("\nNearly there")
			continue
		}

		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}

func rangeFor() {

	food := []string{"burger", "pizza", "chocolate"}
	moreFood := []string{"pasta", "cheese", "pizza"}

	// This is a label that allows the break in the second loop to break out of both loops
OutterLoop:
	// The underscore is the unused variable which here is the index number.
	for _, i := range food {
		fmt.Println(i)
		for _, j := range moreFood {
			if i == j {
				fmt.Println("\nSomething isn't right", i, "is in both lists")
				break OutterLoop
			}
		}
	}
}
