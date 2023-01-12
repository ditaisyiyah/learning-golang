package main

import (
	"fmt"
)

/*
	5. function type declaration
	6. function as a paramter
*/

// function type declaration, alias for function data type
type (
	CheckCountry func(string) bool
	CheckAge     func(int8) bool
)

func checkCountry(countryName string) bool {
	fromIndonesia := countryName == "indonesia"
	if fromIndonesia {
		fmt.Println("from indonesia ✔")
	} else {
		fmt.Println("from indonesia ❌")
	}

	return fromIndonesia
}

func checkAge(age int8) bool {
	above18 := age > 18
	if above18 {
		fmt.Println("above 18 y.o. ✔")
	} else {
		fmt.Println("above 18 y.o. ❌")
	}

	return above18
}

func checkUser(country string, checkCountry CheckCountry, age int8, checkAge CheckAge) bool {
	isCountryValid := checkCountry(country)
	isAgeValid := checkAge(age)

	fmt.Println("Country validity :", isCountryValid)
	fmt.Println("Age validity :", isAgeValid)

	return isCountryValid && isAgeValid
}

func checkCurrentUser(age int8, checkAge CheckAge) {
	isPermitted := checkAge(age)
	if isPermitted {
		fmt.Println("Happy watching!")
	} else {
		fmt.Println("You are not allowed to watch this")
	}
}

func main() {
	var (
		fullName      = "Dita Larasati"
		country       = "Russia"
		age      int8 = 25
	)

	// function as parameter
	isValidUser := checkUser(country, checkCountry, age, checkAge)
	fmt.Println(fullName, "validity :", isValidUser)

	checkCurrentUser(age, checkAge)
	checkCurrentUser(age, func(age int8) bool {
		above18 := age > 18
		if above18 {
			fmt.Println("above 18 y.o. ✔")
		} else {
			fmt.Println("above 18 y.o. ❌")
		}

		return above18
	})

}
