package main

import "fmt"

func main() {
	fmt.Println(fungsiFaktorial(3))
}

func fungsiFaktorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * fungsiFaktorial(value-1)
	}
}
