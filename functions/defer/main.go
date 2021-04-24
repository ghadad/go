package main

import "fmt"

func main() {
	defer fmt.Println("before this.") // runs at the end
	fmt.Println("This is printed ")   // runs first
}
