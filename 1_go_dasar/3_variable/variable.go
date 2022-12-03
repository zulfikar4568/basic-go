package main

import "fmt"

func main() {
	var msg string
	msg = "Deklarasi tanpa inisialisasi"
	fmt.Println(msg)

	var msg2 string = "Deklarasi dengan inisialisasi"
	fmt.Println(msg2)

	var (
		firstName = "Zulfikar"
		lastName  = "Isnaen"
	)
	fmt.Println(firstName, lastName)

	var angka1, angka2 int = 1, 2 // Deklarasi multi variable
	fmt.Println(angka1, angka2)

	var msg3 = "Tipe data bisa apa saja"
	// msg3 = 123 // Akan error jika di isi number
	fmt.Println(msg3)

	foo := 42 //Otomatis tipe data nya, dan tipe akan selalu implicit
	foo = 123
	fmt.Println((foo))

	const constant = "Ini adalah konstanta"
	const const1, const2 = "phi1", "phi2"
	// const1 = "ubah" // Tidak bisa di ubah
	fmt.Println(constant)
	fmt.Println(const1, const2)

	const (
		_ = iota
		a
		b
		c = 1 << iota
		d
	)
	fmt.Println(a, b) // 1 2 (0 is skipped)
	fmt.Println(c, d) // 8 16 (2^3, 2^4)
}
