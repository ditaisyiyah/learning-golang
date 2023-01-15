package main

import "fmt"

/*
by default, all data movements in Go are pass by value
except slice, it always refers to an array

there is no mutable or immutable concept in Go
because it depends the pointer and addressof usage

pointer is a data type
it stores an address of another data

* => pointer
& => addressof

*/

type Book struct {
	Type  string
	Pages int16
}

func main() {
	novel := Book{
		Type:  "Novel",
		Pages: 100,
	}

	// 1. pass by value
	fmt.Println("==== pass by value")
	novel1 := novel
	novel1.Pages = 200

	fmt.Println("novel1 :", novel1)
	fmt.Println("novel :", novel) // novel do not change

	// 2. pass by reference
	fmt.Println("==== pass by reference")
	// novel2 := &novel // short declaration
	var novel2 *Book = &novel
	// novel2 points toward novel
	// novel2 is a "pointer"
	// novel2 stores "address of" novel

	novel2.Pages = 300

	fmt.Println("novel2 :", novel2)
	fmt.Println("novel :", novel) // novel change

	// 3. reassign to this new (address of) struct
	fmt.Println("==== reassign to new struct")
	novel2 = &Book{
		Type:  "Novel2 reassigned",
		Pages: 330,
	}

	fmt.Println("novel2 reassigned :", novel2)
	fmt.Println("novel :", novel) // novel do not change

	// 4. pointing all the structs that point to address of novel
	// to this new (address of) struct
	fmt.Println("==== pointing to new struct")
	var novel3 *Book = &novel
	var novel4 *Book = &novel
	// var novel4 *Book = novel3

	// (*) acts like an operator,
	// to access/change the ori value w/out accessing the ori variable
	*novel3 = Book{
		Type:  "Magazine",
		Pages: 20,
	}
	// novel = Book{
	// 	Type:  "Magazine",
	// 	Pages: 20,
	// }

	fmt.Println("novel3:", novel3)
	fmt.Println("novel4 :", novel4)
	fmt.Println("novel :", novel)

	// explore
	myString := "dita"
	var pointerMyString *string = &myString

	// value of pointerMyString, stores address of myString
	fmt.Println(pointerMyString) // 0xc000054290

	// address of pointerMyString
	fmt.Println(&pointerMyString) // 0xc000006030

	// referred value of pointerMyString
	fmt.Println(*pointerMyString) // dita

	// address of (referred value of pointerMyString)
	fmt.Println(&*pointerMyString) // 0xc000054290

	// value of (address of pointerMyString)
	fmt.Println(*&pointerMyString) // 0xc000054290

}
