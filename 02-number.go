package main

import "fmt"

func main() {
	fmt.Println("nomor satu =", 1)
	fmt.Println("nomor dua =", 2)
	fmt.Println("nomor tiga koma lima =", 3.5)
}

/*
Number data types: integer and floating point

integer: intx and uintx
where x could be 8, 16, 32, and 64
for example int8, the value range is -128 until 127
2^8 = 256, divided by 2 results 128
thus, minimal number is -128 and the maximal number is 127 (because of 0)

floating point: float32, float64, complex64, and complex128
float32: 1.18 x 10^-38 until 3.4 x 10^38
float64: 2.23 x 10^-308 until 1.8 x 10^308

aliases:
byte => uint8
rune => int32
int => int32 or int64 (depens on OS bit)
uint => uint32 or uint64 (depens on OS bit)

*/
