package main

import (
	"fmt"
)

/*
	1. non-void function
	2. named return function for return multiple values
*/

func getArea(length int8, width int8) int16 {
	if length == width {
		fmt.Println("Area of square :")
	} else {
		fmt.Println("Area of rectangular :")
	}

	return int16(length) * int16(width)
	// return int16(length * width) // this is wrong
}

func getCircumference(sides []int8) int16 {
	fmt.Println("Circumference of flat", len(sides), "sides :")

	var circumference int16
	for _, value := range sides {
		circumference += int16(value)
	}

	return circumference
}

// only in GO
// implement named return function for return multiple values
func getAreaAndCircumference(length, width int8, sides []int8) (area, circumference int16) {
	area = int16(length) * int16(width)

	for _, value := range sides {
		circumference += int16(value)
	}

	return
	// implicitly return these below
	// return area, circumference
}

func main() {
	var (
		length int8 = 10
		width  int8 = 20
	)

	fmt.Println("length :", length)
	fmt.Println("width :", width)

	var area int16 = getArea(length, width)
	fmt.Println(area)

	sides := []int8{length, width}
	var circumference int16 = getCircumference(sides)
	fmt.Println(circumference)

	var area2, circumference2 int16 = getAreaAndCircumference(length, width, sides)
	fmt.Println("area :", area2)
	fmt.Println("circumference :", circumference2)
}
