package main

import (
	"fmt"
	"reflect"
)

var (
	name, course  string
	module        float64
	nameInfered   = "Will"
	courseInfered = "GoLang"
	moduleInfered = 3.1
)

func main() {

	nameFunc := "Will func"
	moduleFunc := 3.2
	ptr := &moduleFunc

	fmt.Println("Name is set to ", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Module is set to ", module, "and is of type ", reflect.TypeOf(module))

	fmt.Println("NameFunc is set to ", nameInfered, "and is of type ", reflect.TypeOf(nameInfered))
	fmt.Println("ModuleFunc is set to ", moduleInfered, "and is of type ", reflect.TypeOf(moduleInfered))

	fmt.Println("NameFunc is set to ", nameFunc, "and is of type ", reflect.TypeOf(nameFunc))
	fmt.Println("ModuleFunc is set to ", moduleFunc, "and is of type ", reflect.TypeOf(moduleFunc))

	fmt.Println("Memory address of *moduleFunc* variable is ", ptr, "and the value of *moduleFunc* is", *ptr)
}
