package main

import "fmt"

func main() {
	// number conversion
	var (
		number32 int32 = 130
		number16 int16 = int16(number32)
		number8  int8  = int8(number32)
	)

	fmt.Println(number32, number16)
	fmt.Println(number8)
	// due to max number of int8 is 127, it goes down to min number
	// so, be careful!

	// string conversion
	var (
		myName      = "Larasati"
		firstLetter = myName[0]
		// firstLetter byte = myName[0]
		// firstLetter int8 = myName[0]
		l = string(firstLetter)
	)

	fmt.Println("first letter of my last name is =", l)

	// curious
	firstByte := -128
	lastByte := 127

	fmt.Println(string(firstByte), string(lastByte))
	fmt.Println(string(128), string(129), string(200))

}
