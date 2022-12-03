// Pointer kemampuan membuat reference ke lokasi data di memori yang sama, tanpa menduplikasi data yang sudah ada
// atau bisa kita bilang Pass By Reference

package main

import "fmt"

func main() {
	var address1 Address = Address{"Bogor", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // nilai dari address1 tidak di copy ke address 2, sehingga address2 akan mereference ke address1
	var address3 *Address = new(Address)
	address3.City = "Depok"
	fmt.Println(*address3)

	address2.City = "Bandung"

	fmt.Println(address1) // address1 akan berubah
	fmt.Println(*address2)

	fmt.Println("###############################")

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // Address1 tidak berubah
	// fmt.Println(*address2)

	// fmt.Println("###############################")

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // Address1 ikut berubah
	fmt.Println(*address2)
}

type Address struct {
	City, Province, Country string
}
