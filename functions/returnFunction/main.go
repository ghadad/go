package main

import (
	"fmt"
)

func main() {
	f := getFunction("John") // returns a function
	f()                      // prints "Hello, John"
}

func getFunction(name string) func() {
	return func() {
		fmt.Printf("Hello, %s", name)
	}
}
