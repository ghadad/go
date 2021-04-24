package main

import (
	"fmt"
)

func main() {

	func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}("John") // prints "Hello, John"

	func(name string) {
		fmt.Printf("Hello, %s", name+name)
	}("John") // prints "Hello, John"
}
