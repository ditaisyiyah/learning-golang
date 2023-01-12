package main

import (
	"fmt"
	"sort"
)

type (
	Users []User

	User struct {
		Name string
		Age  int
	}
)

func (users Users) Len() int {
	return len(users)
}

func (users Users) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func (users Users) Less(i, j int) bool {
	return users[i].Age < users[j].Age
}

func main() {
	users := Users{
		{Name: "Dita", Age: 25},
		{Name: "Adi", Age: 31},
		{Name: "Juned", Age: 1},
	}
	fmt.Println(users)

	sort.Sort(users)
	fmt.Println(users)

}
