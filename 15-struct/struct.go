// struct adalah template untuk menggabungkan 0 atau lebih tipe data
// struct biasanya representasi data dalam program yang kita buat
// data di struct disimpan di field
// struct = kumpulan field

// struct = template data atau prototype
// kita bisa membuat object dari struct yang telat dibuat

package main

import "fmt"

type Customer struct {
	Nama, Alamat string
	Umur         int
}

func main() {
	var customer1 Customer
	customer1.Nama = "Zulfikar"
	customer1.Alamat = "Bandung"
	customer1.Umur = 23

	fmt.Println(customer1)

	// Struct Literal
	zul := Customer{
		Nama:   "Zulfikar",
		Alamat: "Cimahi",
		Umur:   12,
	}

	zul2 := Customer{"Zul", "Jkt", 12}

	fmt.Println(zul, zul2)
}
