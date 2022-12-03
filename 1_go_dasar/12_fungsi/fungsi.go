package main

import "fmt"

func main() {
	fungsiBiasa()
	fungsiParameter("hallo", 23, false)
	fungsiParameterSamaTipe("zul", "fikar", 2, 3)
	fmt.Println(fungsiBalikan())
	var x, str = fungsiMultiBalikan()
	var x2, _ = fungsiMultiBalikan()
	fmt.Println(x, str)
	fmt.Println(x2)

	var a, b = fungsiMultiNamaBalikan()
	c, d := fungsiMultiNamaBalikan()
	fmt.Println(a, b)
	fmt.Println(c, d)
}

func fungsiBiasa() {
	fmt.Println("Ini fungsi biasa")
}

func fungsiParameter(param1 string, param2 int, param3 bool) {
	fmt.Println(param1, param2, param3)
}

func fungsiParameterSamaTipe(param1, param2 string, test, test2 int) {
	fmt.Println(param1, param2, test, test2)
}

func fungsiBalikan() string {
	return "Hallo"
}

func fungsiMultiBalikan() (string, int) {
	return "umur", 23
}

func fungsiMultiNamaBalikan() (a string, b int) {
	a = "hahaha"
	b = 123
	return
}
