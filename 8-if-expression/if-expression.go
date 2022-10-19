package main

import "fmt"

func main() {
	firstName := "zulfikar"
	lastName := "isnaen"

	if firstName == "Zulfikar" {
		fmt.Println("Hai Zulfikar")
	} else if firstName == "Zulfikar" && lastName == "isnaen" {
		fmt.Println("Hai " + firstName + " " + lastName)
	} else {
		fmt.Println(firstName)
	}

	// if short statement
	if length := len(firstName); length > 2 {
		fmt.Println("Panjang")
	} else {
		fmt.Println("Pendek")
	}
}
