package main

import "fmt"

func main() {
	total := sumAll("hai", 5, 5, 5, 5, 5)
	total2 := sumAll("hai")
	fmt.Println(total2)
	fmt.Println(total)

	dataSlice := []int{3, 3, 3, 3, 3}
	total3 := sumAll("zul", dataSlice...)
	fmt.Println(total3)
}

func sumAll(test string, numbers ...int) int {
	fmt.Println(test)
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
