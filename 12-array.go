package main

import (
	"fmt"
)

func main() {
	// initialize an array (length is a must)
	var myWords [3]string
	fmt.Println("empty array", myWords, "has length", len(myWords))

	myWords[0] = "ku"
	myWords[1] = "harus"
	myWords[2] = "merelakanmu"
	// myWords[3] = "?" // cannot change the length of array!

	// array be printed withou comma in between
	fmt.Println("filled array", myWords)

	// create a full array (length is a must)
	// var myNumbers [3]int8 = [3]int8{10, 12, 14}
	myNumbers := [3]int8{10, 12, 14}

	fmt.Println(myNumbers)

}

/*
Array functions:
len(array) 						=> to get length of array
array[index] 					=> to get element at index number
array[index] = value 	=> to replace element at index number

*/
