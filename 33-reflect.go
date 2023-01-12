package main

import (
	"fmt"
	"reflect"
)

type (
	User struct {
		Name string `required:"true" alias:"ajun"`
		Age  int    `required:"true" min:"1"`
		// Name string
		// Age  int
	}
)

func isValid(data interface{}) bool {
	dataReflect := reflect.TypeOf(data)

	for i := 0; i < dataReflect.NumField(); i++ {
		field := dataReflect.Field(i)

		// ok => to check the existence of key
		// if value, ok := field.Tag.Lookup("required"); !ok {
		// 	fmt.Println(value, ok)
		// 	return false
		// }

		if field.Tag.Get("required") == "true" {
			fmt.Println(reflect.ValueOf(data).Field(i).Interface())
			if reflect.ValueOf(data).Field(i).Interface() == "" {
				return false
			}
		}

	}
	return true
}

func main() {
	user := User{
		Name: "Juned",
		Age:  1,
	}

	userReflect := reflect.TypeOf(user)
	fmt.Println(userReflect.Field(0))

	// fmt.Println(userReflect.Field(0).Name)
	// fmt.Println(userReflect.Field(0).PkgPath)
	// fmt.Println(userReflect.Field(0).Type)
	// fmt.Println(userReflect.Field(0).Tag)
	// fmt.Println(userReflect.Field(0).Offset)
	// fmt.Println(userReflect.Field(0).Index)
	// fmt.Println(userReflect.Field(0).Anonymous)

	fmt.Println(userReflect.Field(0).Tag.Get("required"))
	fmt.Println(userReflect.Field(1).Tag.Get("min"))

	fmt.Println(userReflect.Field(1).Tag.Lookup("required"))

	fmt.Println(isValid(user))
}
