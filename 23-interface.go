package main

import (
	"fmt"
)

/*
An interface type is defined as a set of method signatures.
A value of interface type CAN hold any value THAT implements THOSE methods.
Interface is abstract, not allowed to create an instance of the interface.
*/

type (
	Intro interface {
		GetVoice() string
		GetName() string
	}

	Animal struct {
		Name, Voice string
		Feet        int8
	}

	Human struct {
		Name, Voice string
	}
)

// sayHello use Intro interface as a contract
func sayHello(intro Intro) {
	fmt.Println("Hello,", intro.GetVoice())
}

func (a Animal) GetVoice() string {
	return a.Voice
}

func (a Animal) GetName() string {
	return a.Name
}

func greeting(intro Intro) {
	fmt.Println(intro.GetVoice(), intro.GetName())
}

func (h Human) GetVoice() string {
	return h.Voice
}

func (h Human) GetName() string {
	return h.Name
}

func main() {
	cat := Animal{
		Name:  "Osa",
		Voice: "Meoong",
		Feet:  4,
	}

	catVoice := cat.GetVoice()
	catName := cat.GetName()
	fmt.Println(catVoice, catName)

	// any struct that want to invoke a func that has a contract with interface
	// that struct must have all the prerequisite that interface has
	sayHello(cat)

	person := Human{
		Name:  "Laras",
		Voice: "Haii",
	}

	greeting(person)

	var myInterface Intro
	result := myInterface.GetName()
	fmt.Println(result, "just want you to know, this could be donw")

}
