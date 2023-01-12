package main

import (
	"fmt"
)

func main() {
	const (
		standardScore   = 70
		standardPresent = 90
	)

	var (
		scoreA     = 72
		presentA   = 90
		curiosityA = false

		scoreB     = 68
		presentB   = 100
		curiosityB = true
	)

	isAPassed := scoreA >= standardScore && presentA >= standardPresent
	isBPassed := scoreB >= standardScore && presentB >= standardPresent

	fmt.Println(isAPassed)
	fmt.Println(isBPassed)

	isAPassed = isAPassed || curiosityA
	isBPassed = isBPassed || curiosityB
	fmt.Println(isAPassed)
	fmt.Println(isBPassed)

	// there is no truthy or falsy value in go
	// box := "will print this" && "not print this"
	// fmt.Println(box)

}

/*
Logical operators operates boolean data and result boolean
&&    logical AND, both must be true
||    logical OR, one of both must be true
!     negation, change the boolean value

*/
