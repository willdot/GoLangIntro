package main

import (
	"fmt"
	"strconv"
)

type SomeService interface {
	DoSomething(string) string
	DoSomethingElse(int, string) string
}

type aService struct{}

func (aService) DoSomething(s string) string {
	return s + "!"
}

func (aService) DoSomethingElse(i int, s string) string {
	return strconv.Itoa(i) + s
}

func printSomething(x SomeService) {
	fmt.Println(x.DoSomething("Hello again"))
}

func main() {
	serv := aService{}

	response := serv.DoSomething("Go Go Go")

	fmt.Println(response)

	response = serv.DoSomethingElse(3, " is the magic number")

	fmt.Println(response)

	printSomething(serv)

	genInterface(serv)
}

func genInterface(i interface{}) {

	v, ok := i.(aService)

	if ok == true {
		fmt.Println(v.DoSomething("This was generic"))
	}
}

/*type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type someService struct{}
type someOtherService struct{}

func (someService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty String")
	}
	return strings.ToUpper(s), nil
}

func (someOtherService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty String")
	}
	return strings.ToUpper(s), nil
}

func (someService) Count(_ context.Context, s string) int {
	return len(s)
}

func (someOtherService) Count(_ context.Context, s string) int {
	return len(s)
}

func DoSomething(svc StringService) {
	fmt.Println("Did something", reflect.TypeOf(svc))

	fmt.Println(svc.Count(nil, "something"))
	fmt.Println(svc.Uppercase(nil, "something"))
}

func main() {
	x := someService{}
	y := someOtherService{}

	DoSomething(x)
	DoSomething(y)
}*/
