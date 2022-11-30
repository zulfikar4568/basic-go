// Walau method akan menempel di struct, tapi sebenernya data struct yang di akses di dalam method adalah pass by value
// Sangat di rekomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu men duplikasi ketika memanggil method

package main

import "fmt"

func main() {
	zul := Man{"Zul"}
	zul.Married()
	fmt.Println(zul) // tidak berubah

	zul.MarriedPointer()
	fmt.Println(zul) // akan berubah
}

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr " + man.Name
}
