// Interger & Floating Point
/*
Nilai positif dan negatif
int8 -128 to 127
int16 -2^15 to 2^15 -1
int32 -2^31 to 2^31 -1
int64 -2^63 to 2^63 -1

Nilai positif
uint8 0 to 255
uint16 0 to 2^16 -1
uint32 0 to 2^32 -1
uint64 0 to 2^64 -1

float32 1.18x10^-38 to 3.4x10^38
float64 2.23x10^-308 to 1.80x10^308
complex64
complex128

Alias
byte uint8
rune int32
int minimal int32 mengikuti OS, jika OS 64 maka akan int64
unit minimal uint32 mengikuti OS, jika OS 64 maka akan uint64


*/

package main

import "fmt"

func main() {
	fmt.Println(1, 2, 3.4, 5.4498)
}
