package main

import (
	"fmt"
)

func main() {
	var (
		// today   = "Wednesday"
		tomorrow = "Wednesday"
		message  string
		message2 string
	)

	// switch short statement
	switch today := "Wednesday"; today {
	case "Monday":
		message = "Happy money day!"
	case "Tuesday", "Wednesday", "Thursday":
		message = "Ganbatte!"
	case "Friday":
		message = "Weekend is coming!!"
	case "Saturday":
		message = "Happy Weekend!"
	case "Sunday":
		message = "Tomorrow is Moday, remember!"
	default:
		message = "Happy whatt?!?"
	}

	fmt.Println(message)

	// do switch AS IF if-else-if
	switch {
	case tomorrow == "Monday":
		message2 = "Happy money day!"
	case tomorrow == "Saturday" || tomorrow == "Sunday":
		message2 = "Happy Weekend!"
	default:
		message2 = "Have a great day!"
	}

	fmt.Println(message2)
}
