package main

import (
	"fmt"
	"reflect"
)

const pi = 3.14

var course string = "Golang"

var (
	fn, ln string
	country = "Costa Rica"
	age int = 28
)

func main() {
	ageS := "123"
	fmt.Printf("Hello World(%s): %s %s %d\n", course, fn, ln, age)
	fmt.Println("Age is of type", reflect.TypeOf(age))
	fmt.Println("AgeS is of type", reflect.TypeOf(ageS))

	defineShortHandVariable()
	usePointers()

	funcUsingPointers(&age)
	fmt.Println("The age was changed", age)

	if ageS < "132" {
		fmt.Println("Coerced?")
	}

}

func defineShortHandVariable() {
	varMain := "Main"
	fmt.Println(varMain)
}

func usePointers() {
	text := "Hello World"
	pointer := &text

	fmt.Println("The mainVar address is", &pointer , "and the value is", *pointer)
}

func funcUsingPointers(age *int) int {
	*age = 29
	return *age
}