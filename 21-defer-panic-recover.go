package main

import (
	"fmt"
)

func endApp() {
	fmt.Println("Application is ending...")
	// catch the error at the end of app
	// in order to prevent application stop working
	recover()
}

func isSumOdd(number1, number2 int8) {
	// nomatter, defer always be executed
	// the prerequisite is must be declare at the start
	defer endApp()

	result := number1 + number2
	if result%2 == 0 {
		panic("The sum result is not an odd number")
	}
	// if an error happen or by design throw an error
	// code after that line will not be executed
	fmt.Println("The sum result :", result)
}

func main() {
	fmt.Println("Application is starting...")

	isSumOdd(8, 9)

	fmt.Println("You must recover the error to make this printed")
}
