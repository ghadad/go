package main

import (
	"fmt"
)

type First func(int) int // declare type

func getFunction() First { // use it
	return func(val int) int {
		return val * 5
	}
}

func main() {

	f := getFunction() // returns a function of type First
	fmt.Println(f(12))
}
