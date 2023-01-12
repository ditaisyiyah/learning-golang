package main

import (
	"fmt"
)

/*
struct is a data type of multiple fields,
could have uneven data type

as if an object in other languages
*/

type People struct {
	Name, Sex string
	Age       int8
}

func sayLove(fromName string, p People) {
	fmt.Println(fromName, "loves", p.Name)
}

func (p People) sayHate(fromName string) {
	fmt.Println(fromName, "hates", p.Name)
}

func main() {
	var person People

	person.Name = "Laras"
	person.Sex = "Female"
	person.Age = 25

	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Sex)
	fmt.Println(person.Age)

	employee := People{
		Name: "Adi",
		Sex:  "Male",
		Age:  31,
	}
	fmt.Println(employee)

	parent := People{"Jono", "Male", 60} // not recommended
	fmt.Println(parent)

	// struct as parameter
	sayLove("Adi", person)

	// struct function or method
	// for aiming the func to be owned only by the struct
	person.sayHate("Adi")
	// sayHate("Adi") // wrong way ya

}
