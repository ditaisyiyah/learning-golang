package main

import (
	"fmt"
)

/*
	3. variadic function
	4. named function (not recommended)
*/

// variadic function => varags as paramter
// argument will be in slice data type
func getTotal(numbers ...int8) int16 {
	var total int16
	for _, value := range numbers {
		total += int16(value)
	}

	return total
}

func main() {
	var (
		length int8 = 10
		width  int8 = 20
	)

	fmt.Println("length :", length)
	fmt.Println("width :", width)

	myNumbers := []int8{2, 4, 6, 8}
	myTotal := getTotal(myNumbers...) // spread the slice values

	fmt.Println("my total :", myTotal)

	// named function => func as a value, save in variable
	checkOdd := func(number int8) {
		isOdd := number%2 == 0
		if isOdd {
			fmt.Println(number, "is an even number")
		} else {
			fmt.Println(number, "is odd number")
		}
	}

	checkOdd(5)

	// var getSum func(...int8) int16 = getTotal
	getSum := getTotal
	total := getSum(length, width) // equals to getTotal(length, width)
	fmt.Println("total :", total)

}
