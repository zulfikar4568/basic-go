package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var nama string = "Zulfikar"
	var z byte = nama[0]
	var zString string = string(z)

	fmt.Println(zString)

}
