package main

import (
	"fmt"
)

func main() {
	var (
		sembilan = 9
		sepuluh  = 10

		lowerDita = "dita"
		upperDita = "Dita"
	)

	fmt.Println(sepuluh > sembilan) // true

	// only compare to the same data type
	// no strictly or loosely comparison

	// fmt.Println(sepuluh < "10")
	// fmt.Println(sepuluh > "10")
	// fmt.Println(sepuluh == "10")

	fmt.Println("dita to Dita")
	fmt.Println(">", lowerDita > upperDita)   // true lexicographically
	fmt.Println("<", lowerDita < upperDita)   // false
	fmt.Println("==", lowerDita == upperDita) // false
	fmt.Println("!=", lowerDita != upperDita) // true alphabetically

	fmt.Println("dita to dita")
	fmt.Println(">", lowerDita > lowerDita)   // false
	fmt.Println("<", lowerDita < lowerDita)   // false
	fmt.Println("==", lowerDita == lowerDita) // true alphabetically
	fmt.Println("!=", lowerDita != lowerDita) // false

	fmt.Println("lexicograhically test")
	fmt.Println("ditA" < "dita")
	fmt.Println("A" < "B")
	fmt.Println("a" < "b")

}

/*
Comparison operators results boolean value
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal


*/
