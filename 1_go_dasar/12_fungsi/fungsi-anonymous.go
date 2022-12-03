package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "sue"
	}

	registerUser("zul", blacklist)
	registerUser("sue", func(name string) bool {
		return name == "sue"
	})

}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blocked")
	} else {
		fmt.Println("Ok")
	}
}
