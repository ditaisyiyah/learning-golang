package greeting

import "fmt"

// use uppercamelcase to let be accessed from outside package
func Greet() {
	fmt.Println("Hello from the outside~")
}

func greet() {
	fmt.Println("Hello you can't hear me~")
}
