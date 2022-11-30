package main

import "fmt"

type Customer struct {
	Nama, Alamat string
	Umur         int
}

func main() {
	zul := Customer{
		Nama:   "Zulfikar",
		Alamat: "Cimahi",
		Umur:   12,
	}
	zul.sayHello("Isnaen")
}

func (customer Customer) sayHello(param1 string) {
	fmt.Println("Nama saya adalah", customer.Nama, param1)
}
