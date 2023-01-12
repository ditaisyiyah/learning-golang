package main

import (
	"fmt"
)

func main() {
	var (
		score   = 80
		message string
	)

	if score < 50 {
		message = "You got D!"
	} else if score < 75 {
		message = "You got C!"
	} else if score < 90 {
		message = "You got B!"
	} else if score <= 100 {
		message = "You got A!"
	}

	fmt.Println(message)

	// if short statement
	if today := "Saturda"; today == "Saturday" || today == "Monday" {
		fmt.Println("Happy weekend!")
	} else if today == "Sunday" || today == "Tuesday" || today == "Wednesday" || today == "Thursday" || today == "Friday" {
		fmt.Println("Have a nice day!")
	} else {
		fmt.Println("Hmm")
	}

}
