package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		counter := 1 // Jika kita tidak mendefinisikan ini, maka counter akan diambil dari scope luar increment function
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}
