package main

import "fmt"

func main() {
	// data slice mempunyai 3 data, pointer, length, capacity
	bulan := [...]string{
		"jan",
		"feb",
		"maret",
		"april",
		"mei",
		"jun",
		"jul",
		"agus",
		"sept",
		"okt",
		"nov",
		"des",
	}
	slice := bulan[3:5]
	slice2 := bulan[3:]
	fmt.Println(slice2)
	fmt.Println((bulan))
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	bulanBaru := append(slice2, "tambahan")
	fmt.Println(bulanBaru)

	bulan[3] = "April"
	fmt.Println(slice)

	// from scratch

	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "zul"
	sliceBaru[1] = "fikar"

	fmt.Println(sliceBaru)
	fmt.Println(len(sliceBaru))
	fmt.Println(cap(sliceBaru))
}
