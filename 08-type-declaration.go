package main

import "fmt"

func main() {
	type (
		Name    string
		Married bool
		// like built-in data type aliases
		// such as byte, rune, int, and uint
	)

	var fullName Name = "Dita Larasati"
	var myStatus Married = false
	// myStatus Married := false // cannot use shorthand way

	fmt.Println(fullName, myStatus)
}
