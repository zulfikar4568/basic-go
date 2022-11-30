package main

import (
	"fmt"

	mypackage "go-hello/pkg_hello"
)

func main() {
	fmt.Println("Hello, Modules!")

	mypackage.PrintHello()
}
