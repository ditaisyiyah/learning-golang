package main

import "fmt"

func main() {
	// const keyword is variable declaration (it's constant)
	const phi float32 = 3.14
	const heliumNumber int8 = 1

	// unreassignable!
	// phi = 2.17

	fmt.Println(phi, heliumNumber)

	// multiple const variable declaration
	const (
		ten   = 10
		black = "#000000"
		// unknown bool // the value of const must be initialized!
	)

	// it is okay if the const value not used yet

}
