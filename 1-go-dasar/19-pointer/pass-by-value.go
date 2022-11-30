package main

import "fmt"

func main() {
	// Pass By Value
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := address1 // nilai dari address1 di copy ke address 2, sehingga ada 2 data dan 2 lokasi memori yang berbeda

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah, jika di bahasa program lain ini berubah
	fmt.Println(address2)
}

type Address struct {
	City, Province, Country string
}
