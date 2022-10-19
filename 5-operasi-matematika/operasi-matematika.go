package main

import "fmt"

func main() {
	var data1 = 10
	var data2 int = 23
	var logic bool = false

	var result = data1 + data2
	result += 30
	var result1 = data1 - data2
	result1 -= 12
	var result2 = data1 * data2
	result2 *= 2
	var result3 = data1 / data2
	result2 /= 1
	var result4 = data1 % data2
	result4 %= 2

	var test int = 0
	test++
	result1--
	fmt.Println(!logic)

	fmt.Println(result)
	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(test)
}
