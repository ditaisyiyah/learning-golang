package main

import (
	"fmt"
)

func main() {
	num := 5      // initial statement
	for num > 0 { // condition
		fmt.Println("Iterate", num)
		num-- // post statement
	}

	// for short statement
	for num := 0; num < 5; num++ {
		fmt.Println("at number", num)
		if num == 3 {
			break // exit from iteration
		}

		if num%2 == 0 {
			continue // go to the next iteration
		}
		fmt.Println("it's odd number", num)

	}

	// for loop can iterate array and slice
	// for range can iterate array, slice, and map

	myArray := [3]string{"Dita", "Larasati", "Caem"}
	mySlice := myArray[:2]
	myMap := map[string]string{
		"firsName": "Dita Larasati",
		"jobTitle": "Software Engineer",
	}

	for i := 0; i < len(myArray); i++ {
		fmt.Println("myArray index", i, "has value", myArray[i])
	}
	for index, value := range myArray {
		fmt.Println("myArray index", index, "has value", value)
	}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
	for _, value := range mySlice { // use underscore if no need
		fmt.Println(value)
	}

	// to access a map, only can use for range
	for _, value := range myMap {
		fmt.Println(value)
	}

}

/*
There is no while loop or do-while loop
Use for range to iterate iterable data entirely
*/
