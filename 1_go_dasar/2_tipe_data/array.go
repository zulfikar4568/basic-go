package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "zul"
	nama[1] = "fikar"
	nama[2] = "isnaen"
	// nama[3] = "out index" // Akan Error

	var data = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(nama)
	fmt.Println(data)

	fmt.Println(len(nama))
	fmt.Println(data[2])
	data[2] = 4
	fmt.Println(data[2])
}
