package main

import "fmt"

func main() {
	nama := "zulfikar"

	for i := 0; i < len(nama); i++ {
		fmt.Println(string(nama[i]))
	}

	counter := 1
	for counter <= 10 {
		fmt.Println("Loop ke ", counter)
		counter++ // Jika tidak diberi ini akan loop infinite
	}

	slice := []string{"zul", "fikar", "is", "naen"}
	for index, nama := range slice { // Jika kita tidak menggunakan index atau nama, maka ganti dengan _
		fmt.Println("index", index, "=", nama)
	}

}
