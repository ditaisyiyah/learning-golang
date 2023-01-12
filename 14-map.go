package main

import (
	"fmt"
)

func main() {
	// create a full map
	// var myMap map[string]string = map[string]string{
	// 	"name":   "Dita",
	// 	"status": "Married",
	// }
	myMap := map[string]string{
		"firstName": "Dita",
		"lastName":  "Larasati",
	}

	fmt.Println(myMap)

	// initialize a map
	// var aMap map[string]int8 = make(map[string]int8)
	aMap := make(map[string]int16)
	fmt.Println(aMap)

	aMap["birthYear"] = 1997
	aMap["age"] = 25

	fmt.Println(aMap)
	fmt.Println(len(aMap))
	fmt.Println(aMap["age"])
	fmt.Println(aMap["birthYear"])

	delete(aMap, "birthYear")

	fmt.Println(aMap)
	fmt.Println(len(aMap))
	fmt.Println(aMap["age"])
	fmt.Println(aMap["birthYear"]) // not causes error, it results 0

}

/*
Map functions:
make(map[keyType]valueType) => to make a map
len(map) 										=> to get length of map
map[index] 									=> to get element by keyname
map[index] = value		 			=> to replace element by keyname
delete(map, key) 						=> to delete element by keyname

*/
