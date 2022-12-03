// Interface tipe data abstract, dan digunakan hanya untuk kontrak

package main

import "fmt"

func main() {
	person := Person{Name: "Zulfikar"}
	animal := Animal{Name: "Ojan kucing"}
	SayHello(person)
	SayHello(animal)
}

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println(hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name + " Person"
}

type Animal struct {
	Name string
}

func (person Animal) GetName() string {
	return person.Name + " Animal"
}
