package main

import (
	"fmt"
)

/*

slice is part of array, it refers to an array
slice is made pass by reference
except made by copy()

contains pointer, length, and capacity

*/

func main() {
	days := [7]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	// create a slice from an explicit array
	weekdays := days[:5]
	weekend := days[5:]

	fmt.Println("weekdays :", weekdays)
	fmt.Println("total days :", len(weekdays))

	fmt.Println("weekend :", weekend)
	fmt.Println("total days :", len(weekend))

	// capacity is length of array starts from pointer
	sliceA := days[2:4]
	sliceB := days[4:6]
	fmt.Println(cap(sliceA), sliceA) // (7 - 2)
	fmt.Println(cap(sliceB), sliceB) // (7 - 4)

	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	firstQuartal := months[:4]
	// fmt.Println(firstQuartal)

	// change slice element => change the array it is referred
	firstQuartal[1] = "LARAS"
	// fmt.Println(firstQuartal)
	// fmt.Println(months)

	// change array element => change the all slices that refersit refer
	months[1] = "February"
	// fmt.Println(firstQuartal)
	// fmt.Println(months)

	// to add an element into a slice,  must create a new different slice
	fmt.Println(len(firstQuartal))
	fmt.Println(cap(firstQuartal))
	newFirstQuartal := append(firstQuartal, "AfterApril")
	fmt.Println(newFirstQuartal)
	fmt.Println(firstQuartal) // doesn't change the initial slice
	fmt.Println(months)       // but change the original array

	// different case if there is NO capacity
	lastQuartal := months[8:]
	fmt.Println(len(lastQuartal))
	fmt.Println(cap(lastQuartal))

	newLastQuartal := append(lastQuartal, "AfterDesember")
	fmt.Println(newLastQuartal) // newLastQuartal refers to a new array in the behind, not months
	fmt.Println(lastQuartal)    // deosn't change the initial slice
	fmt.Println(months)         // doesn't change the original array

	// initialize a slice (from implcit array)
	// var copiedSlice []string = make([]string, len(firstQuartal), cap(firstQuartal))
	copiedSlice := make([]string, len(firstQuartal), cap(firstQuartal))

	// copy a slice => pass by value
	copy(copiedSlice, firstQuartal)
	fmt.Println(copiedSlice)

	// create a full slice
	// var aSlice []int8 = []int8{1, 2, 3, 4}
	aSlice := []int8{1, 2, 3, 4}

	// length of an array is a must
	// use ... if do not know the exact number
	anArray := [...]int8{1, 2, 3, 4}

	fmt.Println(anArray, aSlice)
}

/*
Slice functions:
make([]type, len, cap) 		=> to make a slice from an implicit array
array[low:high]						=> to make a slice from an explicit array
													=> start index low until before index high
copy(toSlice, fromSlice) 	=> to copy a slice to a new slice
len(slice) 								=> to get length of slice
cap(slice) 								=> to get element at index number
slice[index] 							=> to get element at index number
slice[index] = value 			=> to replace element at index number
append(slice, value)			=> to append/replace an element in new slice/referred array

*/
