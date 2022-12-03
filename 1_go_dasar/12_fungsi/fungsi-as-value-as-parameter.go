package main

import "fmt"

func main() {
	cetak := fungsiValue
	fmt.Println(cetak("zulfikar"))

	fungsiAsParameter("zulfikar", toxicFilter)
}

func fungsiValue(name string) string {
	return "cetak " + name
}

func fungsiAsParameter(nama string, filter func(string) string) {
	fmt.Println("Hallo", filter(nama))
}

// Menggunakan Type
type TFilter func(string) string

func fungsiAsParameterUsingType(nama string, filter TFilter) {
	fmt.Println("Hallo", filter(nama))
}

func toxicFilter(nama string) string {
	if nama == "Sue" {
		return "***"
	} else {
		return nama
	}
}
