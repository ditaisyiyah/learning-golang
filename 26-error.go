package main

import (
	"errors"
	"fmt"
)

/*
error is an interface
*/

type (
	// TODO: buat tiruan interface error yg dipake sbg return type
	MyError interface {
		MyError()
	}
	Singer struct {
		Name, Voice string
	}
)

// func printError(myError MyError) {
// 	fmt.Println("ini error :", myError.MyError())
// }

// func (s Singer) MyError() string {
// 	return "tes error"
// }

func multiply(number1 int8, number2 int8) (int16, error) {
	if number1 == 0 || number2 == 0 {
		return 0, errors.New("dikali 0? hasilnya udah jelas 0 dong")
	} else {
		return int16(number1) * int16(number2), nil
	}
}

func main() {
	// person := Singer{
	// 	Name:  "Adi",
	// 	Voice: "duung",
	// }

	// printError(person)

	var myError error = errors.New("ini pesan error")
	// fmt.Println(myError) // WHY dont just use this intead below?
	fmt.Println(myError.Error())

	result, err := multiply(9, 9)
	if err == nil {
		fmt.Println("result :", result)
	} else {
		fmt.Println("error :", err)
	}

}
