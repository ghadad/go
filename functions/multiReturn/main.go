package main

import (
	"fmt"
)

func main() {
	vAdd, vSub, msg := addSub(35, 25)
	fmt.Printf("35 + 25 = %d\n", vAdd) // prints "35 + 25 = 60"
	fmt.Printf("35 - 25 = %d\n", vSub) // prints "35 - 25 = 10"
	fmt.Printf("Message is %s\n", msg) // prints "35 - 25 = 10"
}

func addSub(x, y int) (int, int, string) { // multiple return values (int, int)
	return x + y, x - y, "literal"
}
