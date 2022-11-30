// fmt.Println(a ...interface[])
// panic(v interface[])
// recover() interface[]
// etc..
package main

import "fmt"

func main() {
	kosong := Ups(2)
	fmt.Println(kosong)
}

func Ups(data int) interface{} {
	// Bisa me return apa saja tipe datanya
	if data == 1 {
		return 1
	} else if data == 2 {
		return true
	} else {
		return "Ups"
	}
}
