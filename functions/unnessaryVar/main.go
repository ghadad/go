package main

import (
	"fmt"
)

func main() {
	j, _ := returnTwoNumbers() // ignores the second return value
	fmt.Println(j)
}

func returnTwoNumbers() (int, int) {
	return 12, 23
}
