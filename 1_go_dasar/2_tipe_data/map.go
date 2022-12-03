package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":   "zulfikar",
		"alamat": "bandung",
	}

	person["nama"] = "Zulfikar"

	fmt.Println(person["nama"])
	fmt.Println(person["alamat"])
	fmt.Println(person["gaada"])

	fmt.Println(len(person))
	var person2 map[string]string = make(map[string]string)
	person2["nama"] = "Arief"
	person2["alamat"] = "Bandung"
	fmt.Println(person2)
	fmt.Println(len(person2))

	delete(person2, "nama")

	fmt.Println(person2)
	fmt.Println(len(person2))
}
