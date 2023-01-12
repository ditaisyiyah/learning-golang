package main

import (
	"fmt"
)

/*
interface is a contract
interface is kind of data type
mostly its fields are a func or method (void or nonvoid)
*/

type (
	Voice interface {
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

// sayHello use Voice interface as a contract
func sayHello(voice Voice) {
	fmt.Println("Hello,", voice.GetVoice())
}

func (a Animal) GetVoice() string {
	return a.Voice
}

func (a Animal) GetName() string {
	return a.Name
}

func greeting(voice Voice) {
	fmt.Println(voice.GetVoice(), voice.GetName())
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

	// any struct that want to invoke a func that has a contract with interface
	// that struct must have all the prerequisite that interface has
	sayHello(cat)

	person := Human{
		Name:  "Laras",
		Voice: "Haii",
	}

	greeting(person)

}
