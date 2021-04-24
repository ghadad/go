package main

import (
	"fmt"
)

func main() {
	greetUser("John")      // prints "Hello, John"
	fmt.Println(add10(12)) // prints 22
}

func greetUser(user string) {
	fmt.Printf("Hello, %s\n", user)
}

func add10(value int) int {
	return value + 10
}
