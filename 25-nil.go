package main

import (
	"fmt"
)

/*
nil is like null in other language
could be set into interface, function, map, slice, pointer, and channel

unintizlized data will be set to its default value, for instance:
interface => no field/property
number => 0
string =>
boolean => false

*/

func isEmpty(data map[string]int8) bool {
	for _, value := range data {
		if value != 0 {
			return true
		}
	}
	return false
}

func isNil(data map[string]int8) {
	if data == nil {
		fmt.Println("data is nil")
	} else {
		fmt.Println(data)
	}
}

func main() {
	var unInitMap map[string]int8
	fmt.Println(unInitMap)

	myMap := map[string]int8{
		"satu": 1,
		"dua":  2,
		"tiga": 3,
	}
	fmt.Println(myMap)

	fmt.Println(isEmpty(unInitMap))
	fmt.Println(isEmpty(myMap))

	var initNilMap map[string]int8 = nil
	isNil(initNilMap)

	// thus, init with nil to check the entirely emptiness of data
	// otherwise must be check one by one

}
