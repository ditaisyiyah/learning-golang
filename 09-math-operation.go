package main

import (
	"fmt"
)

func main() {
	var (
		height = 100
		width  = 20
	)

	area := height * width
	circumference := (height + width) * 2

	fmt.Println("area", area)
	fmt.Println("circumference", circumference)

	// augmented assignments
	fmt.Println("augmented assignments")
	var (
		x = 5
		y = 3
	)

	x = x / 2
	fmt.Println(x)
	x /= 3 // quotatient results truncated numbers
	fmt.Println(x)

	y = y % 2
	fmt.Println(y)
	y %= 3
	fmt.Println(y)
	// remainder always result integers

	// unary operation
	fmt.Println("unary operation")
	var (
		a = 2
		b = 4
		c = true
	)

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	b = -b
	fmt.Println(b)
	b = -b
	fmt.Println(b)
	// those mean 0 - b
	// so, b = b means 0 + b, do not result positive

	c = !c
	fmt.Println(c)
	c = !c
	fmt.Println(c)

}

/*
Arithmetic operators
+    sum                    integers, floats, complex values, strings (concatenation)
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << integer >= 0
>>   right shift            integer >> integer >= 0

Precedence operator
5	 	*  /  %  <<  >>  &  &^
4	 	+  -  |  ^
3	 	==  !=  <  <=  >  >=
2	 	&&
1	 	||

*/
