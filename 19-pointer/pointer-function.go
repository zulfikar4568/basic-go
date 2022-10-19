// Parameter di function secara default pass by value, artinya data akan di copy lalu di kirim ke func tersebut
// Jika kita mengubah data di dalam func, data yang asli tidak akan pernah berubah
// Sehingga variable akan menjadi aman, karena tidak bisa berubah
// Jika kita ingin mengubah dat aslinya juga maka gunakan pointer, gunakan operator *

package main

import "fmt"

func main() {
	var address Address = Address{"Bogor", "Jawa Barat", ""}
	ChangeAddressToIndonesia(address)
	fmt.Println(address) // tidak berubah

	var alamatPointer *Address = &address
	ChangeAddressToIndonesiaPointer(alamatPointer)
	fmt.Println(address) // data berubah
}

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndonesiaPointer(address *Address) {
	address.Country = "Indonesia"
}
