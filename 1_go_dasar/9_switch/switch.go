package main

import "fmt"

func main() {
	nama := "zulfikar"

	switch nama {
	case "zul":
		fmt.Println("zul")
	case "fikar":
		fmt.Println("fikar")
	default:
		fmt.Println(nama)
	}

	// Switch Short Statement
	switch length := len(nama); length > 2 {
	case true:
		fmt.Println("Kepanjangan")
	case false:
		fmt.Println("Pendek")
	}

	// Switch Tanpa Kondisi
	length := len(nama)
	switch {
	case length > 8:
		fmt.Println("Panjang")
	case length < 8:
		fmt.Println("Pendek")
	default:
		fmt.Println("Sudah benar")
	}
}
