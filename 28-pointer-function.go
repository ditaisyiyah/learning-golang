package main

import "fmt"

type Book struct {
	Type  string
	Pages int16
}

func changeTypeBook(book Book) {
	book.Type = "Magazine"
}

// due to a pointer always stores an address
// so the parameter is in a pointer type
// and the argument that will be passed in an address
// it's actually make sense, do not be confused by the operator XD
func changeTypeBookTrully(book *Book) {
	book.Type = "Magazine"
}

func main() {
	novel := Book{
		Type:  "Novel",
		Pages: 100,
	}

	fmt.Println(novel)

	// pass by value
	changeTypeBook(novel)
	fmt.Println(novel)

	// pass by reference or the address
	// addressOfNovel := &novel
	// changeTypeBookTrully(addressOfNovel)
	changeTypeBookTrully(&novel)
	fmt.Println(novel)

}
