package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	fmt.Println(today)
	fmt.Println(today.Date())
	fmt.Println(today.Year())
	fmt.Println(today.Month())
	fmt.Println(today.Day())
	fmt.Println(today.Hour())
	fmt.Println(today.Minute())
	fmt.Println(today.Second())
	fmt.Println(today.Location())
	fmt.Println(today.Weekday())
	fmt.Println(today.Zone())
	fmt.Println(today.UTC())

	layout := "2006-01-02"
	formattedToday, _ := time.Parse(layout, "1997-12-10")
	fmt.Println(formattedToday)

	myDate := time.Date(2022, time.December, 28, 10, 10, 10, 10, time.Local)
	fmt.Println(myDate)

}
