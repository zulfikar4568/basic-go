// Object yang belum di inisialisasi datanya digolang maka nilainya akan di isi default value tergantung tipe datanya
// nil merupakan data kosong
// nil biasanya digunakan di beberapa tipe data interface, func, map, slice, pointer, channel

package main

import "fmt"

func main() {
	var person map[string]string = nil
	var kosong = NewMap("Zul")
	fmt.Println(person)
	fmt.Println(kosong)
}

func NewMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
}
