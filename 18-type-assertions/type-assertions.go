// Type Assertions kemampuan merubah tipe data menjadi tipe data yang di inginkan
// Fitur ini biasanya digunakan ketika bertemu dengan interface kosong

package main

import "fmt"

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // Panic
	// fmt.Println(resultInt)

	// Dengan bantuan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}

func random() interface{} {
	return "OK"
}
