package main

import "fmt"

func main() {
	type NoKTP string
	type Success bool

	var noKTP NoKTP = "09123723423"
	var statusKarir Success = true
	fmt.Println(noKTP)
	fmt.Println(statusKarir)
}
