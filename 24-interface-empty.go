package main

import (
	"fmt"
)

/*
type assertion => force or explicitly conversion data type
usually used when facing an empty interface
*/

type (
	Kosong interface {
	}

	Sapien struct {
		Name, Voice string
	}
)

func printFreely(kosong Kosong) {
	fmt.Println("I use empty interface as a contract, which means, no contract")
}

func returnAny(kosong Kosong) Kosong {
	if kosong == 1 {
		return "one"
	} else if kosong == "two" {
		return 2
	} else {
		return false
	}
}

func main() {
	person := Sapien{
		Name:  "Laras",
		Voice: "Haii",
	}

	printFreely(person)
	printFreely(1)

	var anyData Kosong = returnAny(1)
	fmt.Println(anyData)

	var anyData2 interface{} = returnAny("one")
	fmt.Println(anyData2)

	// by here, Kosong means interface{}
	// empty interface actually is implemented by fmt.Println()

	// error due to not only return a string
	// var anyData string = returnAny(1)

	// do this if you 100% sure that it will return particular data type
	var stringData string = (returnAny(1)).(string)
	fmt.Println(stringData)

	// this will be panic
	var stringDataError int8 = returnAny(1).(int8)
	fmt.Println(stringDataError)

	// use switch case to handle the panic such above
	switch returnAny(1).(type) {
	case int8:
		fmt.Println("do what you wanna do with by knowing this an integer")
	case string:
		fmt.Println("do what you wanna do with by knowing this an string")
	default:
		fmt.Println("do what you wanna do with by knowing this data type not what you want")
	}

}
