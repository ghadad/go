package main

import (
	"fmt"
)

func main() {
	showAge("Paul", 22)                     // prints "Paul is 22 years old."
	fmt.Println(multiplyTwoNumbers(12, 23)) // prints 276
}

func showAge(name string, age int) {
	fmt.Printf("%s is %d years old.\n", name, age)
}

func multiplyTwoNumbers(x, y int) int { // same type of arguments
	return x * y
}
