package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"os"
	"strconv"
)

func MultiTable(number int) string {
	result := ""
	for i := 1; i <= 10; i++ {
		result = result + fmt.Sprintf("%d * 5 = %d \n", i, i*5)
	}
	return result
}

func main() {
	// package os //
	args := os.Args // to know the location of compile file
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error", err.Error())
	}

	// package container list //
	/*
		doubly linked list
		element is linked each other,
		but, the first and last element are linked to the nil
	*/

	// var myList *list.List = list.New()
	myList := list.New()
	// fmt.Println(myList)

	myList.PushFront("Dita")
	myList.PushFront("Sebelum Dita")
	myList.PushBack("Laras")
	myList.PushBack("Setelah Laras")

	// fmt.Println(myList)
	// fmt.Println("first el :\n", myList.Front())
	// fmt.Println("last el :\n", myList.Back())

	for el := myList.Front(); el != nil; el = el.Next() {
		// for el := myList.Back(); el != nil; el = el.Prev() {
		fmt.Println(el)
	}

	// fmt.Println(myList)
	// myList.Init()
	// fmt.Println(myList)

	// package container ring //
	/*
		circular list
		element is linked each other
		the first element linked to the last element, vice versa
	*/

	// var myRing *ring.Ring = ring.New(5)
	myRing := ring.New(5)
	// fmt.Println(myRing)

	// to init the ring value
	for i := 0; i < myRing.Len(); i++ {
		myRing.Value = "Data ke-" + strconv.FormatInt(int64(i), 10)
		myRing = myRing.Next() // reassign the ring header
	}

	// to iterate ring entirely
	myRing.Do(func(value interface{}) {
		fmt.Println(value)
	})

}
