package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000
	b := 3

	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type ", reflect.TypeOf(b))

	c := int(a) + b // a is a floating point so convert to int first

	fmt.Println("\nC has value: ", c, "and is of type ", reflect.TypeOf(c))
}
