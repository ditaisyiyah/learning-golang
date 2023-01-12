package main

import "fmt"

func main() {
	// var keyword is variable declaration
	var myName string = "Laras"

	// implicit data type
	var sex = "Female"

	// shorthand of var (and type implicitly)
	age := 25

	fmt.Println(myName, sex, age)

	// reassign
	myName = "Namjoon"
	sex = "Male"
	age = 28

	fmt.Println(myName, sex, age)

	// cannot reassign to another data type
	// myName = 12
	// sex = 20
	// age = true

	// multiple var variable declartion
	var (
		firstName = "Dita"
		lastName  = "Larasati"
		fullName  string
	)

	fullName = firstName + lastName

	fmt.Println(firstName, lastName, fullName)

}
